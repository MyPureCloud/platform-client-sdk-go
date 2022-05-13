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

func (o *Evaluationquestiongroupscore) MarshalJSON() ([]byte, error) {
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
		QuestionGroupId: o.QuestionGroupId,
		
		TotalScore: o.TotalScore,
		
		MaxTotalScore: o.MaxTotalScore,
		
		MarkedNA: o.MarkedNA,
		
		TotalCriticalScore: o.TotalCriticalScore,
		
		MaxTotalCriticalScore: o.MaxTotalCriticalScore,
		
		TotalNonCriticalScore: o.TotalNonCriticalScore,
		
		MaxTotalNonCriticalScore: o.MaxTotalNonCriticalScore,
		
		TotalScoreUnweighted: o.TotalScoreUnweighted,
		
		MaxTotalScoreUnweighted: o.MaxTotalScoreUnweighted,
		
		TotalCriticalScoreUnweighted: o.TotalCriticalScoreUnweighted,
		
		MaxTotalCriticalScoreUnweighted: o.MaxTotalCriticalScoreUnweighted,
		
		TotalNonCriticalScoreUnweighted: o.TotalNonCriticalScoreUnweighted,
		
		MaxTotalNonCriticalScoreUnweighted: o.MaxTotalNonCriticalScoreUnweighted,
		
		QuestionScores: o.QuestionScores,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationquestiongroupscore) UnmarshalJSON(b []byte) error {
	var EvaluationquestiongroupscoreMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationquestiongroupscoreMap)
	if err != nil {
		return err
	}
	
	if QuestionGroupId, ok := EvaluationquestiongroupscoreMap["questionGroupId"].(string); ok {
		o.QuestionGroupId = &QuestionGroupId
	}
    
	if TotalScore, ok := EvaluationquestiongroupscoreMap["totalScore"].(float64); ok {
		TotalScoreFloat32 := float32(TotalScore)
		o.TotalScore = &TotalScoreFloat32
	}
	
	if MaxTotalScore, ok := EvaluationquestiongroupscoreMap["maxTotalScore"].(float64); ok {
		MaxTotalScoreFloat32 := float32(MaxTotalScore)
		o.MaxTotalScore = &MaxTotalScoreFloat32
	}
	
	if MarkedNA, ok := EvaluationquestiongroupscoreMap["markedNA"].(bool); ok {
		o.MarkedNA = &MarkedNA
	}
    
	if TotalCriticalScore, ok := EvaluationquestiongroupscoreMap["totalCriticalScore"].(float64); ok {
		TotalCriticalScoreFloat32 := float32(TotalCriticalScore)
		o.TotalCriticalScore = &TotalCriticalScoreFloat32
	}
	
	if MaxTotalCriticalScore, ok := EvaluationquestiongroupscoreMap["maxTotalCriticalScore"].(float64); ok {
		MaxTotalCriticalScoreFloat32 := float32(MaxTotalCriticalScore)
		o.MaxTotalCriticalScore = &MaxTotalCriticalScoreFloat32
	}
	
	if TotalNonCriticalScore, ok := EvaluationquestiongroupscoreMap["totalNonCriticalScore"].(float64); ok {
		TotalNonCriticalScoreFloat32 := float32(TotalNonCriticalScore)
		o.TotalNonCriticalScore = &TotalNonCriticalScoreFloat32
	}
	
	if MaxTotalNonCriticalScore, ok := EvaluationquestiongroupscoreMap["maxTotalNonCriticalScore"].(float64); ok {
		MaxTotalNonCriticalScoreFloat32 := float32(MaxTotalNonCriticalScore)
		o.MaxTotalNonCriticalScore = &MaxTotalNonCriticalScoreFloat32
	}
	
	if TotalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["totalScoreUnweighted"].(float64); ok {
		TotalScoreUnweightedFloat32 := float32(TotalScoreUnweighted)
		o.TotalScoreUnweighted = &TotalScoreUnweightedFloat32
	}
	
	if MaxTotalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["maxTotalScoreUnweighted"].(float64); ok {
		MaxTotalScoreUnweightedFloat32 := float32(MaxTotalScoreUnweighted)
		o.MaxTotalScoreUnweighted = &MaxTotalScoreUnweightedFloat32
	}
	
	if TotalCriticalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["totalCriticalScoreUnweighted"].(float64); ok {
		TotalCriticalScoreUnweightedFloat32 := float32(TotalCriticalScoreUnweighted)
		o.TotalCriticalScoreUnweighted = &TotalCriticalScoreUnweightedFloat32
	}
	
	if MaxTotalCriticalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["maxTotalCriticalScoreUnweighted"].(float64); ok {
		MaxTotalCriticalScoreUnweightedFloat32 := float32(MaxTotalCriticalScoreUnweighted)
		o.MaxTotalCriticalScoreUnweighted = &MaxTotalCriticalScoreUnweightedFloat32
	}
	
	if TotalNonCriticalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["totalNonCriticalScoreUnweighted"].(float64); ok {
		TotalNonCriticalScoreUnweightedFloat32 := float32(TotalNonCriticalScoreUnweighted)
		o.TotalNonCriticalScoreUnweighted = &TotalNonCriticalScoreUnweightedFloat32
	}
	
	if MaxTotalNonCriticalScoreUnweighted, ok := EvaluationquestiongroupscoreMap["maxTotalNonCriticalScoreUnweighted"].(float64); ok {
		MaxTotalNonCriticalScoreUnweightedFloat32 := float32(MaxTotalNonCriticalScoreUnweighted)
		o.MaxTotalNonCriticalScoreUnweighted = &MaxTotalNonCriticalScoreUnweightedFloat32
	}
	
	if QuestionScores, ok := EvaluationquestiongroupscoreMap["questionScores"].([]interface{}); ok {
		QuestionScoresString, _ := json.Marshal(QuestionScores)
		json.Unmarshal(QuestionScoresString, &o.QuestionScores)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationquestiongroupscore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
