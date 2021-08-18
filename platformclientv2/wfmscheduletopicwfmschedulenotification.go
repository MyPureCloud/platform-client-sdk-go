package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmscheduletopicwfmschedulenotification
type Wfmscheduletopicwfmschedulenotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// PercentComplete
	PercentComplete *int `json:"percentComplete,omitempty"`


	// EventType
	EventType *string `json:"eventType,omitempty"`

}

func (u *Wfmscheduletopicwfmschedulenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmscheduletopicwfmschedulenotification

	

	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		PercentComplete *int `json:"percentComplete,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		*Alias
	}{ 
		Status: u.Status,
		
		OperationId: u.OperationId,
		
		DownloadUrl: u.DownloadUrl,
		
		PercentComplete: u.PercentComplete,
		
		EventType: u.EventType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmscheduletopicwfmschedulenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
