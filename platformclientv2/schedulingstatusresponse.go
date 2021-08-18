package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	PercentComplete *int `json:"percentComplete,omitempty"`

}

func (u *Schedulingstatusresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulingstatusresponse

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ErrorDetails *[]Schedulingprocessingerror `json:"errorDetails,omitempty"`
		
		SchedulingResultUri *string `json:"schedulingResultUri,omitempty"`
		
		PercentComplete *int `json:"percentComplete,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Status: u.Status,
		
		ErrorDetails: u.ErrorDetails,
		
		SchedulingResultUri: u.SchedulingResultUri,
		
		PercentComplete: u.PercentComplete,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Schedulingstatusresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
