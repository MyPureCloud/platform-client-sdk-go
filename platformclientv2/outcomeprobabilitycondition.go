package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outcomeprobabilitycondition
type Outcomeprobabilitycondition struct { 
	// OutcomeId - The outcome ID.
	OutcomeId *string `json:"outcomeId,omitempty"`


	// MaximumProbability - Probability value for the selected outcome at or above which the action map will trigger.
	MaximumProbability *float32 `json:"maximumProbability,omitempty"`


	// Probability - Additional probability condition, where if set, the action map will trigger if the current outcome probability is lower or equal to the value.
	Probability *float32 `json:"probability,omitempty"`

}

func (o *Outcomeprobabilitycondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outcomeprobabilitycondition
	
	return json.Marshal(&struct { 
		OutcomeId *string `json:"outcomeId,omitempty"`
		
		MaximumProbability *float32 `json:"maximumProbability,omitempty"`
		
		Probability *float32 `json:"probability,omitempty"`
		*Alias
	}{ 
		OutcomeId: o.OutcomeId,
		
		MaximumProbability: o.MaximumProbability,
		
		Probability: o.Probability,
		Alias:    (*Alias)(o),
	})
}

func (o *Outcomeprobabilitycondition) UnmarshalJSON(b []byte) error {
	var OutcomeprobabilityconditionMap map[string]interface{}
	err := json.Unmarshal(b, &OutcomeprobabilityconditionMap)
	if err != nil {
		return err
	}
	
	if OutcomeId, ok := OutcomeprobabilityconditionMap["outcomeId"].(string); ok {
		o.OutcomeId = &OutcomeId
	}
    
	if MaximumProbability, ok := OutcomeprobabilityconditionMap["maximumProbability"].(float64); ok {
		MaximumProbabilityFloat32 := float32(MaximumProbability)
		o.MaximumProbability = &MaximumProbabilityFloat32
	}
	
	if Probability, ok := OutcomeprobabilityconditionMap["probability"].(float64); ok {
		ProbabilityFloat32 := float32(Probability)
		o.Probability = &ProbabilityFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outcomeprobabilitycondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
