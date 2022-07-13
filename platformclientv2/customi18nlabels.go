package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Customi18nlabels - The localization settings for homescreen
type Customi18nlabels struct { 
	// Language - Language of localized labels in homescreen app (eg. en-us, de-de)
	Language *string `json:"language,omitempty"`


	// LocalizedLabels - Contains localized labels used in homescreen app
	LocalizedLabels *[]Localizedlabels `json:"localizedLabels,omitempty"`

}

func (o *Customi18nlabels) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Customi18nlabels
	
	return json.Marshal(&struct { 
		Language *string `json:"language,omitempty"`
		
		LocalizedLabels *[]Localizedlabels `json:"localizedLabels,omitempty"`
		*Alias
	}{ 
		Language: o.Language,
		
		LocalizedLabels: o.LocalizedLabels,
		Alias:    (*Alias)(o),
	})
}

func (o *Customi18nlabels) UnmarshalJSON(b []byte) error {
	var Customi18nlabelsMap map[string]interface{}
	err := json.Unmarshal(b, &Customi18nlabelsMap)
	if err != nil {
		return err
	}
	
	if Language, ok := Customi18nlabelsMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if LocalizedLabels, ok := Customi18nlabelsMap["localizedLabels"].([]interface{}); ok {
		LocalizedLabelsString, _ := json.Marshal(LocalizedLabels)
		json.Unmarshal(LocalizedLabelsString, &o.LocalizedLabels)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Customi18nlabels) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
