package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Testexecutionresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
