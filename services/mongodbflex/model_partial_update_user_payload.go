/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

type PartialUpdateUserPayload struct {
	Database *string   `json:"database,omitempty"`
	Roles    *[]string `json:"roles,omitempty"`
}