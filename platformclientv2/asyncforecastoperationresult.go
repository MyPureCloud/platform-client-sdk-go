package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Asyncforecastoperationresult
type Asyncforecastoperationresult struct { 
	// Status - The status of the operation
	Status *string `json:"status,omitempty"`


	// OperationId - The ID for the operation
	OperationId *string `json:"operationId,omitempty"`


	// Result - The result of the operation.  Null unless status == Complete
	Result *Bushorttermforecast `json:"result,omitempty"`


	// Progress - Percent progress for the operation
	Progress *int `json:"progress,omitempty"`

}

func (o *Asyncforecastoperationresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Asyncforecastoperationresult
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		Result *Bushorttermforecast `json:"result,omitempty"`
		
		Progress *int `json:"progress,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		OperationId: o.OperationId,
		
		Result: o.Result,
		
		Progress: o.Progress,
		Alias:    (*Alias)(o),
	})
}

func (o *Asyncforecastoperationresult) UnmarshalJSON(b []byte) error {
	var AsyncforecastoperationresultMap map[string]interface{}
	err := json.Unmarshal(b, &AsyncforecastoperationresultMap)
	if err != nil {
		return err
	}
	
	if Status, ok := AsyncforecastoperationresultMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if OperationId, ok := AsyncforecastoperationresultMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if Result, ok := AsyncforecastoperationresultMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if Progress, ok := AsyncforecastoperationresultMap["progress"].(float64); ok {
		ProgressInt := int(Progress)
		o.Progress = &ProgressInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Asyncforecastoperationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
