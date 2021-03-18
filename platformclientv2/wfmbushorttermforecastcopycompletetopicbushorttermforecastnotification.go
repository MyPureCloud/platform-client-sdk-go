package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastcopycompletetopicbushorttermforecastnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
