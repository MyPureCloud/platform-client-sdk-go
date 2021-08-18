package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Assessmentquestiongroupscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assessmentquestiongroupscore

	

	return json.Marshal(&struct { 
		QuestionGroupId *string `json:"questionGroupId,omitempty"`
		
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`
		
		MaxTotalCriticalScore *float32 `json:"maxTotalCriticalScore,omitempty"`
		
		TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`
		
		MaxTotalNonCriticalScore *float32 `json:"maxTotalNonCriticalScore,omitempty"`
		
		TotalScoreUnweighted *float32 `json:"totalScoreUnweighted,omitempty"`
		
		MaxTotalScoreUnweighted *float32 `json:"maxTotalScoreUnweighted,omitempty"`
		
		TotalCriticalScoreUnweighted *float32 `json:"totalCriticalScoreUnweighted,omitempty"`
		
		MaxTotalCriticalScoreUnweighted *float32 `json:"maxTotalCriticalScoreUnweighted,omitempty"`
		
		TotalNonCriticalScoreUnweighted *float32 `json:"totalNonCriticalScoreUnweighted,omitempty"`
		
		MaxTotalNonCriticalScoreUnweighted *float32 `json:"maxTotalNonCriticalScoreUnweighted,omitempty"`
		
		QuestionScores *[]Assessmentquestionscore `json:"questionScores,omitempty"`
		*Alias
	}{ 
		QuestionGroupId: u.QuestionGroupId,
		
		TotalScore: u.TotalScore,
		
		MaxTotalScore: u.MaxTotalScore,
		
		MarkedNA: u.MarkedNA,
		
		TotalCriticalScore: u.TotalCriticalScore,
		
		MaxTotalCriticalScore: u.MaxTotalCriticalScore,
		
		TotalNonCriticalScore: u.TotalNonCriticalScore,
		
		MaxTotalNonCriticalScore: u.MaxTotalNonCriticalScore,
		
		TotalScoreUnweighted: u.TotalScoreUnweighted,
		
		MaxTotalScoreUnweighted: u.MaxTotalScoreUnweighted,
		
		TotalCriticalScoreUnweighted: u.TotalCriticalScoreUnweighted,
		
		MaxTotalCriticalScoreUnweighted: u.MaxTotalCriticalScoreUnweighted,
		
		TotalNonCriticalScoreUnweighted: u.TotalNonCriticalScoreUnweighted,
		
		MaxTotalNonCriticalScoreUnweighted: u.MaxTotalNonCriticalScoreUnweighted,
		
		QuestionScores: u.QuestionScores,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Assessmentquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
