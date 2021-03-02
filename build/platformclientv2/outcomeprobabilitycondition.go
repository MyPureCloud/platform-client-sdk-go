package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Outcomeprobabilitycondition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
