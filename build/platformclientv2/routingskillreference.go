package platformclientv2
import (
	"encoding/json"
)

// Routingskillreference
type Routingskillreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routingskillreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
