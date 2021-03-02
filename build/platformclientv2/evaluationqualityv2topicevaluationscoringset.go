package platformclientv2
import (
	"encoding/json"
)

// Evaluationqualityv2topicevaluationscoringset
type Evaluationqualityv2topicevaluationscoringset struct { 
	// TotalScore
	TotalScore *int `json:"totalScore,omitempty"`


	// TotalCriticalScore
	TotalCriticalScore *int `json:"totalCriticalScore,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationqualityv2topicevaluationscoringset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
