package platformclientv2
import (
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

}

func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		Result *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting `json:"result,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		OperationId: o.OperationId,
		
		Result: o.Result,
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
