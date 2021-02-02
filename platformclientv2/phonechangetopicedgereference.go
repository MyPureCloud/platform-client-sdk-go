package platformclientv2
import (
	"encoding/json"
)

// Phonechangetopicedgereference
type Phonechangetopicedgereference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonechangetopicedgereference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
