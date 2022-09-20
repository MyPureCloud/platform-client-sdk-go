package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Testmoderesults - Information about trigger test mode execution
type Testmoderesults struct { 
	// SchemaValidation - Information about the validation of the schema of the event body passed in to test mode
	SchemaValidation *Testschemaoperation `json:"schemaValidation,omitempty"`


	// TargetValidation - Information about the validation of the trigger target
	TargetValidation *Testtargetoperation `json:"targetValidation,omitempty"`


	// JsonPathValidation - Information about the json path matching criteria
	JsonPathValidation *Testmatchesoperation `json:"jsonPathValidation,omitempty"`


	// TriggerMatches - Whether the trigger would have matched on the provided event body
	TriggerMatches *bool `json:"triggerMatches,omitempty"`

}

func (o *Testmoderesults) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Testmoderesults
	
	return json.Marshal(&struct { 
		SchemaValidation *Testschemaoperation `json:"schemaValidation,omitempty"`
		
		TargetValidation *Testtargetoperation `json:"targetValidation,omitempty"`
		
		JsonPathValidation *Testmatchesoperation `json:"jsonPathValidation,omitempty"`
		
		TriggerMatches *bool `json:"triggerMatches,omitempty"`
		*Alias
	}{ 
		SchemaValidation: o.SchemaValidation,
		
		TargetValidation: o.TargetValidation,
		
		JsonPathValidation: o.JsonPathValidation,
		
		TriggerMatches: o.TriggerMatches,
		Alias:    (*Alias)(o),
	})
}

func (o *Testmoderesults) UnmarshalJSON(b []byte) error {
	var TestmoderesultsMap map[string]interface{}
	err := json.Unmarshal(b, &TestmoderesultsMap)
	if err != nil {
		return err
	}
	
	if SchemaValidation, ok := TestmoderesultsMap["schemaValidation"].(map[string]interface{}); ok {
		SchemaValidationString, _ := json.Marshal(SchemaValidation)
		json.Unmarshal(SchemaValidationString, &o.SchemaValidation)
	}
	
	if TargetValidation, ok := TestmoderesultsMap["targetValidation"].(map[string]interface{}); ok {
		TargetValidationString, _ := json.Marshal(TargetValidation)
		json.Unmarshal(TargetValidationString, &o.TargetValidation)
	}
	
	if JsonPathValidation, ok := TestmoderesultsMap["jsonPathValidation"].(map[string]interface{}); ok {
		JsonPathValidationString, _ := json.Marshal(JsonPathValidation)
		json.Unmarshal(JsonPathValidationString, &o.JsonPathValidation)
	}
	
	if TriggerMatches, ok := TestmoderesultsMap["triggerMatches"].(bool); ok {
		o.TriggerMatches = &TriggerMatches
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Testmoderesults) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
