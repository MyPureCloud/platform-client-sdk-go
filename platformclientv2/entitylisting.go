package platformclientv2
import (
	"encoding/json"
)

// Entitylisting
type Entitylisting struct { 
	// Entities
	Entities *[]map[string]interface{} `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Entitylisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
