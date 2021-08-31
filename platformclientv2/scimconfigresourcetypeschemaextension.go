package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimconfigresourcetypeschemaextension - Defines a SCIM resource type's schema extension.
type Scimconfigresourcetypeschemaextension struct { 
	// Schema - The URI of an extended schema, for example, \"urn:edu:2.0:Staff\". Must be equal to the \"id\" attribute of a schema.
	Schema *string `json:"schema,omitempty"`


	// Required - Indicates whether a schema extension is required.
	Required *bool `json:"required,omitempty"`

}

func (o *Scimconfigresourcetypeschemaextension) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimconfigresourcetypeschemaextension
	
	return json.Marshal(&struct { 
		Schema *string `json:"schema,omitempty"`
		
		Required *bool `json:"required,omitempty"`
		*Alias
	}{ 
		Schema: o.Schema,
		
		Required: o.Required,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimconfigresourcetypeschemaextension) UnmarshalJSON(b []byte) error {
	var ScimconfigresourcetypeschemaextensionMap map[string]interface{}
	err := json.Unmarshal(b, &ScimconfigresourcetypeschemaextensionMap)
	if err != nil {
		return err
	}
	
	if Schema, ok := ScimconfigresourcetypeschemaextensionMap["schema"].(string); ok {
		o.Schema = &Schema
	}
	
	if Required, ok := ScimconfigresourcetypeschemaextensionMap["required"].(bool); ok {
		o.Required = &Required
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimconfigresourcetypeschemaextension) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
