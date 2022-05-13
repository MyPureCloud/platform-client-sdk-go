package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Autostart
type Autostart struct { 
	// Enabled - whether or not auto start is enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (o *Autostart) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Autostart
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Autostart) UnmarshalJSON(b []byte) error {
	var AutostartMap map[string]interface{}
	err := json.Unmarshal(b, &AutostartMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := AutostartMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Autostart) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
