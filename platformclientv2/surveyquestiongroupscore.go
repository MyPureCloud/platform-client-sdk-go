package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyquestiongroupscore
type Surveyquestiongroupscore struct { 
	// QuestionGroupId
	QuestionGroupId *string `json:"questionGroupId,omitempty"`


	// TotalScore - Score of all questions in the group
	TotalScore *float32 `json:"totalScore,omitempty"`


	// MaxTotalScore - Maximum possible score of all questions in the group
	MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`


	// MarkedNA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// QuestionScores
	QuestionScores *[]Surveyquestionscore `json:"questionScores,omitempty"`

}

func (o *Surveyquestiongroupscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyquestiongroupscore
	
	return json.Marshal(&struct { 
		QuestionGroupId *string `json:"questionGroupId,omitempty"`
		
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		QuestionScores *[]Surveyquestionscore `json:"questionScores,omitempty"`
		*Alias
	}{ 
		QuestionGroupId: o.QuestionGroupId,
		
		TotalScore: o.TotalScore,
		
		MaxTotalScore: o.MaxTotalScore,
		
		MarkedNA: o.MarkedNA,
		
		QuestionScores: o.QuestionScores,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyquestiongroupscore) UnmarshalJSON(b []byte) error {
	var SurveyquestiongroupscoreMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyquestiongroupscoreMap)
	if err != nil {
		return err
	}
	
	if QuestionGroupId, ok := SurveyquestiongroupscoreMap["questionGroupId"].(string); ok {
		o.QuestionGroupId = &QuestionGroupId
	}
	
	if TotalScore, ok := SurveyquestiongroupscoreMap["totalScore"].(float64); ok {
		TotalScoreFloat32 := float32(TotalScore)
		o.TotalScore = &TotalScoreFloat32
	}
	
	if MaxTotalScore, ok := SurveyquestiongroupscoreMap["maxTotalScore"].(float64); ok {
		MaxTotalScoreFloat32 := float32(MaxTotalScore)
		o.MaxTotalScore = &MaxTotalScoreFloat32
	}
	
	if MarkedNA, ok := SurveyquestiongroupscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
	
	if QuestionScores, ok := SurveyquestiongroupscoreMap["questionScores"].([]interface{}); ok {
		QuestionScoresString, _ := json.Marshal(QuestionScores)
		json.Unmarshal(QuestionScoresString, &o.QuestionScores)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
