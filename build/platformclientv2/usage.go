package platformclientv2
import (
	"encoding/json"
)

// Usage
type Usage struct { 
	// Types
	Types *[]Usageitem `json:"types,omitempty"`

}

// String returns a JSON representation of the model
func (o *Usage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
