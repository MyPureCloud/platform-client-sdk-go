package platformclientv2
import (
	"encoding/json"
)

// Replacerequest
type Replacerequest struct { 
	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// AuthToken
	AuthToken *string `json:"authToken,omitempty"`

}

// String returns a JSON representation of the model
func (o *Replacerequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
