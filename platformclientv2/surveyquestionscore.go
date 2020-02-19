package platformclientv2
import (
	"encoding/json"
)

// Surveyquestionscore
type Surveyquestionscore struct { 
	// QuestionId
	QuestionId *string `json:"questionId,omitempty"`


	// AnswerId
	AnswerId *string `json:"answerId,omitempty"`


	// Score
	Score *int32 `json:"score,omitempty"`


	// MarkedNA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// NpsScore
	NpsScore *int32 `json:"npsScore,omitempty"`


	// NpsTextAnswer
	NpsTextAnswer *string `json:"npsTextAnswer,omitempty"`


	// FreeTextAnswer
	FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyquestionscore) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
