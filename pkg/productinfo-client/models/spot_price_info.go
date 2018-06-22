// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// SpotPriceInfo SpotPriceInfo represents different prices per availability zones
// swagger:model SpotPriceInfo
type SpotPriceInfo map[string]float64

// Validate validates this spot price info
func (m SpotPriceInfo) Validate(formats strfmt.Registry) error {
	return nil
}