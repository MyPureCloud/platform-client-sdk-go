package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Reportrunentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportrunentry

	
	RunTime := new(string)
	if u.RunTime != nil {
		
		*RunTime = timeutil.Strftime(u.RunTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RunTime = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ReportId *string `json:"reportId,omitempty"`
		
		RunTime *string `json:"runTime,omitempty"`
		
		RunStatus *string `json:"runStatus,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		RunDurationMsec *int `json:"runDurationMsec,omitempty"`
		
		ReportUrl *string `json:"reportUrl,omitempty"`
		
		ReportFormat *string `json:"reportFormat,omitempty"`
		
		ScheduleUri *string `json:"scheduleUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ReportId: u.ReportId,
		
		RunTime: RunTime,
		
		RunStatus: u.RunStatus,
		
		ErrorMessage: u.ErrorMessage,
		
		RunDurationMsec: u.RunDurationMsec,
		
		ReportUrl: u.ReportUrl,
		
		ReportFormat: u.ReportFormat,
		
		ScheduleUri: u.ScheduleUri,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportrunentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
