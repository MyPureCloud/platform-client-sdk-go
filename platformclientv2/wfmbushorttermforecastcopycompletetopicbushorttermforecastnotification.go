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

func (u *Wfmbushorttermforecastcopycompletetopicbushorttermforecastnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastcopycompletetopicbushorttermforecastnotification

	

	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		Result *Wfmbushorttermforecastcopycompletetopicbushorttermforecast `json:"result,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		*Alias
	}{ 
		Status: u.Status,
		
		Result: u.Result,
		
		OperationId: u.OperationId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastcopycompletetopicbushorttermforecastnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
