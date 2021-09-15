package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outcomeeventscore
type Outcomeeventscore struct { 
	// Outcome - The outcome that the score was calculated for.
	Outcome *Addressableentityref `json:"outcome,omitempty"`


	// SessionMaxProbability - Represents the max probability reached in the session.
	SessionMaxProbability *float32 `json:"sessionMaxProbability,omitempty"`


	// Probability - Represents the likelihood of a customer reaching or achieving a given outcome.
	Probability *float32 `json:"probability,omitempty"`

}

func (o *Outcomeeventscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outcomeeventscore
	
	return json.Marshal(&struct { 
		Outcome *Addressableentityref `json:"outcome,omitempty"`
		
		SessionMaxProbability *float32 `json:"sessionMaxProbability,omitempty"`
		
		Probability *float32 `json:"probability,omitempty"`
		*Alias
	}{ 
		Outcome: o.Outcome,
		
		SessionMaxProbability: o.SessionMaxProbability,
		
		Probability: o.Probability,
		Alias:    (*Alias)(o),
	})
}

func (o *Outcomeeventscore) UnmarshalJSON(b []byte) error {
	var OutcomeeventscoreMap map[string]interface{}
	err := json.Unmarshal(b, &OutcomeeventscoreMap)
	if err != nil {
		return err
	}
	
	if Outcome, ok := OutcomeeventscoreMap["outcome"].(map[string]interface{}); ok {
		OutcomeString, _ := json.Marshal(Outcome)
		json.Unmarshal(OutcomeString, &o.Outcome)
	}
	
	if SessionMaxProbability, ok := OutcomeeventscoreMap["sessionMaxProbability"].(float64); ok {
		SessionMaxProbabilityFloat32 := float32(SessionMaxProbability)
		o.SessionMaxProbability = &SessionMaxProbabilityFloat32
	}
	
	if Probability, ok := OutcomeeventscoreMap["probability"].(float64); ok {
		ProbabilityFloat32 := float32(Probability)
		o.Probability = &ProbabilityFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outcomeeventscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
