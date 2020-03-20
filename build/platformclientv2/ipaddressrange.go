package platformclientv2
import (
	"encoding/json"
)

// Ipaddressrange
type Ipaddressrange struct { 
	// Cidr
	Cidr *string `json:"cidr,omitempty"`


	// Service
	Service *string `json:"service,omitempty"`


	// Region
	Region *string `json:"region,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ipaddressrange) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
