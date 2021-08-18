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

func (u *Outcomeprobabilitycondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outcomeprobabilitycondition

	

	return json.Marshal(&struct { 
		OutcomeId *string `json:"outcomeId,omitempty"`
		
		MaximumProbability *float32 `json:"maximumProbability,omitempty"`
		
		Probability *float32 `json:"probability,omitempty"`
		*Alias
	}{ 
		OutcomeId: u.OutcomeId,
		
		MaximumProbability: u.MaximumProbability,
		
		Probability: u.Probability,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Outcomeprobabilitycondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
