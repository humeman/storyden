package google

import (
	"context"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/fmsg"
	"github.com/Southclaws/fault/ftag"
	petname "github.com/dustinkirkland/golang-petname"
	"github.com/gosimple/slug"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	ga "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/account/authentication"
	"github.com/Southclaws/storyden/app/services/account/register"
	"github.com/Southclaws/storyden/app/services/authentication/provider/oauth/all"
	"github.com/Southclaws/storyden/internal/config"
	"github.com/Southclaws/storyden/internal/infrastructure/endec"
)

var (
	ErrAccessToken  = fault.New("failed to get access token")
	ErrMissingToken = fault.New("no access token in response")
)

var (
	service   = authentication.ServiceOAuthGoogle
	tokenType = authentication.TokenTypeOAuth
)

type Provider struct {
	register *register.Registrar
	ed       endec.EncrypterDecrypter

	callback string
	config   *all.Configuration
}

func New(
	cfg config.Config,
	register *register.Registrar,
	ed endec.EncrypterDecrypter,
) (*Provider, error) {
	config, err := all.LoadProvider(service)
	if err != nil {
		return nil, fault.Wrap(err)
	}

	callback := all.Redirect(cfg, service)

	return &Provider{
		register: register,
		ed:       ed,
		config:   config,
		callback: callback,
	}, nil
}

func (p *Provider) Service() authentication.Service { return service }
func (p *Provider) Token() authentication.TokenType { return tokenType }

func (p *Provider) Enabled(ctx context.Context) (bool, error) {
	return p.config != nil, nil
}

func (p *Provider) oauthConfig() *oauth2.Config {
	if p.config == nil {
		return nil
	}

	return &oauth2.Config{
		ClientID:     p.config.ClientID,
		ClientSecret: p.config.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  p.callback,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}

func (p *Provider) Link(redirectPath string) (string, error) {
	state, err := p.ed.Encrypt(map[string]any{
		"redirect": redirectPath,
	}, time.Minute*10)
	if err != nil {
		return "", fault.Wrap(err)
	}

	oac := p.oauthConfig()

	return oac.AuthCodeURL(state, oauth2.AccessTypeOffline), nil
}

func (p *Provider) Login(ctx context.Context, state, code string) (*account.Account, error) {
	_, err := p.ed.Decrypt(state)
	if err != nil {
		return nil, fault.Wrap(err,
			fctx.With(ctx),
			fmsg.WithDesc("failed to decrypt state value", "This link has expired, please try again."),
		)
	}

	oac := p.oauthConfig()

	token, err := oac.Exchange(ctx, code, oauth2.AccessTypeOffline)
	if err != nil {
		return nil, fault.Wrap(err,
			fctx.With(ctx),
			ftag.With(ftag.InvalidArgument),
			fmsg.WithDesc("failed to exchange code for token", "This login token may have expired, please try again."),
		)
	}

	gs, err := ga.NewService(ctx, option.WithTokenSource(oauth2.StaticTokenSource(token)))
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	u, err := ga.NewUserinfoV2MeService(gs).Get().Do()
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	handle := slug.Make(fmt.Sprintf("%s %s", u.GivenName, petname.Generate(2, "-")))

	name := fmt.Sprint(u.GivenName, " ", u.FamilyName)

	email, err := mail.ParseAddress(u.Email)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	authName := fmt.Sprintf("Google (%s)", email.Address)

	acc, err := p.register.GetOrCreateViaEmail(ctx,
		service,
		authName,
		strings.ToLower(u.Id),
		token.AccessToken,
		handle,
		name,
		*email,
	)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return acc, nil
}
