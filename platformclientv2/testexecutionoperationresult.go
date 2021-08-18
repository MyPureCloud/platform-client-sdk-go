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

func (u *Testexecutionoperationresult) MarshalJSON() ([]byte, error) {
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
		Step: u.Step,
		
		Name: u.Name,
		
		Success: u.Success,
		
		Result: u.Result,
		
		VarError: u.VarError,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Testexecutionoperationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
