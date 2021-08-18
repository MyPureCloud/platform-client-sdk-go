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

func (u *Credentialtype) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		Properties: u.Properties,
		
		DisplayOrder: u.DisplayOrder,
		
		Required: u.Required,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Credentialtype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
