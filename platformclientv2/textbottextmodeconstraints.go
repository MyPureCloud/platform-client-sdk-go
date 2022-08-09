package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbottextmodeconstraints - Mode constraints to observe when operating on a bot flow.
type Textbottextmodeconstraints struct { 
	// LanguagePreferences - The list of language preferences by their ISO language code.
	LanguagePreferences *[]string `json:"languagePreferences,omitempty"`


	// NoInputTimeoutMilliseconds - The amount of time, in milliseconds, before the client should send the 'NoInput' event  to trigger the \"no input\" bot response and handling on digital channels.  Note: This optional field will only be returned for 'Digital Bot Flow' turns.
	NoInputTimeoutMilliseconds *int `json:"noInputTimeoutMilliseconds,omitempty"`

}

func (o *Textbottextmodeconstraints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbottextmodeconstraints
	
	return json.Marshal(&struct { 
		LanguagePreferences *[]string `json:"languagePreferences,omitempty"`
		
		NoInputTimeoutMilliseconds *int `json:"noInputTimeoutMilliseconds,omitempty"`
		*Alias
	}{ 
		LanguagePreferences: o.LanguagePreferences,
		
		NoInputTimeoutMilliseconds: o.NoInputTimeoutMilliseconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbottextmodeconstraints) UnmarshalJSON(b []byte) error {
	var TextbottextmodeconstraintsMap map[string]interface{}
	err := json.Unmarshal(b, &TextbottextmodeconstraintsMap)
	if err != nil {
		return err
	}
	
	if LanguagePreferences, ok := TextbottextmodeconstraintsMap["languagePreferences"].([]interface{}); ok {
		LanguagePreferencesString, _ := json.Marshal(LanguagePreferences)
		json.Unmarshal(LanguagePreferencesString, &o.LanguagePreferences)
	}
	
	if NoInputTimeoutMilliseconds, ok := TextbottextmodeconstraintsMap["noInputTimeoutMilliseconds"].(float64); ok {
		NoInputTimeoutMillisecondsInt := int(NoInputTimeoutMilliseconds)
		o.NoInputTimeoutMilliseconds = &NoInputTimeoutMillisecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbottextmodeconstraints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
