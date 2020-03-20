package platformclientv2
import (
	"encoding/json"
)

// Contentsortitem
type Contentsortitem struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Ascending
	Ascending *bool `json:"ascending,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentsortitem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
