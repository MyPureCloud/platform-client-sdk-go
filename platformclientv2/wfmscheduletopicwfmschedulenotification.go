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

func (o *Wfmscheduletopicwfmschedulenotification) MarshalJSON() ([]byte, error) {
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
		Status: o.Status,
		
		OperationId: o.OperationId,
		
		DownloadUrl: o.DownloadUrl,
		
		PercentComplete: o.PercentComplete,
		
		EventType: o.EventType,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmscheduletopicwfmschedulenotification) UnmarshalJSON(b []byte) error {
	var WfmscheduletopicwfmschedulenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmscheduletopicwfmschedulenotificationMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmscheduletopicwfmschedulenotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if OperationId, ok := WfmscheduletopicwfmschedulenotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if DownloadUrl, ok := WfmscheduletopicwfmschedulenotificationMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if PercentComplete, ok := WfmscheduletopicwfmschedulenotificationMap["percentComplete"].(float64); ok {
		PercentCompleteInt := int(PercentComplete)
		o.PercentComplete = &PercentCompleteInt
	}
	
	if EventType, ok := WfmscheduletopicwfmschedulenotificationMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmscheduletopicwfmschedulenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
