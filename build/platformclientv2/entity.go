package platformclientv2
import (
	"encoding/json"
)

// Entity
type Entity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Entity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
