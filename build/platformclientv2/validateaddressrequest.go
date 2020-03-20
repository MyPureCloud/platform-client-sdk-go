package platformclientv2
import (
	"encoding/json"
)

// Validateaddressrequest
type Validateaddressrequest struct { 
	// Address - Address schema
	Address *Streetaddress `json:"address,omitempty"`

}

// String returns a JSON representation of the model
func (o *Validateaddressrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
