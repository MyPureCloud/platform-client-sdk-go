package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Assessmentquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
