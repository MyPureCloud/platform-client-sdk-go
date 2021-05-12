package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Domainnetworkroute
type Domainnetworkroute struct { 
	// Prefix - The IPv4 or IPv6 route prefix in CIDR notation.
	Prefix *string `json:"prefix,omitempty"`


	// Nexthop - The IPv4 or IPv6 nexthop IP address.
	Nexthop *string `json:"nexthop,omitempty"`


	// Persistent - True if this route will persist on Edge restart.  Routes assigned by DHCP will be returned as false.
	Persistent *bool `json:"persistent,omitempty"`


	// Metric - The metric being used for route. Lower values will have a higher priority.
	Metric *int `json:"metric,omitempty"`


	// Family - The address family for this route.
	Family *int `json:"family,omitempty"`

}

// String returns a JSON representation of the model
func (o *Domainnetworkroute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
