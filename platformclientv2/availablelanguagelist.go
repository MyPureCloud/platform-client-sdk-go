package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availablelanguagelist
type Availablelanguagelist struct { 
	// Languages
	Languages *[]string `json:"languages,omitempty"`

}

func (o *Availablelanguagelist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availablelanguagelist
	
	return json.Marshal(&struct { 
		Languages *[]string `json:"languages,omitempty"`
		*Alias
	}{ 
		Languages: o.Languages,
		Alias:    (*Alias)(o),
	})
}

func (o *Availablelanguagelist) UnmarshalJSON(b []byte) error {
	var AvailablelanguagelistMap map[string]interface{}
	err := json.Unmarshal(b, &AvailablelanguagelistMap)
	if err != nil {
		return err
	}
	
	if Languages, ok := AvailablelanguagelistMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Availablelanguagelist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
