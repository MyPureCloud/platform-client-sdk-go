package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Cobrowsesettings - Settings concerning cobrowse
type Cobrowsesettings struct { 
	// Enabled - Whether or not cobrowse is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// AllowAgentControl - Whether the viewer should have option to request control
	AllowAgentControl *bool `json:"allowAgentControl,omitempty"`


	// MaskSelectors - Mask patterns that will apply to pages being shared
	MaskSelectors *[]string `json:"maskSelectors,omitempty"`

}

func (o *Cobrowsesettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Cobrowsesettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		AllowAgentControl *bool `json:"allowAgentControl,omitempty"`
		
		MaskSelectors *[]string `json:"maskSelectors,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		AllowAgentControl: o.AllowAgentControl,
		
		MaskSelectors: o.MaskSelectors,
		Alias:    (*Alias)(o),
	})
}

func (o *Cobrowsesettings) UnmarshalJSON(b []byte) error {
	var CobrowsesettingsMap map[string]interface{}
	err := json.Unmarshal(b, &CobrowsesettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := CobrowsesettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if AllowAgentControl, ok := CobrowsesettingsMap["allowAgentControl"].(bool); ok {
		o.AllowAgentControl = &AllowAgentControl
	}
	
	if MaskSelectors, ok := CobrowsesettingsMap["maskSelectors"].([]interface{}); ok {
		MaskSelectorsString, _ := json.Marshal(MaskSelectors)
		json.Unmarshal(MaskSelectorsString, &o.MaskSelectors)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Cobrowsesettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
