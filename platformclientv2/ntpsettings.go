package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ntpsettings
type Ntpsettings struct { 
	// Servers - List of NTP servers, in priority order
	Servers *[]string `json:"servers,omitempty"`

}

func (o *Ntpsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ntpsettings
	
	return json.Marshal(&struct { 
		Servers *[]string `json:"servers,omitempty"`
		*Alias
	}{ 
		Servers: o.Servers,
		Alias:    (*Alias)(o),
	})
}

func (o *Ntpsettings) UnmarshalJSON(b []byte) error {
	var NtpsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &NtpsettingsMap)
	if err != nil {
		return err
	}
	
	if Servers, ok := NtpsettingsMap["servers"].([]interface{}); ok {
		ServersString, _ := json.Marshal(Servers)
		json.Unmarshal(ServersString, &o.Servers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ntpsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
