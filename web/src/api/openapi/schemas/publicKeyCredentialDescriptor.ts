/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { PublicKeyCredentialDescriptorTransportsItem } from "./publicKeyCredentialDescriptorTransportsItem";
import type { PublicKeyCredentialType } from "./publicKeyCredentialType";

/**
 * https://www.w3.org/TR/webauthn-2/#dictdef-publickeycredentialdescriptor

 */
export interface PublicKeyCredentialDescriptor {
  type: PublicKeyCredentialType;
  id: string;
  transports?: PublicKeyCredentialDescriptorTransportsItem[];
}
