package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
