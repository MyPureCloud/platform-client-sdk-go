package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ipaddressauthentication
type Ipaddressauthentication struct { 
	// NetworkWhitelist
	NetworkWhitelist *[]string `json:"networkWhitelist,omitempty"`

}

func (o *Ipaddressauthentication) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ipaddressauthentication
	
	return json.Marshal(&struct { 
		NetworkWhitelist *[]string `json:"networkWhitelist,omitempty"`
		*Alias
	}{ 
		NetworkWhitelist: o.NetworkWhitelist,
		Alias:    (*Alias)(o),
	})
}

func (o *Ipaddressauthentication) UnmarshalJSON(b []byte) error {
	var IpaddressauthenticationMap map[string]interface{}
	err := json.Unmarshal(b, &IpaddressauthenticationMap)
	if err != nil {
		return err
	}
	
	if NetworkWhitelist, ok := IpaddressauthenticationMap["networkWhitelist"].([]interface{}); ok {
		NetworkWhitelistString, _ := json.Marshal(NetworkWhitelist)
		json.Unmarshal(NetworkWhitelistString, &o.NetworkWhitelist)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ipaddressauthentication) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
