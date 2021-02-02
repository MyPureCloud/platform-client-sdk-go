package platformclientv2
import (
	"encoding/json"
)

// Evaluationquestionscore
type Evaluationquestionscore struct { 
	// QuestionId
	QuestionId *string `json:"questionId,omitempty"`


	// AnswerId
	AnswerId *string `json:"answerId,omitempty"`


	// Score
	Score *int `json:"score,omitempty"`


	// MarkedNA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// FailedKillQuestion
	FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`


	// Comments
	Comments *string `json:"comments,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationquestionscore) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
