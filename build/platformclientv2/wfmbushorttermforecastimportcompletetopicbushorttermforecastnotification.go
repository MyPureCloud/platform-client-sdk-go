package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
