package platformclientv2
import (
	"time"
	"encoding/json"
)

// Reportrunentry
type Reportrunentry struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ReportId
	ReportId *string `json:"reportId,omitempty"`


	// RunTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RunTime *time.Time `json:"runTime,omitempty"`


	// RunStatus
	RunStatus *string `json:"runStatus,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// RunDurationMsec
	RunDurationMsec *int `json:"runDurationMsec,omitempty"`


	// ReportUrl
	ReportUrl *string `json:"reportUrl,omitempty"`


	// ReportFormat
	ReportFormat *string `json:"reportFormat,omitempty"`


	// ScheduleUri
	ScheduleUri *string `json:"scheduleUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reportrunentry) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
