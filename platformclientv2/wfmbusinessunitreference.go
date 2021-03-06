package platformclientv2
import (
	"encoding/json"
)

// Wfmbusinessunitreference
type Wfmbusinessunitreference struct { 
	// Id - The ID of the business unit
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbusinessunitreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
