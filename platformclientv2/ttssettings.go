package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ttssettings
type Ttssettings struct { 
	// DefaultEngine - ID of the global default TTS engine
	DefaultEngine *string `json:"defaultEngine,omitempty"`


	// LanguageOverrides - The list of default overrides for specific languages
	LanguageOverrides *[]Languageoverride `json:"languageOverrides,omitempty"`

}

func (o *Ttssettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ttssettings
	
	return json.Marshal(&struct { 
		DefaultEngine *string `json:"defaultEngine,omitempty"`
		
		LanguageOverrides *[]Languageoverride `json:"languageOverrides,omitempty"`
		*Alias
	}{ 
		DefaultEngine: o.DefaultEngine,
		
		LanguageOverrides: o.LanguageOverrides,
		Alias:    (*Alias)(o),
	})
}

func (o *Ttssettings) UnmarshalJSON(b []byte) error {
	var TtssettingsMap map[string]interface{}
	err := json.Unmarshal(b, &TtssettingsMap)
	if err != nil {
		return err
	}
	
	if DefaultEngine, ok := TtssettingsMap["defaultEngine"].(string); ok {
		o.DefaultEngine = &DefaultEngine
	}
	
	if LanguageOverrides, ok := TtssettingsMap["languageOverrides"].([]interface{}); ok {
		LanguageOverridesString, _ := json.Marshal(LanguageOverrides)
		json.Unmarshal(LanguageOverridesString, &o.LanguageOverrides)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ttssettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
