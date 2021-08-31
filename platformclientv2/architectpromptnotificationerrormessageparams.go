package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectpromptnotificationerrormessageparams
type Architectpromptnotificationerrormessageparams struct { 
	// AdditionalProperties
	AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`

}

func (o *Architectpromptnotificationerrormessageparams) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectpromptnotificationerrormessageparams
	
	return json.Marshal(&struct { 
		AdditionalProperties *map[string]string `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectpromptnotificationerrormessageparams) UnmarshalJSON(b []byte) error {
	var ArchitectpromptnotificationerrormessageparamsMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectpromptnotificationerrormessageparamsMap)
	if err != nil {
		return err
	}
	
	if AdditionalProperties, ok := ArchitectpromptnotificationerrormessageparamsMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectpromptnotificationerrormessageparams) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
