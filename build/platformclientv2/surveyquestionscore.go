package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Surveyquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
