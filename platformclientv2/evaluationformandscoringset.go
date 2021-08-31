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

func (o *Evaluationformandscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationformandscoringset
	
	return json.Marshal(&struct { 
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		Answers *Evaluationscoringset `json:"answers,omitempty"`
		*Alias
	}{ 
		EvaluationForm: o.EvaluationForm,
		
		Answers: o.Answers,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationformandscoringset) UnmarshalJSON(b []byte) error {
	var EvaluationformandscoringsetMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationformandscoringsetMap)
	if err != nil {
		return err
	}
	
	if EvaluationForm, ok := EvaluationformandscoringsetMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if Answers, ok := EvaluationformandscoringsetMap["answers"].(map[string]interface{}); ok {
		AnswersString, _ := json.Marshal(Answers)
		json.Unmarshal(AnswersString, &o.Answers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationformandscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
