package platformclientv2
import (
	"time"
	"encoding/json"
)

// Calibration
type Calibration struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Calibrator
	Calibrator *User `json:"calibrator,omitempty"`


	// Agent
	Agent *User `json:"agent,omitempty"`


	// Conversation
	Conversation *Conversation `json:"conversation,omitempty"`


	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// AverageScore
	AverageScore *int32 `json:"averageScore,omitempty"`


	// HighScore
	HighScore *int32 `json:"highScore,omitempty"`


	// LowScore
	LowScore *int32 `json:"lowScore,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// Evaluations
	Evaluations *[]Evaluation `json:"evaluations,omitempty"`


	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`


	// ScoringIndex
	ScoringIndex **Evaluation `json:"scoringIndex,omitempty"`


	// ExpertEvaluator
	ExpertEvaluator *User `json:"expertEvaluator,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Calibration) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
