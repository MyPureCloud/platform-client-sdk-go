package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuintradaydataupdatetopicbuintradaynotification
type Wfmbuintradaydataupdatetopicbuintradaynotification struct { 
	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// Result
	Result *Wfmbuintradaydataupdatetopicbuintradayresult `json:"result,omitempty"`

}

func (o *Wfmbuintradaydataupdatetopicbuintradaynotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbuintradaynotification
	
	return json.Marshal(&struct { 
		OperationId *string `json:"operationId,omitempty"`
		
		Result *Wfmbuintradaydataupdatetopicbuintradayresult `json:"result,omitempty"`
		*Alias
	}{ 
		OperationId: o.OperationId,
		
		Result: o.Result,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuintradaydataupdatetopicbuintradaynotification) UnmarshalJSON(b []byte) error {
	var WfmbuintradaydataupdatetopicbuintradaynotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuintradaydataupdatetopicbuintradaynotificationMap)
	if err != nil {
		return err
	}
	
	if OperationId, ok := WfmbuintradaydataupdatetopicbuintradaynotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
	
	if Result, ok := WfmbuintradaydataupdatetopicbuintradaynotificationMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradaynotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
