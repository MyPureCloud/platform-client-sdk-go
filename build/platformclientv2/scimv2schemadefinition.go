package platformclientv2
import (
	"encoding/json"
)

// Scimv2schemadefinition - Defines a SCIM schema.
type Scimv2schemadefinition struct { 
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readOnly\". \"returned\" is set to \"always\".
	Id *string `json:"id,omitempty"`


	// Name - The name of the schema.
	Name *string `json:"name,omitempty"`


	// Description - The description of the schema.
	Description *string `json:"description,omitempty"`


	// Attributes - The list of service provider attributes.
	Attributes *[]Scimv2schemaattribute `json:"attributes,omitempty"`


	// Meta - The metadata of the SCIM resource. Only \"location\" and \"resourceType\" are set for \"Schema\" resources.
	Meta *Scimmetadata `json:"meta,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2schemadefinition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
