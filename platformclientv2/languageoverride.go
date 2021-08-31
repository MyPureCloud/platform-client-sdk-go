package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Languageoverride
type Languageoverride struct { 
	// Language - The language code of the language being overridden
	Language *string `json:"language,omitempty"`


	// Engine - The ID of the TTS engine to use for this language override
	Engine *string `json:"engine,omitempty"`


	// Voice - The ID of the voice to use for this language override. The voice must be supported by the chosen engine.
	Voice *string `json:"voice,omitempty"`

}

func (o *Languageoverride) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Languageoverride
	
	return json.Marshal(&struct { 
		Language *string `json:"language,omitempty"`
		
		Engine *string `json:"engine,omitempty"`
		
		Voice *string `json:"voice,omitempty"`
		*Alias
	}{ 
		Language: o.Language,
		
		Engine: o.Engine,
		
		Voice: o.Voice,
		Alias:    (*Alias)(o),
	})
}

func (o *Languageoverride) UnmarshalJSON(b []byte) error {
	var LanguageoverrideMap map[string]interface{}
	err := json.Unmarshal(b, &LanguageoverrideMap)
	if err != nil {
		return err
	}
	
	if Language, ok := LanguageoverrideMap["language"].(string); ok {
		o.Language = &Language
	}
	
	if Engine, ok := LanguageoverrideMap["engine"].(string); ok {
		o.Engine = &Engine
	}
	
	if Voice, ok := LanguageoverrideMap["voice"].(string); ok {
		o.Voice = &Voice
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Languageoverride) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
