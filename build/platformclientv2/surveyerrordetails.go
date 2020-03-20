package platformclientv2
import (
	"encoding/json"
)

// Surveyerrordetails
type Surveyerrordetails struct { 
	// FlowDiagnosticInfo - Additional information about any errors that occurred in the survey invite flow.
	FlowDiagnosticInfo *Flowdiagnosticinfo `json:"flowDiagnosticInfo,omitempty"`


	// SurveyErrorReason - An error code indicating the reason for the survey failure.
	SurveyErrorReason *string `json:"surveyErrorReason,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyerrordetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
