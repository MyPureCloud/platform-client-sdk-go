package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentquestionscore
type Assessmentquestionscore struct { 
	// FailedKillQuestion - True if this was a failed Kill question
	FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`


	// Comments - Comments provided for the answer
	Comments *string `json:"comments,omitempty"`


	// QuestionId - The ID of the question
	QuestionId *string `json:"questionId,omitempty"`


	// AnswerId - The ID of the selected answer
	AnswerId *string `json:"answerId,omitempty"`


	// Score - The score received for this question
	Score *int `json:"score,omitempty"`


	// MarkedNA - True if this question was marked as NA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// FreeTextAnswer - Answer for free text answer type
	FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`

}

func (u *Assessmentquestionscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assessmentquestionscore

	

	return json.Marshal(&struct { 
		FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		
		QuestionId *string `json:"questionId,omitempty"`
		
		AnswerId *string `json:"answerId,omitempty"`
		
		Score *int `json:"score,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		FreeTextAnswer *string `json:"freeTextAnswer,omitempty"`
		*Alias
	}{ 
		FailedKillQuestion: u.FailedKillQuestion,
		
		Comments: u.Comments,
		
		QuestionId: u.QuestionId,
		
		AnswerId: u.AnswerId,
		
		Score: u.Score,
		
		MarkedNA: u.MarkedNA,
		
		FreeTextAnswer: u.FreeTextAnswer,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Assessmentquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
