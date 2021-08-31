package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyscoringset
type Surveyscoringset struct { 
	// TotalScore
	TotalScore *float32 `json:"totalScore,omitempty"`


	// NpsScore
	NpsScore *int `json:"npsScore,omitempty"`


	// QuestionGroupScores
	QuestionGroupScores *[]Surveyquestiongroupscore `json:"questionGroupScores,omitempty"`

}

func (o *Surveyscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyscoringset
	
	return json.Marshal(&struct { 
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		NpsScore *int `json:"npsScore,omitempty"`
		
		QuestionGroupScores *[]Surveyquestiongroupscore `json:"questionGroupScores,omitempty"`
		*Alias
	}{ 
		TotalScore: o.TotalScore,
		
		NpsScore: o.NpsScore,
		
		QuestionGroupScores: o.QuestionGroupScores,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyscoringset) UnmarshalJSON(b []byte) error {
	var SurveyscoringsetMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyscoringsetMap)
	if err != nil {
		return err
	}
	
	if TotalScore, ok := SurveyscoringsetMap["totalScore"].(float64); ok {
		TotalScoreFloat32 := float32(TotalScore)
		o.TotalScore = &TotalScoreFloat32
	}
	
	if NpsScore, ok := SurveyscoringsetMap["npsScore"].(float64); ok {
		NpsScoreInt := int(NpsScore)
		o.NpsScore = &NpsScoreInt
	}
	
	if QuestionGroupScores, ok := SurveyscoringsetMap["questionGroupScores"].([]interface{}); ok {
		QuestionGroupScoresString, _ := json.Marshal(QuestionGroupScores)
		json.Unmarshal(QuestionGroupScoresString, &o.QuestionGroupScores)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
