package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Schema
type Schema struct { 
	// Title - A core type's title
	Title *string `json:"title,omitempty"`


	// Description - A core type's description
	Description *string `json:"description,omitempty"`


	// VarType - An array of fundamental JSON Schema primitive types on which the core type is based
	VarType *[]string `json:"type,omitempty"`


	// Items - Denotes the type and pattern of the items in an enum core type
	Items *Items `json:"items,omitempty"`


	// Pattern - For the \"date\" and \"datetime\" core types, denotes the regex prescribing the allowable date/datetime format
	Pattern *string `json:"pattern,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
