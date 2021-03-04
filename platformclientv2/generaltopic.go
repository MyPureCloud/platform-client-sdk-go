package platformclientv2
import (
	"encoding/json"
)

// Generaltopic
type Generaltopic struct { 
	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Generaltopic) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
