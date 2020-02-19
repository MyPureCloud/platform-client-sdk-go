package platformclientv2
import (
	"encoding/json"
)

// Ipaddressrangelisting
type Ipaddressrangelisting struct { 
	// Entities
	Entities *[]Ipaddressrange `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ipaddressrangelisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
