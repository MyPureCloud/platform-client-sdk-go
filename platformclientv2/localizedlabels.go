package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Localizedlabels - Contains localized labels used in messenger apps
type Localizedlabels struct { 
	// Key - Contains localized label key used in messenger homescreen
	Key *string `json:"key,omitempty"`


	// Value - Contains localized label value used in messenger homescreen
	Value *string `json:"value,omitempty"`

}

func (o *Localizedlabels) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Localizedlabels
	
	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Key: o.Key,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Localizedlabels) UnmarshalJSON(b []byte) error {
	var LocalizedlabelsMap map[string]interface{}
	err := json.Unmarshal(b, &LocalizedlabelsMap)
	if err != nil {
		return err
	}
	
	if Key, ok := LocalizedlabelsMap["key"].(string); ok {
		o.Key = &Key
	}
    
	if Value, ok := LocalizedlabelsMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Localizedlabels) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
