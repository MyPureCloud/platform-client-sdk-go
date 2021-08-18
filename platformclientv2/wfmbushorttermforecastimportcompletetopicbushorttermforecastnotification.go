package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification
type Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// Result
	Result *Wfmbushorttermforecastimportcompletetopicbushorttermforecast `json:"result,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`

}

func (u *Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification

	

	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		Result *Wfmbushorttermforecastimportcompletetopicbushorttermforecast `json:"result,omitempty"`
		
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
func (o *Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
