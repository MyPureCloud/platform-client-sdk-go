package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassessmentscoringrequest
type Learningassessmentscoringrequest struct { 
	// AssessmentForm - The assessment form to score against
	AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`


	// Answers - The answers to score
	Answers *Assessmentscoringset `json:"answers,omitempty"`

}

func (o *Learningassessmentscoringrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassessmentscoringrequest
	
	return json.Marshal(&struct { 
		AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`
		
		Answers *Assessmentscoringset `json:"answers,omitempty"`
		*Alias
	}{ 
		AssessmentForm: o.AssessmentForm,
		
		Answers: o.Answers,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassessmentscoringrequest) UnmarshalJSON(b []byte) error {
	var LearningassessmentscoringrequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassessmentscoringrequestMap)
	if err != nil {
		return err
	}
	
	if AssessmentForm, ok := LearningassessmentscoringrequestMap["assessmentForm"].(map[string]interface{}); ok {
		AssessmentFormString, _ := json.Marshal(AssessmentForm)
		json.Unmarshal(AssessmentFormString, &o.AssessmentForm)
	}
	
	if Answers, ok := LearningassessmentscoringrequestMap["answers"].(map[string]interface{}); ok {
		AnswersString, _ := json.Marshal(Answers)
		json.Unmarshal(AnswersString, &o.Answers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassessmentscoringrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
