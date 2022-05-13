package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationscoringset
type Evaluationscoringset struct { 
	// TotalScore - Score of all questions
	TotalScore *float32 `json:"totalScore,omitempty"`


	// TotalCriticalScore - Score of only the critical questions
	TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`


	// TotalNonCriticalScore - Score of only the non-critical questions
	TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`


	// QuestionGroupScores
	QuestionGroupScores *[]Evaluationquestiongroupscore `json:"questionGroupScores,omitempty"`


	// AnyFailedKillQuestions - Indicates that at least one fatal question was answered without having the highest score available for the question
	AnyFailedKillQuestions *bool `json:"anyFailedKillQuestions,omitempty"`


	// Comments - Overall comments from the evaluator
	Comments *string `json:"comments,omitempty"`


	// AgentComments - Comments from the agent while reviewing evaluation results
	AgentComments *string `json:"agentComments,omitempty"`


	// TranscriptTopics - List of topics found within the conversation's transcripts
	TranscriptTopics *[]Transcripttopic `json:"transcriptTopics,omitempty"`

}

func (o *Evaluationscoringset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationscoringset
	
	return json.Marshal(&struct { 
		TotalScore *float32 `json:"totalScore,omitempty"`
		
		TotalCriticalScore *float32 `json:"totalCriticalScore,omitempty"`
		
		TotalNonCriticalScore *float32 `json:"totalNonCriticalScore,omitempty"`
		
		QuestionGroupScores *[]Evaluationquestiongroupscore `json:"questionGroupScores,omitempty"`
		
		AnyFailedKillQuestions *bool `json:"anyFailedKillQuestions,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		
		AgentComments *string `json:"agentComments,omitempty"`
		
		TranscriptTopics *[]Transcripttopic `json:"transcriptTopics,omitempty"`
		*Alias
	}{ 
		TotalScore: o.TotalScore,
		
		TotalCriticalScore: o.TotalCriticalScore,
		
		TotalNonCriticalScore: o.TotalNonCriticalScore,
		
		QuestionGroupScores: o.QuestionGroupScores,
		
		AnyFailedKillQuestions: o.AnyFailedKillQuestions,
		
		Comments: o.Comments,
		
		AgentComments: o.AgentComments,
		
		TranscriptTopics: o.TranscriptTopics,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationscoringset) UnmarshalJSON(b []byte) error {
	var EvaluationscoringsetMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationscoringsetMap)
	if err != nil {
		return err
	}
	
	if TotalScore, ok := EvaluationscoringsetMap["totalScore"].(float64); ok {
		TotalScoreFloat32 := float32(TotalScore)
		o.TotalScore = &TotalScoreFloat32
	}
	
	if TotalCriticalScore, ok := EvaluationscoringsetMap["totalCriticalScore"].(float64); ok {
		TotalCriticalScoreFloat32 := float32(TotalCriticalScore)
		o.TotalCriticalScore = &TotalCriticalScoreFloat32
	}
	
	if TotalNonCriticalScore, ok := EvaluationscoringsetMap["totalNonCriticalScore"].(float64); ok {
		TotalNonCriticalScoreFloat32 := float32(TotalNonCriticalScore)
		o.TotalNonCriticalScore = &TotalNonCriticalScoreFloat32
	}
	
	if QuestionGroupScores, ok := EvaluationscoringsetMap["questionGroupScores"].([]interface{}); ok {
		QuestionGroupScoresString, _ := json.Marshal(QuestionGroupScores)
		json.Unmarshal(QuestionGroupScoresString, &o.QuestionGroupScores)
	}
	
	if AnyFailedKillQuestions, ok := EvaluationscoringsetMap["anyFailedKillQuestions"].(bool); ok {
		o.AnyFailedKillQuestions = &AnyFailedKillQuestions
	}
    
	if Comments, ok := EvaluationscoringsetMap["comments"].(string); ok {
		o.Comments = &Comments
	}
    
	if AgentComments, ok := EvaluationscoringsetMap["agentComments"].(string); ok {
		o.AgentComments = &AgentComments
	}
    
	if TranscriptTopics, ok := EvaluationscoringsetMap["transcriptTopics"].([]interface{}); ok {
		TranscriptTopicsString, _ := json.Marshal(TranscriptTopics)
		json.Unmarshal(TranscriptTopicsString, &o.TranscriptTopics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationscoringset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
