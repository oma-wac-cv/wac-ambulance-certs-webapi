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

type Certification struct {

	// The certification ID
	Id string `json:"id,omitempty"`

	// The certification name
	Name string `json:"name"`

	// The certification description
	Description string `json:"description"`

	// The authority that issued the certification
	Authority string `json:"authority"`
}
