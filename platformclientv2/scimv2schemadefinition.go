package platformclientv2
import (
	"encoding/json"
)

// Scimv2schemadefinition - A SCIM schema definition.
type Scimv2schemadefinition struct { 
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"Mutability\" is set to \"readOnly\". \"Returned\" is set to \"always\".
	Id *string `json:"id,omitempty"`


	// Name - Schema name.
	Name *string `json:"name,omitempty"`


	// Description - Schema description.
	Description *string `json:"description,omitempty"`


	// Attributes - A complex type that defines service provider attributes and their qualities.
	Attributes *[]Scimv2schemaattribute `json:"attributes,omitempty"`


	// Meta - The metadata of the SCIM resource.
	Meta *Scimmetadata `json:"meta,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2schemadefinition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
