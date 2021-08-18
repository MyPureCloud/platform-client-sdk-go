package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Jsonschemadocument) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Jsonschemadocument

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Schema *string `json:"$schema,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Required *[]string `json:"required,omitempty"`
		
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Schema: u.Schema,
		
		Title: u.Title,
		
		Description: u.Description,
		
		VarType: u.VarType,
		
		Required: u.Required,
		
		Properties: u.Properties,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Jsonschemadocument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
