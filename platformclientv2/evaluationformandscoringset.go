package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
