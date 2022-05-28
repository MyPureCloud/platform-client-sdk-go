package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcenterfeedbacksettings
type Supportcenterfeedbacksettings struct { 
	// Enabled - Whether or not requesting customer feedback on article content and article search results is enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (o *Supportcenterfeedbacksettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportcenterfeedbacksettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportcenterfeedbacksettings) UnmarshalJSON(b []byte) error {
	var SupportcenterfeedbacksettingsMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcenterfeedbacksettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := SupportcenterfeedbacksettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcenterfeedbacksettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
