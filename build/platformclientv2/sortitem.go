package platformclientv2
import (
	"encoding/json"
)

// Sortitem
type Sortitem struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Ascending
	Ascending *bool `json:"ascending,omitempty"`

}

// String returns a JSON representation of the model
func (o *Sortitem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
