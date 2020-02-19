package platformclientv2
import (
	"encoding/json"
)

// Surveyscoringset
type Surveyscoringset struct { 
	// TotalScore
	TotalScore *float32 `json:"totalScore,omitempty"`


	// NpsScore
	NpsScore *int32 `json:"npsScore,omitempty"`


	// QuestionGroupScores
	QuestionGroupScores *[]Surveyquestiongroupscore `json:"questionGroupScores,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyscoringset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
