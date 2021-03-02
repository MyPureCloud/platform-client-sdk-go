package platformclientv2
import (
	"encoding/json"
)

// Evaluationformandscoringset
type Evaluationformandscoringset struct { 
	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// Answers
	Answers *Evaluationscoringset `json:"answers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationformandscoringset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
