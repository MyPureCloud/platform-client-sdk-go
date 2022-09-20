package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Testmodeeventresults - Information about event test mode execution
type Testmodeeventresults struct { 
	// SchemaValidation - Information about the validation of the schema of the event body passed in to test mode
	SchemaValidation *Testschemaoperation `json:"schemaValidation,omitempty"`


	// TriggerMatchValidation - Information about matched and unmatched triggers
	TriggerMatchValidation *Testmatcheseventoperation `json:"triggerMatchValidation,omitempty"`

}

func (o *Testmodeeventresults) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Testmodeeventresults
	
	return json.Marshal(&struct { 
		SchemaValidation *Testschemaoperation `json:"schemaValidation,omitempty"`
		
		TriggerMatchValidation *Testmatcheseventoperation `json:"triggerMatchValidation,omitempty"`
		*Alias
	}{ 
		SchemaValidation: o.SchemaValidation,
		
		TriggerMatchValidation: o.TriggerMatchValidation,
		Alias:    (*Alias)(o),
	})
}

func (o *Testmodeeventresults) UnmarshalJSON(b []byte) error {
	var TestmodeeventresultsMap map[string]interface{}
	err := json.Unmarshal(b, &TestmodeeventresultsMap)
	if err != nil {
		return err
	}
	
	if SchemaValidation, ok := TestmodeeventresultsMap["schemaValidation"].(map[string]interface{}); ok {
		SchemaValidationString, _ := json.Marshal(SchemaValidation)
		json.Unmarshal(SchemaValidationString, &o.SchemaValidation)
	}
	
	if TriggerMatchValidation, ok := TestmodeeventresultsMap["triggerMatchValidation"].(map[string]interface{}); ok {
		TriggerMatchValidationString, _ := json.Marshal(TriggerMatchValidation)
		json.Unmarshal(TriggerMatchValidationString, &o.TriggerMatchValidation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Testmodeeventresults) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
