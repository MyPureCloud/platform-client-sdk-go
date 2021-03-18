package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationquestiongroupscore
type Evaluationquestiongroupscore struct { 
	// QuestionGroupId
	QuestionGroupId *string `json:"questionGroupId,omitempty"`


	// TotalScore
	TotalScore *float32 `json:"totalScore,omitempty"`


	// MaxTotalScore
	MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`


	// MarkedNA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// TotalCriticalScore
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`


	// MaxTotalCriticalScore
	MaxTotalCriticalScore *float32 `json:"maxTotalCriticalScore,omitempty"`


	// TotalNonCriticalScore
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`


	// MaxTotalNonCriticalScore
	MaxTotalNonCriticalScore *float32 `json:"maxTotalNonCriticalScore,omitempty"`


	// TotalScoreUnweighted
	TotalScoreUnweighted *float32 `json:"totalScoreUnweighted,omitempty"`


	// MaxTotalScoreUnweighted
	MaxTotalScoreUnweighted *float32 `json:"maxTotalScoreUnweighted,omitempty"`


	// TotalCriticalScoreUnweighted
	TotalCriticalScoreUnweighted *float32 `json:"totalCriticalScoreUnweighted,omitempty"`


	// MaxTotalCriticalScoreUnweighted
	MaxTotalCriticalScoreUnweighted *float32 `json:"maxTotalCriticalScoreUnweighted,omitempty"`


	// TotalNonCriticalScoreUnweighted
	TotalNonCriticalScoreUnweighted *float32 `json:"totalNonCriticalScoreUnweighted,omitempty"`


	// MaxTotalNonCriticalScoreUnweighted
	MaxTotalNonCriticalScoreUnweighted *float32 `json:"maxTotalNonCriticalScoreUnweighted,omitempty"`


	// QuestionScores
	QuestionScores *[]Evaluationquestionscore `json:"questionScores,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
