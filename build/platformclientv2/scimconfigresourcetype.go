package platformclientv2
import (
	"encoding/json"
)

// Scimconfigresourcetype - Defines a SCIM resource.
type Scimconfigresourcetype struct { 
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readOnly\". \"returned\" is set to \"always\".
	Id *string `json:"id,omitempty"`


	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`


	// Name - The name of the resource type.
	Name *string `json:"name,omitempty"`


	// Description - The description of the resource type.
	Description *string `json:"description,omitempty"`


	// Schema - The URI of the primary or base schema for the resource type.
	Schema *string `json:"schema,omitempty"`


	// SchemaExtensions - The list of schema extensions for the resource type.
	SchemaExtensions *[]Scimconfigresourcetypeschemaextension `json:"schemaExtensions,omitempty"`


	// Endpoint - The HTTP-addressable endpoint of the resource type. Appears after the base URL.
	Endpoint *string `json:"endpoint,omitempty"`


	// Meta - The metadata of the SCIM resource. Only \"location\" and \"resourceType\" are set for \"ResourceType\" resources.
	Meta *Scimmetadata `json:"meta,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimconfigresourcetype) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
