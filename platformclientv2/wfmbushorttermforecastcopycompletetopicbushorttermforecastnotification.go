package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastcopycompletetopicbushorttermforecastnotification
type Wfmbushorttermforecastcopycompletetopicbushorttermforecastnotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// Result
	Result *Wfmbushorttermforecastcopycompletetopicbushorttermforecast `json:"result,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`

}

func (o *Wfmbushorttermforecastcopycompletetopicbushorttermforecastnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastcopycompletetopicbushorttermforecastnotification
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		Result *Wfmbushorttermforecastcopycompletetopicbushorttermforecast `json:"result,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		Result: o.Result,
		
		OperationId: o.OperationId,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbushorttermforecastcopycompletetopicbushorttermforecastnotification) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastcopycompletetopicbushorttermforecastnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastcopycompletetopicbushorttermforecastnotificationMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Result, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastnotificationMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if OperationId, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastnotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastcopycompletetopicbushorttermforecastnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
