package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyquestiongroupscore
type Surveyquestiongroupscore struct { 
	// QuestionGroupId
	QuestionGroupId *string `json:"questionGroupId,omitempty"`


	// TotalScore
	TotalScore *float32 `json:"totalScore,omitempty"`


	// MaxTotalScore
	MaxTotalScore *float32 `json:"maxTotalScore,omitempty"`


	// MarkedNA
	MarkedNA *bool `json:"markedNA,omitempty"`


	// QuestionScores
	QuestionScores *[]Surveyquestionscore `json:"questionScores,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
