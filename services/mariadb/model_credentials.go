/*
STACKIT MariaDB API

The STACKIT MariaDB API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mariadb

type Credentials struct {
	// REQUIRED
	Host  *string   `json:"host"`
	Hosts *[]string `json:"hosts,omitempty"`
	// for rabbitmq only
	HttpApiUri *string `json:"http_api_uri,omitempty"`
	Name       *string `json:"name,omitempty"`
	// REQUIRED
	Password  *string                 `json:"password"`
	Port      *int64                  `json:"port,omitempty"`
	Protocols *map[string]interface{} `json:"protocols,omitempty"`
	Uri       *string                 `json:"uri,omitempty"`
	// REQUIRED
	Username *string `json:"username"`
}
