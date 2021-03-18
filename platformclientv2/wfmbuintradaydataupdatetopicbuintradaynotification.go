package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradaynotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
