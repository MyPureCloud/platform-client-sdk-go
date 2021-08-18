package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportedlanguage
type Supportedlanguage struct { 
	// Language - Architect supported language tag, e.g. en-us, es-us
	Language *string `json:"language,omitempty"`


	// IsDefault - Whether or not this language is the default language
	IsDefault *bool `json:"isDefault,omitempty"`

}

func (u *Supportedlanguage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportedlanguage

	

	return json.Marshal(&struct { 
		Language *string `json:"language,omitempty"`
		
		IsDefault *bool `json:"isDefault,omitempty"`
		*Alias
	}{ 
		Language: u.Language,
		
		IsDefault: u.IsDefault,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Supportedlanguage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
