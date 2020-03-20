package platformclientv2
import (
	"encoding/json"
)

// Asyncqueryresponse
type Asyncqueryresponse struct { 
	// JobId - Unique identifier for the async query execution. Can be used to check the status of the query and retrieve results.
	JobId *string `json:"jobId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Asyncqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
