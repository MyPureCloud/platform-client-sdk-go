package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Embeddedintegration
type Embeddedintegration struct { 
	// EnableWhitelist
	EnableWhitelist *bool `json:"enableWhitelist,omitempty"`


	// DomainWhitelist
	DomainWhitelist *[]string `json:"domainWhitelist,omitempty"`

}

func (o *Embeddedintegration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Embeddedintegration
	
	return json.Marshal(&struct { 
		EnableWhitelist *bool `json:"enableWhitelist,omitempty"`
		
		DomainWhitelist *[]string `json:"domainWhitelist,omitempty"`
		*Alias
	}{ 
		EnableWhitelist: o.EnableWhitelist,
		
		DomainWhitelist: o.DomainWhitelist,
		Alias:    (*Alias)(o),
	})
}

func (o *Embeddedintegration) UnmarshalJSON(b []byte) error {
	var EmbeddedintegrationMap map[string]interface{}
	err := json.Unmarshal(b, &EmbeddedintegrationMap)
	if err != nil {
		return err
	}
	
	if EnableWhitelist, ok := EmbeddedintegrationMap["enableWhitelist"].(bool); ok {
		o.EnableWhitelist = &EnableWhitelist
	}
    
	if DomainWhitelist, ok := EmbeddedintegrationMap["domainWhitelist"].([]interface{}); ok {
		DomainWhitelistString, _ := json.Marshal(DomainWhitelist)
		json.Unmarshal(DomainWhitelistString, &o.DomainWhitelist)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Embeddedintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
