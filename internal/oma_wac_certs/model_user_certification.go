/*
 * Waiting List Api
 *
 * Ambulance Waiting List API
 *
 * API version: 1.0.3
 * Contact: xmartinkao@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oma_wac_certs

type UserCertification struct {

	// The relation ID
	Id string `json:"id,omitempty"`

	// The user ID
	UserId string `json:"user_id"`

	// The certification ID
	CertificationId string `json:"certification_id"`

	// The expiration date of the certification
	ExpiresAt string `json:"expires_at"`

	// The issue date of the certification
	IssuedAt string `json:"issued_at"`
}
