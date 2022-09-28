package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Humanize
type Humanize struct { 
	// Enabled - Whether or not humanize conversations setting is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// Bot - Bot messenger profile setting
	Bot *Botmessengerprofile `json:"bot,omitempty"`

}

func (o *Humanize) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Humanize
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		Bot *Botmessengerprofile `json:"bot,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		Bot: o.Bot,
		Alias:    (*Alias)(o),
	})
}

func (o *Humanize) UnmarshalJSON(b []byte) error {
	var HumanizeMap map[string]interface{}
	err := json.Unmarshal(b, &HumanizeMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := HumanizeMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Bot, ok := HumanizeMap["bot"].(map[string]interface{}); ok {
		BotString, _ := json.Marshal(Bot)
		json.Unmarshal(BotString, &o.Bot)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Humanize) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
