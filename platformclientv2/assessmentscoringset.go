package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentscoringset
type Assessmentscoringset struct { 
	// TotalScore - The total score of the answers
	TotalScore *float32 `json:"totalScore,omitempty"`


	// TotalCriticalScore - The total score for the critical questions
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`


	// TotalNonCriticalScore - The total score for the non-critical questions
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`


	// QuestionGroupScores - The individual scores for each question group
	QuestionGroupScores *[]Assessmentquestiongroupscore `json:"questionGroupScores,omitempty"`


	// FailureReasons - If the assessment was not passed, the reasons for failure.
	FailureReasons *[]string `json:"failureReasons,omitempty"`


	// Comments - Comments provided for these answers.
	Comments *string `json:"comments,omitempty"`


	// AgentComments - Comments provided by agent.
	AgentComments *string `json:"agentComments,omitempty"`


	// IsPassed - True if the assessment was passed
	IsPassed *bool `json:"isPassed,omitempty"`

}

func (o *Assessmentscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assessmentscoringset
	
	return json.Marshal(&struct { 
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`
		
		TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`
		
		QuestionGroupScores *[]Assessmentquestiongroupscore `json:"questionGroupScores,omitempty"`
		
		FailureReasons *[]string `json:"failureReasons,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		
		AgentComments *string `json:"agentComments,omitempty"`
		
		IsPassed *bool `json:"isPassed,omitempty"`
		*Alias
	}{ 
		TotalScore: o.TotalScore,
		
		TotalCriticalScore: o.TotalCriticalScore,
		
		TotalNonCriticalScore: o.TotalNonCriticalScore,
		
		QuestionGroupScores: o.QuestionGroupScores,
		
		FailureReasons: o.FailureReasons,
		
		Comments: o.Comments,
		
		AgentComments: o.AgentComments,
		
		IsPassed: o.IsPassed,
		Alias:    (*Alias)(o),
	})
}

func (o *Assessmentscoringset) UnmarshalJSON(b []byte) error {
	var AssessmentscoringsetMap map[string]interface{}
	err := json.Unmarshal(b, &AssessmentscoringsetMap)
	if err != nil {
		return err
	}
	
	if TotalScore, ok := AssessmentscoringsetMap["totalScore"].(float64); ok {
		TotalScoreFloat32 := float32(TotalScore)
		o.TotalScore = &TotalScoreFloat32
	}
	
	if TotalCriticalScore, ok := AssessmentscoringsetMap["totalCriticalScore"].(float64); ok {
		TotalCriticalScoreFloat32 := float32(TotalCriticalScore)
		o.TotalCriticalScore = &TotalCriticalScoreFloat32
	}
	
	if TotalNonCriticalScore, ok := AssessmentscoringsetMap["totalNonCriticalScore"].(float64); ok {
		TotalNonCriticalScoreFloat32 := float32(TotalNonCriticalScore)
		o.TotalNonCriticalScore = &TotalNonCriticalScoreFloat32
	}
	
	if QuestionGroupScores, ok := AssessmentscoringsetMap["questionGroupScores"].([]interface{}); ok {
		QuestionGroupScoresString, _ := json.Marshal(QuestionGroupScores)
		json.Unmarshal(QuestionGroupScoresString, &o.QuestionGroupScores)
	}
	
	if FailureReasons, ok := AssessmentscoringsetMap["failureReasons"].([]interface{}); ok {
		FailureReasonsString, _ := json.Marshal(FailureReasons)
		json.Unmarshal(FailureReasonsString, &o.FailureReasons)
	}
	
	if Comments, ok := AssessmentscoringsetMap["comments"].(string); ok {
		o.Comments = &Comments
	}
    
	if AgentComments, ok := AssessmentscoringsetMap["agentComments"].(string); ok {
		o.AgentComments = &AgentComments
	}
    
	if IsPassed, ok := AssessmentscoringsetMap["isPassed"].(bool); ok {
		o.IsPassed = &IsPassed
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Assessmentscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
