package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lastattemptoverallconditionsettings
type Lastattemptoverallconditionsettings struct { 
	// MediaTypes - A list of media types to evaluate.
	MediaTypes *[]string `json:"mediaTypes,omitempty"`


	// Operator - The operator to use when comparing values.
	Operator *string `json:"operator,omitempty"`


	// Value - The period value to compare against the contact's data.
	Value *string `json:"value,omitempty"`

}

func (o *Lastattemptoverallconditionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lastattemptoverallconditionsettings
	
	return json.Marshal(&struct { 
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		MediaTypes: o.MediaTypes,
		
		Operator: o.Operator,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Lastattemptoverallconditionsettings) UnmarshalJSON(b []byte) error {
	var LastattemptoverallconditionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &LastattemptoverallconditionsettingsMap)
	if err != nil {
		return err
	}
	
	if MediaTypes, ok := LastattemptoverallconditionsettingsMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	
	if Operator, ok := LastattemptoverallconditionsettingsMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := LastattemptoverallconditionsettingsMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Lastattemptoverallconditionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
