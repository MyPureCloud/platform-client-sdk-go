package platformclientv2
import (
	"time"
	"encoding/json"
)

// Calibrationcreate
type Calibrationcreate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Calibrator
	Calibrator *User `json:"calibrator,omitempty"`


	// Agent
	Agent *User `json:"agent,omitempty"`


	// Conversation - The conversation to use for the calibration.
	Conversation *Conversation `json:"conversation,omitempty"`


	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// AverageScore
	AverageScore *int `json:"averageScore,omitempty"`


	// HighScore
	HighScore *int `json:"highScore,omitempty"`


	// LowScore
	LowScore *int `json:"lowScore,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// Evaluations
	Evaluations *[]Evaluation `json:"evaluations,omitempty"`


	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`


	// ScoringIndex
	ScoringIndex *Evaluation `json:"scoringIndex,omitempty"`


	// ExpertEvaluator
	ExpertEvaluator *User `json:"expertEvaluator,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Calibrationcreate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
