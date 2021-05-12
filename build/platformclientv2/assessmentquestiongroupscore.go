package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentquestiongroupscore
type Assessmentquestiongroupscore struct { 
	// QuestionGroupId - The ID of the question group
	QuestionGroupId *string `json:"questionGroupId,omitempty"`


	// TotalScore - The total score for the questions
	TotalScore *float32 `json:"totalScore,omitempty"`


	// MaxTotalScore - The maximum total score for the questions
	MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`


	// MarkedNA - True if this question group is marked NA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// TotalCriticalScore - The total score for the critical questions
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`


	// MaxTotalCriticalScore - The maximum total score for the critical questions
	MaxTotalCriticalScore *float32 `json:"maxTotalCriticalScore,omitempty"`


	// TotalNonCriticalScore - The total score for the non-critical questions
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`


	// MaxTotalNonCriticalScore - The maximum total score for the non-critical questions
	MaxTotalNonCriticalScore *float32 `json:"maxTotalNonCriticalScore,omitempty"`


	// TotalScoreUnweighted - The unweighted total score for this question group
	TotalScoreUnweighted *float32 `json:"totalScoreUnweighted,omitempty"`


	// MaxTotalScoreUnweighted - The maximum unweighted total score for this question group
	MaxTotalScoreUnweighted *float32 `json:"maxTotalScoreUnweighted,omitempty"`


	// TotalCriticalScoreUnweighted - The unweighted total score for the critical questions
	TotalCriticalScoreUnweighted *float32 `json:"totalCriticalScoreUnweighted,omitempty"`


	// MaxTotalCriticalScoreUnweighted - The maximum unweighted total score for the critical questions
	MaxTotalCriticalScoreUnweighted *float32 `json:"maxTotalCriticalScoreUnweighted,omitempty"`


	// TotalNonCriticalScoreUnweighted - The total unweighted score for the non-critical questions
	TotalNonCriticalScoreUnweighted *float32 `json:"totalNonCriticalScoreUnweighted,omitempty"`


	// MaxTotalNonCriticalScoreUnweighted - The maximum unweighted total score for the non-critical questions
	MaxTotalNonCriticalScoreUnweighted *float32 `json:"maxTotalNonCriticalScoreUnweighted,omitempty"`


	// QuestionScores - The individual question scores
	QuestionScores *[]Assessmentquestionscore `json:"questionScores,omitempty"`

}

// String returns a JSON representation of the model
func (o *Assessmentquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
