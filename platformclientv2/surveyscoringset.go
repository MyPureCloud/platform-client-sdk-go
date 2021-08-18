package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Surveyscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyscoringset

	

	return json.Marshal(&struct { 
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		NpsScore *int `json:"npsScore,omitempty"`
		
		QuestionGroupScores *[]Surveyquestiongroupscore `json:"questionGroupScores,omitempty"`
		*Alias
	}{ 
		TotalScore: u.TotalScore,
		
		NpsScore: u.NpsScore,
		
		QuestionGroupScores: u.QuestionGroupScores,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Surveyscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
