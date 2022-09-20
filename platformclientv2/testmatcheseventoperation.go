package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Testmatcheseventoperation - Results from evaluating matching criteria against test input
type Testmatcheseventoperation struct { 
	// Name - The name of the processing step
	Name *string `json:"name,omitempty"`


	// Step - The number of the processing step
	Step *int `json:"step,omitempty"`


	// MatchedTriggers - Triggers that matched
	MatchedTriggers *[]Testmodetrigger `json:"matchedTriggers,omitempty"`


	// UnmatchedTriggers - Triggers that did not match
	UnmatchedTriggers *[]Testmodetrigger `json:"unmatchedTriggers,omitempty"`

}

func (o *Testmatcheseventoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Testmatcheseventoperation
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Step *int `json:"step,omitempty"`
		
		MatchedTriggers *[]Testmodetrigger `json:"matchedTriggers,omitempty"`
		
		UnmatchedTriggers *[]Testmodetrigger `json:"unmatchedTriggers,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Step: o.Step,
		
		MatchedTriggers: o.MatchedTriggers,
		
		UnmatchedTriggers: o.UnmatchedTriggers,
		Alias:    (*Alias)(o),
	})
}

func (o *Testmatcheseventoperation) UnmarshalJSON(b []byte) error {
	var TestmatcheseventoperationMap map[string]interface{}
	err := json.Unmarshal(b, &TestmatcheseventoperationMap)
	if err != nil {
		return err
	}
	
	if Name, ok := TestmatcheseventoperationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Step, ok := TestmatcheseventoperationMap["step"].(float64); ok {
		StepInt := int(Step)
		o.Step = &StepInt
	}
	
	if MatchedTriggers, ok := TestmatcheseventoperationMap["matchedTriggers"].([]interface{}); ok {
		MatchedTriggersString, _ := json.Marshal(MatchedTriggers)
		json.Unmarshal(MatchedTriggersString, &o.MatchedTriggers)
	}
	
	if UnmatchedTriggers, ok := TestmatcheseventoperationMap["unmatchedTriggers"].([]interface{}); ok {
		UnmatchedTriggersString, _ := json.Marshal(UnmatchedTriggers)
		json.Unmarshal(UnmatchedTriggersString, &o.UnmatchedTriggers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Testmatcheseventoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
