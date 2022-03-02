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


	// AssistedAnswerId - AnswerId found with evaluation assistance conditions
	AssistedAnswerId *string `json:"assistedAnswerId,omitempty"`


	// FailedKillQuestion - Applicable only on fatal questions. Indicates that the answer selected was not the highest score available for the question
	FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`


	// Comments - Comments from the evaluator specific to this question
	Comments *string `json:"comments,omitempty"`

}

func (o *Evaluationquestionscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationquestionscore
	
	return json.Marshal(&struct { 
		QuestionId *string `json:"questionId,omitempty"`
		
		AnswerId *string `json:"answerId,omitempty"`
		
		Score *int `json:"score,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		AssistedAnswerId *string `json:"assistedAnswerId,omitempty"`
		
		FailedKillQuestion *bool `json:"failedKillQuestion,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		*Alias
	}{ 
		QuestionId: o.QuestionId,
		
		AnswerId: o.AnswerId,
		
		Score: o.Score,
		
		MarkedNA: o.MarkedNA,
		
		AssistedAnswerId: o.AssistedAnswerId,
		
		FailedKillQuestion: o.FailedKillQuestion,
		
		Comments: o.Comments,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationquestionscore) UnmarshalJSON(b []byte) error {
	var EvaluationquestionscoreMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationquestionscoreMap)
	if err != nil {
		return err
	}
	
	if QuestionId, ok := EvaluationquestionscoreMap["questionId"].(string); ok {
		o.QuestionId = &QuestionId
	}
	
	if AnswerId, ok := EvaluationquestionscoreMap["answerId"].(string); ok {
		o.AnswerId = &AnswerId
	}
	
	if Score, ok := EvaluationquestionscoreMap["score"].(float64); ok {
		ScoreInt := int(Score)
		o.Score = &ScoreInt
	}
	
	if MarkedNA, ok := EvaluationquestionscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
	
	if AssistedAnswerId, ok := EvaluationquestionscoreMap["assistedAnswerId"].(string); ok {
		o.AssistedAnswerId = &AssistedAnswerId
	}
	
	if FailedKillQuestion, ok := EvaluationquestionscoreMap["failedKillQuestion"].(bool); ok {
		o.FailedKillQuestion = &FailedKillQuestion
	}
	
	if Comments, ok := EvaluationquestionscoreMap["comments"].(string); ok {
		o.Comments = &Comments
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationquestionscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
