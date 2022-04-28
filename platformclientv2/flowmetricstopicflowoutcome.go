package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowmetricstopicflowoutcome
type Flowmetricstopicflowoutcome struct { 
	// FlowOutcome - Combination of unique flow outcome identifier and its value separated by colon
	FlowOutcome *string `json:"flowOutcome,omitempty"`


	// FlowOutcomeId - Unique identifier of a flow outcome
	FlowOutcomeId *string `json:"flowOutcomeId,omitempty"`


	// FlowOutcomeValue - Flow outcome value, e.g. SUCCESS
	FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`

}

func (o *Flowmetricstopicflowoutcome) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowmetricstopicflowoutcome
	
	return json.Marshal(&struct { 
		FlowOutcome *string `json:"flowOutcome,omitempty"`
		
		FlowOutcomeId *string `json:"flowOutcomeId,omitempty"`
		
		FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`
		*Alias
	}{ 
		FlowOutcome: o.FlowOutcome,
		
		FlowOutcomeId: o.FlowOutcomeId,
		
		FlowOutcomeValue: o.FlowOutcomeValue,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowmetricstopicflowoutcome) UnmarshalJSON(b []byte) error {
	var FlowmetricstopicflowoutcomeMap map[string]interface{}
	err := json.Unmarshal(b, &FlowmetricstopicflowoutcomeMap)
	if err != nil {
		return err
	}
	
	if FlowOutcome, ok := FlowmetricstopicflowoutcomeMap["flowOutcome"].(string); ok {
		o.FlowOutcome = &FlowOutcome
	}
	
	if FlowOutcomeId, ok := FlowmetricstopicflowoutcomeMap["flowOutcomeId"].(string); ok {
		o.FlowOutcomeId = &FlowOutcomeId
	}
	
	if FlowOutcomeValue, ok := FlowmetricstopicflowoutcomeMap["flowOutcomeValue"].(string); ok {
		o.FlowOutcomeValue = &FlowOutcomeValue
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowmetricstopicflowoutcome) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
