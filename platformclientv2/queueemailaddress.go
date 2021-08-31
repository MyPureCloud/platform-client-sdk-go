package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueemailaddress
type Queueemailaddress struct { 
	// Domain
	Domain *Domainentityref `json:"domain,omitempty"`


	// Route
	Route *Inboundroute `json:"route,omitempty"`

}

func (o *Queueemailaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueemailaddress
	
	return json.Marshal(&struct { 
		Domain *Domainentityref `json:"domain,omitempty"`
		
		Route *Inboundroute `json:"route,omitempty"`
		*Alias
	}{ 
		Domain: o.Domain,
		
		Route: o.Route,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueemailaddress) UnmarshalJSON(b []byte) error {
	var QueueemailaddressMap map[string]interface{}
	err := json.Unmarshal(b, &QueueemailaddressMap)
	if err != nil {
		return err
	}
	
	if Domain, ok := QueueemailaddressMap["domain"].(map[string]interface{}); ok {
		DomainString, _ := json.Marshal(Domain)
		json.Unmarshal(DomainString, &o.Domain)
	}
	
	if Route, ok := QueueemailaddressMap["route"].(map[string]interface{}); ok {
		RouteString, _ := json.Marshal(Route)
		json.Unmarshal(RouteString, &o.Route)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueemailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
