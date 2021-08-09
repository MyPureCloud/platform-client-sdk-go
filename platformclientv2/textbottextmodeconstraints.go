package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbottextmodeconstraints - Mode constraints to observe when operating on a bot flow.
type Textbottextmodeconstraints struct { 
	// LanguagePreferences - The list of language preferences by their ISO language code.
	LanguagePreferences *[]string `json:"languagePreferences,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbottextmodeconstraints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
