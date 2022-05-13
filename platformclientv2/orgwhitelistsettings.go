package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Orgwhitelistsettings
type Orgwhitelistsettings struct { 
	// EnableWhitelist
	EnableWhitelist *bool `json:"enableWhitelist,omitempty"`


	// DomainWhitelist
	DomainWhitelist *[]string `json:"domainWhitelist,omitempty"`

}

func (o *Orgwhitelistsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Orgwhitelistsettings
	
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

func (o *Orgwhitelistsettings) UnmarshalJSON(b []byte) error {
	var OrgwhitelistsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &OrgwhitelistsettingsMap)
	if err != nil {
		return err
	}
	
	if EnableWhitelist, ok := OrgwhitelistsettingsMap["enableWhitelist"].(bool); ok {
		o.EnableWhitelist = &EnableWhitelist
	}
    
	if DomainWhitelist, ok := OrgwhitelistsettingsMap["domainWhitelist"].([]interface{}); ok {
		DomainWhitelistString, _ := json.Marshal(DomainWhitelist)
		json.Unmarshal(DomainWhitelistString, &o.DomainWhitelist)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Orgwhitelistsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
