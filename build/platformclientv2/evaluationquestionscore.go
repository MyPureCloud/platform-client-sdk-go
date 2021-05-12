package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationquestionscore
type Evaluationquestionscore struct { 
	// QuestionId
	QuestionId *string `json:"questionId,omitempty"`


	// AnswerId
	AnswerId *string `json:"answerId,omitempty"`


	// Score - Unweighted score of the question
	Score *int `json:"score,omitempty"`


	// MarkedNA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// FailedKillQuestion - Applicable only on fatal questions. Indicates that the answer selected was not the highest score available for the question
	FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`


	// Comments - Comments from the evaluator specific to this question
	Comments *string `json:"comments,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
