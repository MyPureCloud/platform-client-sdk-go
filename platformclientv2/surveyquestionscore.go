package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyquestionscore
type Surveyquestionscore struct { 
	// QuestionId
	QuestionId *string `json:"questionId,omitempty"`


	// AnswerId
	AnswerId *string `json:"answerId,omitempty"`


	// Score - Unweighted score of the question
	Score *int `json:"score,omitempty"`


	// MarkedNA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// AssistedAnswerId - AnswerId found with evaluation assistance conditions
	AssistedAnswerId *string `json:"assistedAnswerId,omitempty"`


	// NpsScore
	NpsScore *int `json:"npsScore,omitempty"`


	// NpsTextAnswer
	NpsTextAnswer *string `json:"npsTextAnswer,omitempty"`


	// FreeTextAnswer
	FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`

}

func (o *Surveyquestionscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyquestionscore
	
	return json.Marshal(&struct { 
		QuestionId *string `json:"questionId,omitempty"`
		
		AnswerId *string `json:"answerId,omitempty"`
		
		Score *int `json:"score,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		AssistedAnswerId *string `json:"assistedAnswerId,omitempty"`
		
		NpsScore *int `json:"npsScore,omitempty"`
		
		NpsTextAnswer *string `json:"npsTextAnswer,omitempty"`
		
		FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`
		*Alias
	}{ 
		QuestionId: o.QuestionId,
		
		AnswerId: o.AnswerId,
		
		Score: o.Score,
		
		MarkedNA: o.MarkedNA,
		
		AssistedAnswerId: o.AssistedAnswerId,
		
		NpsScore: o.NpsScore,
		
		NpsTextAnswer: o.NpsTextAnswer,
		
		FreeTextAnswer: o.FreeTextAnswer,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyquestionscore) UnmarshalJSON(b []byte) error {
	var SurveyquestionscoreMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyquestionscoreMap)
	if err != nil {
		return err
	}
	
	if QuestionId, ok := SurveyquestionscoreMap["questionId"].(string); ok {
		o.QuestionId = &QuestionId
	}
	
	if AnswerId, ok := SurveyquestionscoreMap["answerId"].(string); ok {
		o.AnswerId = &AnswerId
	}
	
	if Score, ok := SurveyquestionscoreMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	
	if MarkedNA, ok := SurveyquestionscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
	
	if AssistedAnswerId, ok := SurveyquestionscoreMap["assistedAnswerId"].(string); ok {
		o.AssistedAnswerId = &AssistedAnswerId
	}
	
	if NpsScore, ok := SurveyquestionscoreMap["npsScore"].(float64); ok {
		NpsScoreInt := int(NpsScore)
		o.NpsScore = &NpsScoreInt
	}
	
	if NpsTextAnswer, ok := SurveyquestionscoreMap["npsTextAnswer"].(string); ok {
		o.NpsTextAnswer = &NpsTextAnswer
	}
	
	if FreeTextAnswer, ok := SurveyquestionscoreMap["freeTextAnswer"].(string); ok {
		o.FreeTextAnswer = &FreeTextAnswer
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
