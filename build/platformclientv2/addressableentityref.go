package platformclientv2
import (
	"encoding/json"
)

// Addressableentityref
type Addressableentityref struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Addressableentityref) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
