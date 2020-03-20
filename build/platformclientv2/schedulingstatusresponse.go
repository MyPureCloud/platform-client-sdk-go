package platformclientv2
import (
	"encoding/json"
)

// Schedulingstatusresponse
type Schedulingstatusresponse struct { 
	// Id - The ID generated for the scheduling job.  Use to GET result when job is completed.
	Id *string `json:"id,omitempty"`


	// Status - The status of the scheduling job.
	Status *string `json:"status,omitempty"`


	// ErrorDetails - If the request could not be properly processed, error details will be given here.
	ErrorDetails *[]Schedulingprocessingerror `json:"errorDetails,omitempty"`


	// SchedulingResultUri - The uri of the scheduling result. It has a value if the status is 'Success'.
	SchedulingResultUri *string `json:"schedulingResultUri,omitempty"`


	// PercentComplete - The percentage of the job that is complete.
	PercentComplete *int32 `json:"percentComplete,omitempty"`

}

// String returns a JSON representation of the model
func (o *Schedulingstatusresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
