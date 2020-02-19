package platformclientv2
import (
	"encoding/json"
)

// Credentialtype
type Credentialtype struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Properties - Properties describing credentials of this type.
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// DisplayOrder - Order in which properties should be displayed in the UI.
	DisplayOrder *[]string `json:"displayOrder,omitempty"`


	// Required - Properties that are required fields.
	Required *[]string `json:"required,omitempty"`

}

// String returns a JSON representation of the model
func (o *Credentialtype) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
