package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Domainphysicalcapabilities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
