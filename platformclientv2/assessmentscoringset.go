package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentscoringset
type Assessmentscoringset struct { 
	// TotalScore - The total score of the answers
	TotalScore *float32 `json:"totalScore,omitempty"`


	// TotalCriticalScore - The total score for the critical questions
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`


	// TotalNonCriticalScore - The total score for the non-critical questions
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`


	// QuestionGroupScores - The individual scores for each question group
	QuestionGroupScores *[]Assessmentquestiongroupscore `json:"questionGroupScores,omitempty"`


	// FailureReasons - If the assessment was not passed, the reasons for failure.
	FailureReasons *[]string `json:"failureReasons,omitempty"`


	// Comments - Comments provided for these answers.
	Comments *string `json:"comments,omitempty"`


	// AgentComments - Comments provided by agent.
	AgentComments *string `json:"agentComments,omitempty"`


	// IsPassed - True if the assessment was passed
	IsPassed *bool `json:"isPassed,omitempty"`

}

// String returns a JSON representation of the model
func (o *Assessmentscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
