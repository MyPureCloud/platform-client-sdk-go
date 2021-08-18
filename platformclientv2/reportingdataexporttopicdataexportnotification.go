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

func (u *Reportingdataexporttopicdataexportnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingdataexporttopicdataexportnotification

	
	CreatedDateTime := new(string)
	if u.CreatedDateTime != nil {
		
		*CreatedDateTime = timeutil.Strftime(u.CreatedDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDateTime = nil
	}
	
	ModifiedDateTime := new(string)
	if u.ModifiedDateTime != nil {
		
		*ModifiedDateTime = timeutil.Strftime(u.ModifiedDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		RunId: u.RunId,
		
		Name: u.Name,
		
		Status: u.Status,
		
		ExportFormat: u.ExportFormat,
		
		DownloadUrl: u.DownloadUrl,
		
		ViewType: u.ViewType,
		
		ExportErrorMessagesType: u.ExportErrorMessagesType,
		
		Read: u.Read,
		
		CreatedDateTime: CreatedDateTime,
		
		ModifiedDateTime: ModifiedDateTime,
		
		PercentageComplete: u.PercentageComplete,
		
		EmailStatuses: u.EmailStatuses,
		
		EmailErrorDescription: u.EmailErrorDescription,
		
		ScheduleExpression: u.ScheduleExpression,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportingdataexporttopicdataexportnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
