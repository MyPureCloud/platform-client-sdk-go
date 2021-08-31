package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingdataexporttopicdataexportnotification
type Reportingdataexporttopicdataexportnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// RunId
	RunId *string `json:"runId,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// ExportFormat
	ExportFormat *string `json:"exportFormat,omitempty"`


	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// ViewType
	ViewType *string `json:"viewType,omitempty"`


	// ExportErrorMessagesType
	ExportErrorMessagesType *string `json:"exportErrorMessagesType,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// CreatedDateTime
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`


	// ModifiedDateTime
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`


	// PercentageComplete
	PercentageComplete *float32 `json:"percentageComplete,omitempty"`


	// EmailStatuses
	EmailStatuses *map[string]string `json:"emailStatuses,omitempty"`


	// EmailErrorDescription
	EmailErrorDescription *string `json:"emailErrorDescription,omitempty"`


	// ScheduleExpression
	ScheduleExpression *string `json:"scheduleExpression,omitempty"`

}

func (o *Reportingdataexporttopicdataexportnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingdataexporttopicdataexportnotification
	
	CreatedDateTime := new(string)
	if o.CreatedDateTime != nil {
		
		*CreatedDateTime = timeutil.Strftime(o.CreatedDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDateTime = nil
	}
	
	ModifiedDateTime := new(string)
	if o.ModifiedDateTime != nil {
		
		*ModifiedDateTime = timeutil.Strftime(o.ModifiedDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDateTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		RunId *string `json:"runId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ExportFormat *string `json:"exportFormat,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		ViewType *string `json:"viewType,omitempty"`
		
		ExportErrorMessagesType *string `json:"exportErrorMessagesType,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		CreatedDateTime *string `json:"createdDateTime,omitempty"`
		
		ModifiedDateTime *string `json:"modifiedDateTime,omitempty"`
		
		PercentageComplete *float32 `json:"percentageComplete,omitempty"`
		
		EmailStatuses *map[string]string `json:"emailStatuses,omitempty"`
		
		EmailErrorDescription *string `json:"emailErrorDescription,omitempty"`
		
		ScheduleExpression *string `json:"scheduleExpression,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		RunId: o.RunId,
		
		Name: o.Name,
		
		Status: o.Status,
		
		ExportFormat: o.ExportFormat,
		
		DownloadUrl: o.DownloadUrl,
		
		ViewType: o.ViewType,
		
		ExportErrorMessagesType: o.ExportErrorMessagesType,
		
		Read: o.Read,
		
		CreatedDateTime: CreatedDateTime,
		
		ModifiedDateTime: ModifiedDateTime,
		
		PercentageComplete: o.PercentageComplete,
		
		EmailStatuses: o.EmailStatuses,
		
		EmailErrorDescription: o.EmailErrorDescription,
		
		ScheduleExpression: o.ScheduleExpression,
		Alias:    (*Alias)(o),
	})
}

func (o *Reportingdataexporttopicdataexportnotification) UnmarshalJSON(b []byte) error {
	var ReportingdataexporttopicdataexportnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingdataexporttopicdataexportnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReportingdataexporttopicdataexportnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if RunId, ok := ReportingdataexporttopicdataexportnotificationMap["runId"].(string); ok {
		o.RunId = &RunId
	}
	
	if Name, ok := ReportingdataexporttopicdataexportnotificationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Status, ok := ReportingdataexporttopicdataexportnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if ExportFormat, ok := ReportingdataexporttopicdataexportnotificationMap["exportFormat"].(string); ok {
		o.ExportFormat = &ExportFormat
	}
	
	if DownloadUrl, ok := ReportingdataexporttopicdataexportnotificationMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
	
	if ViewType, ok := ReportingdataexporttopicdataexportnotificationMap["viewType"].(string); ok {
		o.ViewType = &ViewType
	}
	
	if ExportErrorMessagesType, ok := ReportingdataexporttopicdataexportnotificationMap["exportErrorMessagesType"].(string); ok {
		o.ExportErrorMessagesType = &ExportErrorMessagesType
	}
	
	if Read, ok := ReportingdataexporttopicdataexportnotificationMap["read"].(bool); ok {
		o.Read = &Read
	}
	
	if createdDateTimeString, ok := ReportingdataexporttopicdataexportnotificationMap["createdDateTime"].(string); ok {
		CreatedDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateTimeString)
		o.CreatedDateTime = &CreatedDateTime
	}
	
	if modifiedDateTimeString, ok := ReportingdataexporttopicdataexportnotificationMap["modifiedDateTime"].(string); ok {
		ModifiedDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateTimeString)
		o.ModifiedDateTime = &ModifiedDateTime
	}
	
	if PercentageComplete, ok := ReportingdataexporttopicdataexportnotificationMap["percentageComplete"].(float64); ok {
		PercentageCompleteFloat32 := float32(PercentageComplete)
		o.PercentageComplete = &PercentageCompleteFloat32
	}
	
	if EmailStatuses, ok := ReportingdataexporttopicdataexportnotificationMap["emailStatuses"].(map[string]interface{}); ok {
		EmailStatusesString, _ := json.Marshal(EmailStatuses)
		json.Unmarshal(EmailStatusesString, &o.EmailStatuses)
	}
	
	if EmailErrorDescription, ok := ReportingdataexporttopicdataexportnotificationMap["emailErrorDescription"].(string); ok {
		o.EmailErrorDescription = &EmailErrorDescription
	}
	
	if ScheduleExpression, ok := ReportingdataexporttopicdataexportnotificationMap["scheduleExpression"].(string); ok {
		o.ScheduleExpression = &ScheduleExpression
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingdataexporttopicdataexportnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
