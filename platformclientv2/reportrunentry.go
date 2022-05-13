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

func (o *Reportrunentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportrunentry
	
	RunTime := new(string)
	if o.RunTime != nil {
		
		*RunTime = timeutil.Strftime(o.RunTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		ReportId: o.ReportId,
		
		RunTime: RunTime,
		
		RunStatus: o.RunStatus,
		
		ErrorMessage: o.ErrorMessage,
		
		RunDurationMsec: o.RunDurationMsec,
		
		ReportUrl: o.ReportUrl,
		
		ReportFormat: o.ReportFormat,
		
		ScheduleUri: o.ScheduleUri,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Reportrunentry) UnmarshalJSON(b []byte) error {
	var ReportrunentryMap map[string]interface{}
	err := json.Unmarshal(b, &ReportrunentryMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReportrunentryMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ReportrunentryMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ReportId, ok := ReportrunentryMap["reportId"].(string); ok {
		o.ReportId = &ReportId
	}
    
	if runTimeString, ok := ReportrunentryMap["runTime"].(string); ok {
		RunTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", runTimeString)
		o.RunTime = &RunTime
	}
	
	if RunStatus, ok := ReportrunentryMap["runStatus"].(string); ok {
		o.RunStatus = &RunStatus
	}
    
	if ErrorMessage, ok := ReportrunentryMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if RunDurationMsec, ok := ReportrunentryMap["runDurationMsec"].(float64); ok {
		RunDurationMsecInt := int(RunDurationMsec)
		o.RunDurationMsec = &RunDurationMsecInt
	}
	
	if ReportUrl, ok := ReportrunentryMap["reportUrl"].(string); ok {
		o.ReportUrl = &ReportUrl
	}
    
	if ReportFormat, ok := ReportrunentryMap["reportFormat"].(string); ok {
		o.ReportFormat = &ReportFormat
	}
    
	if ScheduleUri, ok := ReportrunentryMap["scheduleUri"].(string); ok {
		o.ScheduleUri = &ScheduleUri
	}
    
	if SelfUri, ok := ReportrunentryMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Reportrunentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
