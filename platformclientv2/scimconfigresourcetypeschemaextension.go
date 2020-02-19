package platformclientv2
import (
	"encoding/json"
)

// Scimconfigresourcetypeschemaextension - Defines a SCIM resource type's schema extension.
type Scimconfigresourcetypeschemaextension struct { 
	// Schema - The URI of an extended schema, for example, \"urn:edu:2.0:Staff\". Must be equal to the \"id\" attribute of a schema.
	Schema *string `json:"schema,omitempty"`


	// Required - Indicates whether a schema extension is required.
	Required *bool `json:"required,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimconfigresourcetypeschemaextension) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
