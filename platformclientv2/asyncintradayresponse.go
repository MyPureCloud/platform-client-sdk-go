package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Asyncintradayresponse
type Asyncintradayresponse struct { 
	// Status - The status of the operation
	Status *string `json:"status,omitempty"`


	// OperationId - The ID for the operation
	OperationId *string `json:"operationId,omitempty"`


	// Result - The result of the operation.  Null unless status == Complete
	Result *Buintradayresponse `json:"result,omitempty"`

}

func (o *Asyncintradayresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Asyncintradayresponse
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		Result *Buintradayresponse `json:"result,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		OperationId: o.OperationId,
		
		Result: o.Result,
		Alias:    (*Alias)(o),
	})
}

func (o *Asyncintradayresponse) UnmarshalJSON(b []byte) error {
	var AsyncintradayresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AsyncintradayresponseMap)
	if err != nil {
		return err
	}
	
	if Status, ok := AsyncintradayresponseMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if OperationId, ok := AsyncintradayresponseMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
	
	if Result, ok := AsyncintradayresponseMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Asyncintradayresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
