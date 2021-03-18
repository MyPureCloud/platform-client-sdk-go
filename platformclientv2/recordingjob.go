package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingjob
type Recordingjob struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// State - The current state of the job.
	State *string `json:"state,omitempty"`


	// RecordingJobsQuery - Original query of the job.
	RecordingJobsQuery *Recordingjobsquery `json:"recordingJobsQuery,omitempty"`


	// DateCreated - Date when the job was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// TotalConversations - Total number of conversations affected.
	TotalConversations *int `json:"totalConversations,omitempty"`


	// TotalRecordings - Total number of recordings affected.
	TotalRecordings *int `json:"totalRecordings,omitempty"`


	// TotalProcessedRecordings - Total number of recordings have been processed.
	TotalProcessedRecordings *int `json:"totalProcessedRecordings,omitempty"`


	// PercentProgress - Progress in percentage based on the number of recordings
	PercentProgress *int `json:"percentProgress,omitempty"`


	// ErrorMessage - Error occurred during the job execution
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// User - Details of the user created the job
	User *Addressableentityref `json:"user,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recordingjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
