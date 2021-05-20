package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Credentialtype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
