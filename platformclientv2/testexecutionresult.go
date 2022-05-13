package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Testexecutionresult
type Testexecutionresult struct { 
	// Operations - Execution operations performed as part of the test
	Operations *[]Testexecutionoperationresult `json:"operations,omitempty"`


	// VarError - The final error encountered during the test that resulted in test failure
	VarError *Errorbody `json:"error,omitempty"`


	// FinalResult - The final result of the test. This is the response that would be returned during normal action execution
	FinalResult *interface{} `json:"finalResult,omitempty"`


	// Success - Indicates whether or not the test was a success
	Success *bool `json:"success,omitempty"`

}

func (o *Testexecutionresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Testexecutionresult
	
	return json.Marshal(&struct { 
		Operations *[]Testexecutionoperationresult `json:"operations,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		
		FinalResult *interface{} `json:"finalResult,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		*Alias
	}{ 
		Operations: o.Operations,
		
		VarError: o.VarError,
		
		FinalResult: o.FinalResult,
		
		Success: o.Success,
		Alias:    (*Alias)(o),
	})
}

func (o *Testexecutionresult) UnmarshalJSON(b []byte) error {
	var TestexecutionresultMap map[string]interface{}
	err := json.Unmarshal(b, &TestexecutionresultMap)
	if err != nil {
		return err
	}
	
	if Operations, ok := TestexecutionresultMap["operations"].([]interface{}); ok {
		OperationsString, _ := json.Marshal(Operations)
		json.Unmarshal(OperationsString, &o.Operations)
	}
	
	if VarError, ok := TestexecutionresultMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if FinalResult, ok := TestexecutionresultMap["finalResult"].(map[string]interface{}); ok {
		FinalResultString, _ := json.Marshal(FinalResult)
		json.Unmarshal(FinalResultString, &o.FinalResult)
	}
	
	if Success, ok := TestexecutionresultMap["success"].(bool); ok {
		o.Success = &Success
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Testexecutionresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
