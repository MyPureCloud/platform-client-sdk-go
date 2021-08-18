package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Domainnetworkroute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainnetworkroute

	

	return json.Marshal(&struct { 
		Prefix *string `json:"prefix,omitempty"`
		
		Nexthop *string `json:"nexthop,omitempty"`
		
		Persistent *bool `json:"persistent,omitempty"`
		
		Metric *int `json:"metric,omitempty"`
		
		Family *int `json:"family,omitempty"`
		*Alias
	}{ 
		Prefix: u.Prefix,
		
		Nexthop: u.Nexthop,
		
		Persistent: u.Persistent,
		
		Metric: u.Metric,
		
		Family: u.Family,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainnetworkroute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
