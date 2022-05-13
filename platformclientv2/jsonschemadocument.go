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

func (o *Jsonschemadocument) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Schema: o.Schema,
		
		Title: o.Title,
		
		Description: o.Description,
		
		VarType: o.VarType,
		
		Required: o.Required,
		
		Properties: o.Properties,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Jsonschemadocument) UnmarshalJSON(b []byte) error {
	var JsonschemadocumentMap map[string]interface{}
	err := json.Unmarshal(b, &JsonschemadocumentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JsonschemadocumentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Schema, ok := JsonschemadocumentMap["$schema"].(string); ok {
		o.Schema = &Schema
	}
    
	if Title, ok := JsonschemadocumentMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := JsonschemadocumentMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if VarType, ok := JsonschemadocumentMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Required, ok := JsonschemadocumentMap["required"].([]interface{}); ok {
		RequiredString, _ := json.Marshal(Required)
		json.Unmarshal(RequiredString, &o.Required)
	}
	
	if Properties, ok := JsonschemadocumentMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if AdditionalProperties, ok := JsonschemadocumentMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Jsonschemadocument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
