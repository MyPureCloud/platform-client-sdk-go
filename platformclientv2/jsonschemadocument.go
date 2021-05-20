package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Jsonschemadocument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
