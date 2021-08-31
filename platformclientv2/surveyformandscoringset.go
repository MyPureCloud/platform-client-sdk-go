package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyformandscoringset
type Surveyformandscoringset struct { 
	// SurveyForm
	SurveyForm *Surveyform `json:"surveyForm,omitempty"`


	// Answers
	Answers *Surveyscoringset `json:"answers,omitempty"`

}

func (o *Surveyformandscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyformandscoringset
	
	return json.Marshal(&struct { 
		SurveyForm *Surveyform `json:"surveyForm,omitempty"`
		
		Answers *Surveyscoringset `json:"answers,omitempty"`
		*Alias
	}{ 
		SurveyForm: o.SurveyForm,
		
		Answers: o.Answers,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyformandscoringset) UnmarshalJSON(b []byte) error {
	var SurveyformandscoringsetMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyformandscoringsetMap)
	if err != nil {
		return err
	}
	
	if SurveyForm, ok := SurveyformandscoringsetMap["surveyForm"].(map[string]interface{}); ok {
		SurveyFormString, _ := json.Marshal(SurveyForm)
		json.Unmarshal(SurveyFormString, &o.SurveyForm)
	}
	
	if Answers, ok := SurveyformandscoringsetMap["answers"].(map[string]interface{}); ok {
		AnswersString, _ := json.Marshal(Answers)
		json.Unmarshal(AnswersString, &o.Answers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyformandscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
