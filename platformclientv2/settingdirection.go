package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Settingdirection
type Settingdirection struct { 
	// Inbound - Status for the Inbound Direction
	Inbound *string `json:"inbound,omitempty"`


	// Outbound - Status for the Outbound Direction
	Outbound *string `json:"outbound,omitempty"`

}

func (o *Settingdirection) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Settingdirection
	
	return json.Marshal(&struct { 
		Inbound *string `json:"inbound,omitempty"`
		
		Outbound *string `json:"outbound,omitempty"`
		*Alias
	}{ 
		Inbound: o.Inbound,
		
		Outbound: o.Outbound,
		Alias:    (*Alias)(o),
	})
}

func (o *Settingdirection) UnmarshalJSON(b []byte) error {
	var SettingdirectionMap map[string]interface{}
	err := json.Unmarshal(b, &SettingdirectionMap)
	if err != nil {
		return err
	}
	
	if Inbound, ok := SettingdirectionMap["inbound"].(string); ok {
		o.Inbound = &Inbound
	}
	
	if Outbound, ok := SettingdirectionMap["outbound"].(string); ok {
		o.Outbound = &Outbound
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Settingdirection) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
