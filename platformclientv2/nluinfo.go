package platformclientv2
import (
	"encoding/json"
)

// Nluinfo
type Nluinfo struct { 
	// Intents
	Intents *[]Intent `json:"intents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nluinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
