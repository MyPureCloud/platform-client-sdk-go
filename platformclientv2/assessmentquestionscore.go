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

func (o *Assessmentquestionscore) MarshalJSON() ([]byte, error) {
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
		FailedKillQuestion: o.FailedKillQuestion,
		
		Comments: o.Comments,
		
		QuestionId: o.QuestionId,
		
		AnswerId: o.AnswerId,
		
		Score: o.Score,
		
		MarkedNA: o.MarkedNA,
		
		FreeTextAnswer: o.FreeTextAnswer,
		Alias:    (*Alias)(o),
	})
}

func (o *Assessmentquestionscore) UnmarshalJSON(b []byte) error {
	var AssessmentquestionscoreMap map[string]interface{}
	err := json.Unmarshal(b, &AssessmentquestionscoreMap)
	if err != nil {
		return err
	}
	
	if FailedKillQuestion, ok := AssessmentquestionscoreMap["failedKillQuestion"].(bool); ok {
		o.FailedKillQuestion = &FailedKillQuestion
	}
    
	if Comments, ok := AssessmentquestionscoreMap["comments"].(string); ok {
		o.Comments = &Comments
	}
    
	if QuestionId, ok := AssessmentquestionscoreMap["questionId"].(string); ok {
		o.QuestionId = &QuestionId
	}
    
	if AnswerId, ok := AssessmentquestionscoreMap["answerId"].(string); ok {
		o.AnswerId = &AnswerId
	}
    
	if Score, ok := AssessmentquestionscoreMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	
	if MarkedNA, ok := AssessmentquestionscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
    
	if FreeTextAnswer, ok := AssessmentquestionscoreMap["freeTextAnswer"].(string); ok {
		o.FreeTextAnswer = &FreeTextAnswer
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Assessmentquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
