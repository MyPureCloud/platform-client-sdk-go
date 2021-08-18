package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Evaluationformandscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationformandscoringset

	

	return json.Marshal(&struct { 
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		Answers *Evaluationscoringset `json:"answers,omitempty"`
		*Alias
	}{ 
		EvaluationForm: u.EvaluationForm,
		
		Answers: u.Answers,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluationformandscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
