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

func (o *Surveyerrordetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyerrordetails
	
	return json.Marshal(&struct { 
		FlowDiagnosticInfo *Flowdiagnosticinfo `json:"flowDiagnosticInfo,omitempty"`
		
		SurveyErrorReason *string `json:"surveyErrorReason,omitempty"`
		*Alias
	}{ 
		FlowDiagnosticInfo: o.FlowDiagnosticInfo,
		
		SurveyErrorReason: o.SurveyErrorReason,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyerrordetails) UnmarshalJSON(b []byte) error {
	var SurveyerrordetailsMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyerrordetailsMap)
	if err != nil {
		return err
	}
	
	if FlowDiagnosticInfo, ok := SurveyerrordetailsMap["flowDiagnosticInfo"].(map[string]interface{}); ok {
		FlowDiagnosticInfoString, _ := json.Marshal(FlowDiagnosticInfo)
		json.Unmarshal(FlowDiagnosticInfoString, &o.FlowDiagnosticInfo)
	}
	
	if SurveyErrorReason, ok := SurveyerrordetailsMap["surveyErrorReason"].(string); ok {
		o.SurveyErrorReason = &SurveyErrorReason
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyerrordetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
