package platformclientv2
import (
	"encoding/json"
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
	Metric *int32 `json:"metric,omitempty"`


	// Family - The address family for this route.
	Family *int32 `json:"family,omitempty"`

}

// String returns a JSON representation of the model
func (o *Domainnetworkroute) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
