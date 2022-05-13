package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainphysicalcapabilities
type Domainphysicalcapabilities struct { 
	// Vlan
	Vlan *bool `json:"vlan,omitempty"`


	// Team
	Team *bool `json:"team,omitempty"`

}

func (o *Domainphysicalcapabilities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainphysicalcapabilities
	
	return json.Marshal(&struct { 
		Vlan *bool `json:"vlan,omitempty"`
		
		Team *bool `json:"team,omitempty"`
		*Alias
	}{ 
		Vlan: o.Vlan,
		
		Team: o.Team,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainphysicalcapabilities) UnmarshalJSON(b []byte) error {
	var DomainphysicalcapabilitiesMap map[string]interface{}
	err := json.Unmarshal(b, &DomainphysicalcapabilitiesMap)
	if err != nil {
		return err
	}
	
	if Vlan, ok := DomainphysicalcapabilitiesMap["vlan"].(bool); ok {
		o.Vlan = &Vlan
	}
    
	if Team, ok := DomainphysicalcapabilitiesMap["team"].(bool); ok {
		o.Team = &Team
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Domainphysicalcapabilities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
