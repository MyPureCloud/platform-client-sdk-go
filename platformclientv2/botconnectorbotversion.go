package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Botconnectorbotversion - A version description for a botConnector bot.
type Botconnectorbotversion struct { 
	// Version - The name of the version. This can be up to 100 characters long and must be comprised of displayable characters without leading or trailing whitespace
	Version *string `json:"version,omitempty"`


	// SupportedLanguages - The supported languages for this bot. EG 'en-us' or 'es', etc; These language codes are W3C language identification tags (ISO 639-1 for the language name and ISO 3166 for the country code)
	SupportedLanguages *[]string `json:"supportedLanguages,omitempty"`


	// Intents - A list of potential intents this bot will return, limit of 50
	Intents *[]Botintent `json:"intents,omitempty"`

}

func (o *Botconnectorbotversion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botconnectorbotversion
	
	return json.Marshal(&struct { 
		Version *string `json:"version,omitempty"`
		
		SupportedLanguages *[]string `json:"supportedLanguages,omitempty"`
		
		Intents *[]Botintent `json:"intents,omitempty"`
		*Alias
	}{ 
		Version: o.Version,
		
		SupportedLanguages: o.SupportedLanguages,
		
		Intents: o.Intents,
		Alias:    (*Alias)(o),
	})
}

func (o *Botconnectorbotversion) UnmarshalJSON(b []byte) error {
	var BotconnectorbotversionMap map[string]interface{}
	err := json.Unmarshal(b, &BotconnectorbotversionMap)
	if err != nil {
		return err
	}
	
	if Version, ok := BotconnectorbotversionMap["version"].(string); ok {
		o.Version = &Version
	}
	
	if SupportedLanguages, ok := BotconnectorbotversionMap["supportedLanguages"].([]interface{}); ok {
		SupportedLanguagesString, _ := json.Marshal(SupportedLanguages)
		json.Unmarshal(SupportedLanguagesString, &o.SupportedLanguages)
	}
	
	if Intents, ok := BotconnectorbotversionMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Botconnectorbotversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
