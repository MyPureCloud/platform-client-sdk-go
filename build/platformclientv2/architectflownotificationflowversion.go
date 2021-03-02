package platformclientv2
import (
	"encoding/json"
)

// Architectflownotificationflowversion
type Architectflownotificationflowversion struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflownotificationflowversion) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
