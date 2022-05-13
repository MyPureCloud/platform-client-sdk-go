package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scorablesurvey
type Scorablesurvey struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SurveyForm - Survey form used for this survey.
	SurveyForm *Surveyform `json:"surveyForm,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// Answers
	Answers *Surveyscoringset `json:"answers,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Scorablesurvey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scorablesurvey
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SurveyForm *Surveyform `json:"surveyForm,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Answers *Surveyscoringset `json:"answers,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SurveyForm: o.SurveyForm,
		
		Status: o.Status,
		
		Answers: o.Answers,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Scorablesurvey) UnmarshalJSON(b []byte) error {
	var ScorablesurveyMap map[string]interface{}
	err := json.Unmarshal(b, &ScorablesurveyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ScorablesurveyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ScorablesurveyMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SurveyForm, ok := ScorablesurveyMap["surveyForm"].(map[string]interface{}); ok {
		SurveyFormString, _ := json.Marshal(SurveyForm)
		json.Unmarshal(SurveyFormString, &o.SurveyForm)
	}
	
	if Status, ok := ScorablesurveyMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Answers, ok := ScorablesurveyMap["answers"].(map[string]interface{}); ok {
		AnswersString, _ := json.Marshal(Answers)
		json.Unmarshal(AnswersString, &o.Answers)
	}
	
	if SelfUri, ok := ScorablesurveyMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scorablesurvey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
