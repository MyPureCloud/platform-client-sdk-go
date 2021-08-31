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

func (o *Evaluationqualityv2topicevaluationscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationqualityv2topicevaluationscoringset
	
	return json.Marshal(&struct { 
		TotalScore *int `json:"totalScore,omitempty"`
		
		TotalCriticalScore *int `json:"totalCriticalScore,omitempty"`
		*Alias
	}{ 
		TotalScore: o.TotalScore,
		
		TotalCriticalScore: o.TotalCriticalScore,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationqualityv2topicevaluationscoringset) UnmarshalJSON(b []byte) error {
	var Evaluationqualityv2topicevaluationscoringsetMap map[string]interface{}
	err := json.Unmarshal(b, &Evaluationqualityv2topicevaluationscoringsetMap)
	if err != nil {
		return err
	}
	
	if TotalScore, ok := Evaluationqualityv2topicevaluationscoringsetMap["totalScore"].(float64); ok {
		TotalScoreInt := int(TotalScore)
		o.TotalScore = &TotalScoreInt
	}
	
	if TotalCriticalScore, ok := Evaluationqualityv2topicevaluationscoringsetMap["totalCriticalScore"].(float64); ok {
		TotalCriticalScoreInt := int(TotalCriticalScore)
		o.TotalCriticalScore = &TotalCriticalScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationqualityv2topicevaluationscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
