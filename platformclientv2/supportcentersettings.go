package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcentersettings - Settings concerning support center
type Supportcentersettings struct { 
	// Enabled - Whether or not support center is enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (o *Supportcentersettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportcentersettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportcentersettings) UnmarshalJSON(b []byte) error {
	var SupportcentersettingsMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcentersettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := SupportcentersettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcentersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
