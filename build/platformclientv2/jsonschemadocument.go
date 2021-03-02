package platformclientv2
import (
	"encoding/json"
)

// Jsonschemadocument - A JSON Schema document.
type Jsonschemadocument struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Schema
	Schema *string `json:"$schema,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Required
	Required *[]string `json:"required,omitempty"`


	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Jsonschemadocument) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
