package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowoutcomenotificationerrormessageparams - The error message params, if the action failed
type Architectflowoutcomenotificationerrormessageparams struct { 
	// AdditionalProperties
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`

}

func (o *Architectflowoutcomenotificationerrormessageparams) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflowoutcomenotificationerrormessageparams
	
	return json.Marshal(&struct { 
		AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectflowoutcomenotificationerrormessageparams) UnmarshalJSON(b []byte) error {
	var ArchitectflowoutcomenotificationerrormessageparamsMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflowoutcomenotificationerrormessageparamsMap)
	if err != nil {
		return err
	}
	
	if AdditionalProperties, ok := ArchitectflowoutcomenotificationerrormessageparamsMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationerrormessageparams) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
