package platformclientv2
import (
	"encoding/json"
)

// Generateshorttermforecastresponse
type Generateshorttermforecastresponse struct { 
	// Status - The status of the request
	Status *string `json:"status,omitempty"`


	// Result - The resulting forecast.  May be sent asynchronously via notification depending on the complexity of the forecast
	Result *Shorttermforecast `json:"result,omitempty"`


	// OperationId - The operation id to watch for on the notification topic
	OperationId *string `json:"operationId,omitempty"`


	// Progress - Percent progress.  Subscribe to the corresponding notification to view progress and await the result
	Progress *int32 `json:"progress,omitempty"`

}

// String returns a JSON representation of the model
func (o *Generateshorttermforecastresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
