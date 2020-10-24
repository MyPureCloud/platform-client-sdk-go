package platformclientv2
import (
	"time"
	"encoding/json"
)

// Recordingjobsquery
type Recordingjobsquery struct { 
	// Action - Operation to perform bulk task
	Action *string `json:"action,omitempty"`


	// ActionDate - The date when the action will be performed. If the operation will cause the delete date of a recording to be older than the export date, the export date will be adjusted to the delete date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ActionDate *time.Time `json:"actionDate,omitempty"`


	// IntegrationId - IntegrationId to Access AWS S3 bucket for bulk recording exports. This field is required and used only for EXPORT action.
	IntegrationId *string `json:"integrationId,omitempty"`


	// IncludeScreenRecordings - Include Screen recordings for export action, default value = true 
	IncludeScreenRecordings *bool `json:"includeScreenRecordings,omitempty"`


	// ConversationQuery - Conversation Query. Note: After the recording is created, it might take up to 48 hours for the recording to be included in the submitted job query.
	ConversationQuery *Asyncconversationquery `json:"conversationQuery,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recordingjobsquery) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
