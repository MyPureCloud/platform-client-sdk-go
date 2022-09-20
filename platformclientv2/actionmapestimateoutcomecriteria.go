package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmapestimateoutcomecriteria
type Actionmapestimateoutcomecriteria struct { 
	// OutcomeId - ID of outcome.
	OutcomeId *string `json:"outcomeId,omitempty"`


	// MaxProbability - Probability value for the selected outcome at or above which the action map will trigger.
	MaxProbability *float32 `json:"maxProbability,omitempty"`


	// Probability - Additional probability condition, where if set, the action map will trigger if the current outcome probability is lower or equal to the value.
	Probability *float32 `json:"probability,omitempty"`

}

func (o *Actionmapestimateoutcomecriteria) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionmapestimateoutcomecriteria
	
	return json.Marshal(&struct { 
		OutcomeId *string `json:"outcomeId,omitempty"`
		
		MaxProbability *float32 `json:"maxProbability,omitempty"`
		
		Probability *float32 `json:"probability,omitempty"`
		*Alias
	}{ 
		OutcomeId: o.OutcomeId,
		
		MaxProbability: o.MaxProbability,
		
		Probability: o.Probability,
		Alias:    (*Alias)(o),
	})
}

func (o *Actionmapestimateoutcomecriteria) UnmarshalJSON(b []byte) error {
	var ActionmapestimateoutcomecriteriaMap map[string]interface{}
	err := json.Unmarshal(b, &ActionmapestimateoutcomecriteriaMap)
	if err != nil {
		return err
	}
	
	if OutcomeId, ok := ActionmapestimateoutcomecriteriaMap["outcomeId"].(string); ok {
		o.OutcomeId = &OutcomeId
	}
    
	if MaxProbability, ok := ActionmapestimateoutcomecriteriaMap["maxProbability"].(float64); ok {
		MaxProbabilityFloat32 := float32(MaxProbability)
		o.MaxProbability = &MaxProbabilityFloat32
	}
	
	if Probability, ok := ActionmapestimateoutcomecriteriaMap["probability"].(float64); ok {
		ProbabilityFloat32 := float32(Probability)
		o.Probability = &ProbabilityFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionmapestimateoutcomecriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
