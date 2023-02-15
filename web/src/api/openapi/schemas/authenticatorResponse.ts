/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */

/**
 * https://www.w3.org/TR/webauthn-2/#authenticatorresponse

 */
export interface AuthenticatorResponse {
  clientDataJSON: string;
  attestationObject?: string;
  transports?: string[];
  authenticatorData?: string;
  signature?: string;
  userHandle?: string;
}
