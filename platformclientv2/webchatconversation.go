package platformclientv2
import (
	"encoding/json"
)

// Webchatconversation
type Webchatconversation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Member - Chat Member
	Member *Webchatmemberinfo `json:"member,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}