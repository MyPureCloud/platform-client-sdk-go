package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduleruntopicbuschedulingrunprogressnotification
type Wfmbuscheduleruntopicbuschedulingrunprogressnotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// Result
	Result *Wfmbuscheduleruntopicbuschedulerun `json:"result,omitempty"`

}

func (o *Wfmbuscheduleruntopicbuschedulingrunprogressnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduleruntopicbuschedulingrunprogressnotification
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		Result *Wfmbuscheduleruntopicbuschedulerun `json:"result,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		OperationId: o.OperationId,
		
		Result: o.Result,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuscheduleruntopicbuschedulingrunprogressnotification) UnmarshalJSON(b []byte) error {
	var WfmbuscheduleruntopicbuschedulingrunprogressnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduleruntopicbuschedulingrunprogressnotificationMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmbuscheduleruntopicbuschedulingrunprogressnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if OperationId, ok := WfmbuscheduleruntopicbuschedulingrunprogressnotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
	
	if Result, ok := WfmbuscheduleruntopicbuschedulingrunprogressnotificationMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduleruntopicbuschedulingrunprogressnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
