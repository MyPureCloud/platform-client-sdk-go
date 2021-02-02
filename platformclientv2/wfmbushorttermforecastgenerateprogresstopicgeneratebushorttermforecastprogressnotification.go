package platformclientv2
import (
	"encoding/json"
)

// Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification
type Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// Result
	Result *Wfmbushorttermforecastgenerateprogresstopicbushorttermforecast `json:"result,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// Progress
	Progress *int `json:"progress,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastgenerateprogresstopicgeneratebushorttermforecastprogressnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
