package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastupdatecompletetopicbushorttermforecastnotification
type Wfmbushorttermforecastupdatecompletetopicbushorttermforecastnotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// Result
	Result *Wfmbushorttermforecastupdatecompletetopicbushorttermforecast `json:"result,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastupdatecompletetopicbushorttermforecastnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
