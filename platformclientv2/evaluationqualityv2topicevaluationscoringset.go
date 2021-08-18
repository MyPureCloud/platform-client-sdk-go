package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationqualityv2topicevaluationscoringset
type Evaluationqualityv2topicevaluationscoringset struct { 
	// TotalScore
	TotalScore *int `json:"totalScore,omitempty"`


	// TotalCriticalScore
	TotalCriticalScore *int `json:"totalCriticalScore,omitempty"`

}

func (u *Evaluationqualityv2topicevaluationscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationqualityv2topicevaluationscoringset

	

	return json.Marshal(&struct { 
		TotalScore *int `json:"totalScore,omitempty"`
		
		TotalCriticalScore *int `json:"totalCriticalScore,omitempty"`
		*Alias
	}{ 
		TotalScore: u.TotalScore,
		
		TotalCriticalScore: u.TotalCriticalScore,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluationqualityv2topicevaluationscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
