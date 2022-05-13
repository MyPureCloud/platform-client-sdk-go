package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification
type Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// Result
	Result *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting `json:"result,omitempty"`


	// WeekDate
	WeekDate *time.Time `json:"weekDate,omitempty"`

}

func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		WeekDate = nil
	}
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		Result *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting `json:"result,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		OperationId: o.OperationId,
		
		Result: o.Result,
		
		WeekDate: WeekDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification) UnmarshalJSON(b []byte) error {
	var WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotificationMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if OperationId, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if Result, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotificationMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if weekDateString, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotificationMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", weekDateString)
		o.WeekDate = &WeekDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
