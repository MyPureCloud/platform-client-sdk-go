package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsevaluation
type Analyticsevaluation struct { 
	// CalibrationId - The calibration ID used for the purpose of training evaluators
	CalibrationId *string `json:"calibrationId,omitempty"`


	// ContextId - A unique identifier for an evaluation form, regardless of version
	ContextId *string `json:"contextId,omitempty"`


	// Deleted - Whether the evaluation has been deleted
	Deleted *bool `json:"deleted,omitempty"`


	// EvaluationId - Unique identifier for the evaluation
	EvaluationId *string `json:"evaluationId,omitempty"`


	// EvaluatorId - A unique identifier of the user who evaluated the interaction
	EvaluatorId *string `json:"evaluatorId,omitempty"`


	// EventTime - Specifies when an evaluation occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// FormId - ID of the evaluation form used
	FormId *string `json:"formId,omitempty"`


	// FormName - Name of the evaluation form used
	FormName *string `json:"formName,omitempty"`


	// QueueId - The ID of the associated queue
	QueueId *string `json:"queueId,omitempty"`


	// Rescored - Whether the evaluation has been rescored at least once
	Rescored *bool `json:"rescored,omitempty"`


	// UserId - ID of the agent the evaluation was performed against
	UserId *string `json:"userId,omitempty"`


	// GetoTotalCriticalScore - Total critical score of the evaluation
	GetoTotalCriticalScore *int `json:"getoTotalCriticalScore,omitempty"`


	// GetoTotalScore - Total score of the evaluation
	GetoTotalScore *int `json:"getoTotalScore,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsevaluation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
