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

func (u *Languageoverride) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Languageoverride

	

	return json.Marshal(&struct { 
		Language *string `json:"language,omitempty"`
		
		Engine *string `json:"engine,omitempty"`
		
		Voice *string `json:"voice,omitempty"`
		*Alias
	}{ 
		Language: u.Language,
		
		Engine: u.Engine,
		
		Voice: u.Voice,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Languageoverride) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
