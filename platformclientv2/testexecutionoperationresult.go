package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Testexecutionoperationresult
type Testexecutionoperationresult struct { 
	// Step - The step number to indicate the order in which the operation was performed
	Step *int `json:"step,omitempty"`


	// Name - Name of the operation performed
	Name *string `json:"name,omitempty"`


	// Success - Indicated whether or not the operation was successful
	Success *bool `json:"success,omitempty"`


	// Result - The result of the operation
	Result *interface{} `json:"result,omitempty"`


	// VarError - Error that occurred during the operation
	VarError *Errorbody `json:"error,omitempty"`

}

func (o *Testexecutionoperationresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Testexecutionoperationresult
	
	return json.Marshal(&struct { 
		Step *int `json:"step,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		
		Result *interface{} `json:"result,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		*Alias
	}{ 
		Step: o.Step,
		
		Name: o.Name,
		
		Success: o.Success,
		
		Result: o.Result,
		
		VarError: o.VarError,
		Alias:    (*Alias)(o),
	})
}

func (o *Testexecutionoperationresult) UnmarshalJSON(b []byte) error {
	var TestexecutionoperationresultMap map[string]interface{}
	err := json.Unmarshal(b, &TestexecutionoperationresultMap)
	if err != nil {
		return err
	}
	
	if Step, ok := TestexecutionoperationresultMap["step"].(float64); ok {
		StepInt := int(Step)
		o.Step = &StepInt
	}
	
	if Name, ok := TestexecutionoperationresultMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Success, ok := TestexecutionoperationresultMap["success"].(bool); ok {
		o.Success = &Success
	}
    
	if Result, ok := TestexecutionoperationresultMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if VarError, ok := TestexecutionoperationresultMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Testexecutionoperationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
