package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Credentialtype
type Credentialtype struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Properties - Properties describing credentials of this type.
	Properties *interface{} `json:"properties,omitempty"`


	// DisplayOrder - Order in which properties should be displayed in the UI.
	DisplayOrder *[]string `json:"displayOrder,omitempty"`


	// Required - Properties that are required fields.
	Required *[]string `json:"required,omitempty"`

}

func (o *Credentialtype) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Credentialtype
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Properties *interface{} `json:"properties,omitempty"`
		
		DisplayOrder *[]string `json:"displayOrder,omitempty"`
		
		Required *[]string `json:"required,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Properties: o.Properties,
		
		DisplayOrder: o.DisplayOrder,
		
		Required: o.Required,
		Alias:    (*Alias)(o),
	})
}

func (o *Credentialtype) UnmarshalJSON(b []byte) error {
	var CredentialtypeMap map[string]interface{}
	err := json.Unmarshal(b, &CredentialtypeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CredentialtypeMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := CredentialtypeMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Properties, ok := CredentialtypeMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if DisplayOrder, ok := CredentialtypeMap["displayOrder"].([]interface{}); ok {
		DisplayOrderString, _ := json.Marshal(DisplayOrder)
		json.Unmarshal(DisplayOrderString, &o.DisplayOrder)
	}
	
	if Required, ok := CredentialtypeMap["required"].([]interface{}); ok {
		RequiredString, _ := json.Marshal(Required)
		json.Unmarshal(RequiredString, &o.Required)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Credentialtype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
