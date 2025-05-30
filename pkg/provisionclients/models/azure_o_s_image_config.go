/*
Portions Copyright (c) Microsoft Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureOSImageConfig azure o s image config
//
// swagger:model AzureOSImageConfig
type AzureOSImageConfig struct {

	// image offer
	ImageOffer *string `json:"imageOffer,omitempty"`

	// image publisher
	ImagePublisher *string `json:"imagePublisher,omitempty"`

	// image sku
	ImageSku *string `json:"imageSku,omitempty"`

	// image version
	ImageVersion *string `json:"imageVersion,omitempty"`
}

// Validate validates this azure o s image config
func (m *AzureOSImageConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure o s image config based on context it is used
func (m *AzureOSImageConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureOSImageConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureOSImageConfig) UnmarshalBinary(b []byte) error {
	var res AzureOSImageConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
