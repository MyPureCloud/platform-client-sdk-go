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


	// NpsScore
	NpsScore *int `json:"npsScore,omitempty"`


	// NpsTextAnswer
	NpsTextAnswer *string `json:"npsTextAnswer,omitempty"`


	// FreeTextAnswer
	FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`

}

func (u *Surveyquestionscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyquestionscore

	

	return json.Marshal(&struct { 
		QuestionId *string `json:"questionId,omitempty"`
		
		AnswerId *string `json:"answerId,omitempty"`
		
		Score *int `json:"score,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		NpsScore *int `json:"npsScore,omitempty"`
		
		NpsTextAnswer *string `json:"npsTextAnswer,omitempty"`
		
		FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`
		*Alias
	}{ 
		QuestionId: u.QuestionId,
		
		AnswerId: u.AnswerId,
		
		Score: u.Score,
		
		MarkedNA: u.MarkedNA,
		
		NpsScore: u.NpsScore,
		
		NpsTextAnswer: u.NpsTextAnswer,
		
		FreeTextAnswer: u.FreeTextAnswer,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Surveyquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
