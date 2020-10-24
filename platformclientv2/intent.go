package platformclientv2
import (
	"encoding/json"
)

// Intent
type Intent struct { 
	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
