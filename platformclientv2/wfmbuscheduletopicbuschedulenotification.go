package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbuschedulenotification
type Wfmbuscheduletopicbuschedulenotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// EventType
	EventType *string `json:"eventType,omitempty"`


	// Result
	Result *Wfmbuscheduletopicbuschedulemetadata `json:"result,omitempty"`

}

func (o *Wfmbuscheduletopicbuschedulenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicbuschedulenotification
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		Result *Wfmbuscheduletopicbuschedulemetadata `json:"result,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		OperationId: o.OperationId,
		
		EventType: o.EventType,
		
		Result: o.Result,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuscheduletopicbuschedulenotification) UnmarshalJSON(b []byte) error {
	var WfmbuscheduletopicbuschedulenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduletopicbuschedulenotificationMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmbuscheduletopicbuschedulenotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if OperationId, ok := WfmbuscheduletopicbuschedulenotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if EventType, ok := WfmbuscheduletopicbuschedulenotificationMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if Result, ok := WfmbuscheduletopicbuschedulenotificationMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbuschedulenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
