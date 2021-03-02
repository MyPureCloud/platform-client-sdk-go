package platformclientv2
import (
	"time"
	"encoding/json"
)

// Analyticsevaluation
type Analyticsevaluation struct { 
	// EvaluationId - Unique identifier for the evaluation
	EvaluationId *string `json:"evaluationId,omitempty"`


	// EvaluatorId - A unique identifier of the PureCloud user who evaluated the interaction
	EvaluatorId *string `json:"evaluatorId,omitempty"`


	// UserId - Unique identifier for the user being evaluated
	UserId *string `json:"userId,omitempty"`


	// EventTime - Specifies when an evaluation occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// QueueId - Unique identifier for the queue the conversation was on
	QueueId *string `json:"queueId,omitempty"`


	// FormId - Unique identifier for the form used to evaluate the conversation/agent
	FormId *string `json:"formId,omitempty"`


	// ContextId - A unique identifier for an evaluation form, regardless of version
	ContextId *string `json:"contextId,omitempty"`


	// FormName - Name of the evaluation form
	FormName *string `json:"formName,omitempty"`


	// CalibrationId - The calibration id used for the purpose of training evaluators
	CalibrationId *string `json:"calibrationId,omitempty"`


	// Rescored - Whether this evaluation has ever been rescored
	Rescored *bool `json:"rescored,omitempty"`


	// Deleted - Whether this evaluation has been deleted
	Deleted *bool `json:"deleted,omitempty"`


	// OTotalScore
	OTotalScore *int `json:"oTotalScore,omitempty"`


	// OTotalCriticalScore
	OTotalCriticalScore *int `json:"oTotalCriticalScore,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsevaluation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
