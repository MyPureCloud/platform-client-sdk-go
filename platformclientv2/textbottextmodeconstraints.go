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

}

func (o *Textbottextmodeconstraints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbottextmodeconstraints
	
	return json.Marshal(&struct { 
		LanguagePreferences *[]string `json:"languagePreferences,omitempty"`
		*Alias
	}{ 
		LanguagePreferences: o.LanguagePreferences,
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbottextmodeconstraints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
