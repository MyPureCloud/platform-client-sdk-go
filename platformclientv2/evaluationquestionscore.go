package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Evaluationquestionscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationquestionscore

	

	return json.Marshal(&struct { 
		QuestionId *string `json:"questionId,omitempty"`
		
		AnswerId *string `json:"answerId,omitempty"`
		
		Score *int `json:"score,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		*Alias
	}{ 
		QuestionId: u.QuestionId,
		
		AnswerId: u.AnswerId,
		
		Score: u.Score,
		
		MarkedNA: u.MarkedNA,
		
		FailedKillQuestion: u.FailedKillQuestion,
		
		Comments: u.Comments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluationquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
