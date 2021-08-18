package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Modelingstatusresponse
type Modelingstatusresponse struct { 
	// Id - The ID generated for the modeling job.  Use to GET result when job is completed.
	Id *string `json:"id,omitempty"`


	// Status - The status of the modeling job.
	Status *string `json:"status,omitempty"`


	// ErrorDetails - If the request could not be properly processed, error details will be given here.
	ErrorDetails *[]Modelingprocessingerror `json:"errorDetails,omitempty"`


	// ModelingResultUri - The uri of the modeling result. It has a value if the status is either 'Success', 'PartialFailure', or 'Failed'.
	ModelingResultUri *string `json:"modelingResultUri,omitempty"`

}

func (u *Modelingstatusresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Modelingstatusresponse

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ErrorDetails *[]Modelingprocessingerror `json:"errorDetails,omitempty"`
		
		ModelingResultUri *string `json:"modelingResultUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Status: u.Status,
		
		ErrorDetails: u.ErrorDetails,
		
		ModelingResultUri: u.ModelingResultUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Modelingstatusresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
