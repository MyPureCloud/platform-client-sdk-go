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

func (u *Wfmbuintradaydataupdatetopicbuintradaynotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbuintradaynotification

	

	return json.Marshal(&struct { 
		OperationId *string `json:"operationId,omitempty"`
		
		Result *Wfmbuintradaydataupdatetopicbuintradayresult `json:"result,omitempty"`
		*Alias
	}{ 
		OperationId: u.OperationId,
		
		Result: u.Result,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradaynotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
