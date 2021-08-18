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

func (u *Asyncforecastoperationresult) MarshalJSON() ([]byte, error) {
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
		Status: u.Status,
		
		OperationId: u.OperationId,
		
		Result: u.Result,
		
		Progress: u.Progress,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Asyncforecastoperationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
