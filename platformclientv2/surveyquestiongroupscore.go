package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyquestiongroupscore
type Surveyquestiongroupscore struct { 
	// QuestionGroupId
	QuestionGroupId *string `json:"questionGroupId,omitempty"`


	// TotalScore - Score of all questions in the group
	TotalScore *float32 `json:"totalScore,omitempty"`


	// MaxTotalScore - Maximum possible score of all questions in the group
	MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`


	// MarkedNA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// QuestionScores
	QuestionScores *[]Surveyquestionscore `json:"questionScores,omitempty"`

}

func (u *Surveyquestiongroupscore) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyquestiongroupscore

	

	return json.Marshal(&struct { 
		QuestionGroupId *string `json:"questionGroupId,omitempty"`
		
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`
		
		MarkedNA *bool `json:"markedNA,omitempty"`
		
		QuestionScores *[]Surveyquestionscore `json:"questionScores,omitempty"`
		*Alias
	}{ 
		QuestionGroupId: u.QuestionGroupId,
		
		TotalScore: u.TotalScore,
		
		MaxTotalScore: u.MaxTotalScore,
		
		MarkedNA: u.MarkedNA,
		
		QuestionScores: u.QuestionScores,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Surveyquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
