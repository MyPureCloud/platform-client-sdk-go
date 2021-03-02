package platformclientv2
import (
	"time"
	"encoding/json"
)

// Asyncquerystatus
type Asyncquerystatus struct { 
	// State - The current state of the asynchronous query
	State *string `json:"state,omitempty"`


	// ErrorMessage - The error associated with the current query, if the state is FAILED
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ExpirationDate - The time at which results for this query will expire. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`


	// SubmissionDate - The time at which the query was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SubmissionDate *time.Time `json:"submissionDate,omitempty"`


	// CompletionDate - The time at which the query completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CompletionDate *time.Time `json:"completionDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Asyncquerystatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
