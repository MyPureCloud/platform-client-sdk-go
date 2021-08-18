package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyerrordetails
type Surveyerrordetails struct { 
	// FlowDiagnosticInfo - Additional information about any errors that occurred in the survey invite flow.
	FlowDiagnosticInfo *Flowdiagnosticinfo `json:"flowDiagnosticInfo,omitempty"`


	// SurveyErrorReason - An error code indicating the reason for the survey failure.
	SurveyErrorReason *string `json:"surveyErrorReason,omitempty"`

}

func (u *Surveyerrordetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyerrordetails

	

	return json.Marshal(&struct { 
		FlowDiagnosticInfo *Flowdiagnosticinfo `json:"flowDiagnosticInfo,omitempty"`
		
		SurveyErrorReason *string `json:"surveyErrorReason,omitempty"`
		*Alias
	}{ 
		FlowDiagnosticInfo: u.FlowDiagnosticInfo,
		
		SurveyErrorReason: u.SurveyErrorReason,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Surveyerrordetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
