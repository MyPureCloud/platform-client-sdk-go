package platformclientv2
import (
	"encoding/json"
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
	return string(j)
}
