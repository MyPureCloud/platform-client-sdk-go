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

func (u *Domainphysicalcapabilities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainphysicalcapabilities

	

	return json.Marshal(&struct { 
		Vlan *bool `json:"vlan,omitempty"`
		
		Team *bool `json:"team,omitempty"`
		*Alias
	}{ 
		Vlan: u.Vlan,
		
		Team: u.Team,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainphysicalcapabilities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
