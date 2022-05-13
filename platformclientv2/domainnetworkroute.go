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

func (o *Domainnetworkroute) MarshalJSON() ([]byte, error) {
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
		Prefix: o.Prefix,
		
		Nexthop: o.Nexthop,
		
		Persistent: o.Persistent,
		
		Metric: o.Metric,
		
		Family: o.Family,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainnetworkroute) UnmarshalJSON(b []byte) error {
	var DomainnetworkrouteMap map[string]interface{}
	err := json.Unmarshal(b, &DomainnetworkrouteMap)
	if err != nil {
		return err
	}
	
	if Prefix, ok := DomainnetworkrouteMap["prefix"].(string); ok {
		o.Prefix = &Prefix
	}
    
	if Nexthop, ok := DomainnetworkrouteMap["nexthop"].(string); ok {
		o.Nexthop = &Nexthop
	}
    
	if Persistent, ok := DomainnetworkrouteMap["persistent"].(bool); ok {
		o.Persistent = &Persistent
	}
    
	if Metric, ok := DomainnetworkrouteMap["metric"].(float64); ok {
		MetricInt := int(Metric)
		o.Metric = &MetricInt
	}
	
	if Family, ok := DomainnetworkrouteMap["family"].(float64); ok {
		FamilyInt := int(Family)
		o.Family = &FamilyInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainnetworkroute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
