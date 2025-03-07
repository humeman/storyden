package openapi

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/fmsg"
	"github.com/Southclaws/fault/ftag"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
)

// ValidatorErrorHandler is an OpenAPI validator function for structured errors.
// This is used with `OapiRequestValidatorWithOptions` in order to handle errors
// that occur at the request validation level before the request reaches the app
// layer. Validation failures happen due to security schemes not being satisfied
// or requests not satisfying the specification. The purpose of this function is
// to process the various types of error that can occur and turn the errors into
// structured errors which can be serialised properly with `fctx` for logging.
func ValidatorErrorHandler() func(c echo.Context, err *echo.HTTPError) error {
	var fn func(ctx context.Context, err error) error

	fn = func(ctx context.Context, err error) error {
		he := &echo.HTTPError{}
		if errors.As(err, &he) {
			if he.Internal == nil {
				// if the wrapped error is nil, just return the error as-is.
				message, ok := he.Message.(string)
				if !ok {
					return err
				}

				return fault.Wrap(err, fmsg.With(message))
			}

			// Most of the time, these errors contain a more detailed error.
			return fn(ctx, he.Internal)
		}

		// These occur when calling the wrong endpoint.
		re := &openapi3filter.RequestError{}
		if errors.As(err, &re) {
			if re.Err == nil {
				return err
			}

			return fn(ctx, fault.Wrap(re.Err,
				fctx.With(ctx),
				fmsg.With(re.Reason)),
			)
		}

		// These occur when the payload doesn't match the schema.
		se := &openapi3.SchemaError{}
		if errors.As(err, &se) {
			ctx = fctx.WithMeta(ctx,
				"schema_error", se.Reason,
				"schema_field", se.SchemaField,
				"schema_description", se.Schema.Description,
				"schema_type", fmt.Sprint(se.Schema.Type),
				"path", strings.Join(se.JSONPointer(), "."),
			)

			// These schema errors become nested when the schema is composed of
			// other schemas using anyOf, allOf, oneOf, etc. This recursion will
			// make sure the root error is the one that is used for reporting.
			if se.Origin != nil {
				return fn(ctx, se.Origin)
			}

			return fault.Wrap(fault.New("openapi schema validation failure"),
				fctx.With(ctx),
				ftag.With(ftag.InvalidArgument),
				fmsg.With("request parameters do not match schema"),
			)
		}

		// These occur when there's an authentication problem.
		sr := &openapi3filter.SecurityRequirementsError{}
		if errors.As(err, &sr) {
			// For some reason, OpenAPI stores the resulting validation
			// error as an array. It's always 1 element long (for now) so
			// just grab the first element if it's there for the err value.
			wrapped := err
			if len(sr.Errors) == 1 {
				wrapped = sr.Errors[0]
			}

			return fault.Wrap(wrapped,
				fctx.With(ctx),
				ftag.With(ftag.Unauthenticated),
				fmsg.WithDesc(sr.Error(), "The request did not contain any authentication information, please check to make sure you are logged in."),
			)
		}

		// By default, validation errors are treated as the client's fault.
		return fault.Wrap(err, ftag.With(ftag.InvalidArgument))
	}

	return func(c echo.Context, err *echo.HTTPError) error {
		return fn(c.Request().Context(), err)
	}
}
