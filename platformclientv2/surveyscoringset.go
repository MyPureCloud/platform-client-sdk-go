package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyscoringset
type Surveyscoringset struct { 
	// TotalScore
	TotalScore *float32 `json:"totalScore,omitempty"`


	// NpsScore
	NpsScore *int `json:"npsScore,omitempty"`


	// QuestionGroupScores
	QuestionGroupScores *[]Surveyquestiongroupscore `json:"questionGroupScores,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
