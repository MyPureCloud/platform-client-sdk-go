package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationquestiongroupscore
type Evaluationquestiongroupscore struct { 
	// QuestionGroupId
	QuestionGroupId *string `json:"questionGroupId,omitempty"`


	// TotalScore - Score of all questions in the group
	TotalScore *float32 `json:"totalScore,omitempty"`


	// MaxTotalScore - Maximum possible score of all questions in the group
	MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`


	// MarkedNA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// TotalCriticalScore - Score of only the critical questions in the group
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`


	// MaxTotalCriticalScore - Maximum possible score of only the critical questions in the group
	MaxTotalCriticalScore *float32 `json:"maxTotalCriticalScore,omitempty"`


	// TotalNonCriticalScore - Score of only the non critical questions in the group
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`


	// MaxTotalNonCriticalScore - Maximum possible score of only the non critical questions in the group
	MaxTotalNonCriticalScore *float32 `json:"maxTotalNonCriticalScore,omitempty"`


	// TotalScoreUnweighted - Unweighted score of all questions in the group
	TotalScoreUnweighted *float32 `json:"totalScoreUnweighted,omitempty"`


	// MaxTotalScoreUnweighted - Maximum possible unweighted score of all questions in the group
	MaxTotalScoreUnweighted *float32 `json:"maxTotalScoreUnweighted,omitempty"`


	// TotalCriticalScoreUnweighted - Unweighted score of only the critical questions in the group
	TotalCriticalScoreUnweighted *float32 `json:"totalCriticalScoreUnweighted,omitempty"`


	// MaxTotalCriticalScoreUnweighted - Maximum possible unweighted score of only the critical questions in the group
	MaxTotalCriticalScoreUnweighted *float32 `json:"maxTotalCriticalScoreUnweighted,omitempty"`


	// TotalNonCriticalScoreUnweighted - Unweighted score of only the non critical questions in the group
	TotalNonCriticalScoreUnweighted *float32 `json:"totalNonCriticalScoreUnweighted,omitempty"`


	// MaxTotalNonCriticalScoreUnweighted - Maximum possible unweighted score of only the non critical questions in the group
	MaxTotalNonCriticalScoreUnweighted *float32 `json:"maxTotalNonCriticalScoreUnweighted,omitempty"`


	// QuestionScores
	QuestionScores *[]Evaluationquestionscore `json:"questionScores,omitempty"`

}

func (u *Evaluationquestiongroupscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationquestiongroupscore

	

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
		
		QuestionScores *[]Evaluationquestionscore `json:"questionScores,omitempty"`
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
func (o *Evaluationquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
