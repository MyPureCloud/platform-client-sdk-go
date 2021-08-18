package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationscoringset
type Evaluationscoringset struct { 
	// TotalScore - Score of all questions
	TotalScore *float32 `json:"totalScore,omitempty"`


	// TotalCriticalScore - Score of only the critical questions
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`


	// TotalNonCriticalScore - Score of only the non-critical questions
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`


	// QuestionGroupScores
	QuestionGroupScores *[]Evaluationquestiongroupscore `json:"questionGroupScores,omitempty"`


	// AnyFailedKillQuestions - Indicates that at least one fatal question was answered without having the highest score available for the question
	AnyFailedKillQuestions *bool `json:"anyFailedKillQuestions,omitempty"`


	// Comments - Overall comments from the evaluator
	Comments *string `json:"comments,omitempty"`


	// AgentComments - Comments from the agent while reviewing evaluation results
	AgentComments *string `json:"agentComments,omitempty"`

}

func (u *Evaluationscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationscoringset

	

	return json.Marshal(&struct { 
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`
		
		TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`
		
		QuestionGroupScores *[]Evaluationquestiongroupscore `json:"questionGroupScores,omitempty"`
		
		AnyFailedKillQuestions *bool `json:"anyFailedKillQuestions,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		
		AgentComments *string `json:"agentComments,omitempty"`
		*Alias
	}{ 
		TotalScore: u.TotalScore,
		
		TotalCriticalScore: u.TotalCriticalScore,
		
		TotalNonCriticalScore: u.TotalNonCriticalScore,
		
		QuestionGroupScores: u.QuestionGroupScores,
		
		AnyFailedKillQuestions: u.AnyFailedKillQuestions,
		
		Comments: u.Comments,
		
		AgentComments: u.AgentComments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluationscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
