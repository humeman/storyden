/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { PublicKeyCredentialType } from "./publicKeyCredentialType";

/**
 * https://www.w3.org/TR/webauthn-2/#dictdef-publickeycredentialparameters

 */
export interface PublicKeyCredentialParameters {
  alg: number;
  type: PublicKeyCredentialType;
}
