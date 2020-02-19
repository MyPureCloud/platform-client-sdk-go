package platformclientv2
import (
	"encoding/json"
)

// Evaluationscoringset
type Evaluationscoringset struct { 
	// TotalScore
	TotalScore *float32 `json:"totalScore,omitempty"`


	// TotalCriticalScore
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`


	// QuestionGroupScores
	QuestionGroupScores *[]Evaluationquestiongroupscore `json:"questionGroupScores,omitempty"`


	// AnyFailedKillQuestions
	AnyFailedKillQuestions *bool `json:"anyFailedKillQuestions,omitempty"`


	// Comments
	Comments *string `json:"comments,omitempty"`


	// AgentComments
	AgentComments *string `json:"agentComments,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationscoringset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
