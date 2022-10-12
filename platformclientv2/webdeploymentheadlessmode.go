package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentheadlessmode
type Webdeploymentheadlessmode struct { 
	// Enabled - Whether or not Headless Mode is enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (o *Webdeploymentheadlessmode) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentheadlessmode
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentheadlessmode) UnmarshalJSON(b []byte) error {
	var WebdeploymentheadlessmodeMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentheadlessmodeMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := WebdeploymentheadlessmodeMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentheadlessmode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
