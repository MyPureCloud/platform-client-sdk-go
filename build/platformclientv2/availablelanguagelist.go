package platformclientv2
import (
	"encoding/json"
)

// Availablelanguagelist
type Availablelanguagelist struct { 
	// Languages
	Languages *[]string `json:"languages,omitempty"`

}

// String returns a JSON representation of the model
func (o *Availablelanguagelist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
