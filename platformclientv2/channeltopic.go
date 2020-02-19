package platformclientv2
import (
	"encoding/json"
)

// Channeltopic
type Channeltopic struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Channeltopic) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
