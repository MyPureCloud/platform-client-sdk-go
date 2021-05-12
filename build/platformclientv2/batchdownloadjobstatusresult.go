package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Batchdownloadjobstatusresult
type Batchdownloadjobstatusresult struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// JobId - JobId returned when job was initially submitted
	JobId *string `json:"jobId,omitempty"`


	// ExpectedResultCount - Number of results expected when job is completed
	ExpectedResultCount *int `json:"expectedResultCount,omitempty"`


	// ResultCount - Current number of results available
	ResultCount *int `json:"resultCount,omitempty"`


	// ErrorCount - Number of error results produced so far
	ErrorCount *int `json:"errorCount,omitempty"`


	// Results - Current set of results for the job
	Results *[]Batchdownloadjobresult `json:"results,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Batchdownloadjobstatusresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
