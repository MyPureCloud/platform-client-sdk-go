package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Markcontactuncontactableactionsettings
type Markcontactuncontactableactionsettings struct { 
	// MediaTypes - A list of media types to evaluate.
	MediaTypes *[]string `json:"mediaTypes,omitempty"`

}

func (o *Markcontactuncontactableactionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Markcontactuncontactableactionsettings
	
	return json.Marshal(&struct { 
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		*Alias
	}{ 
		MediaTypes: o.MediaTypes,
		Alias:    (*Alias)(o),
	})
}

func (o *Markcontactuncontactableactionsettings) UnmarshalJSON(b []byte) error {
	var MarkcontactuncontactableactionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &MarkcontactuncontactableactionsettingsMap)
	if err != nil {
		return err
	}
	
	if MediaTypes, ok := MarkcontactuncontactableactionsettingsMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Markcontactuncontactableactionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
