package platformclientv2
import (
	"encoding/json"
)

// Wfmbuintradaydataupdatetopicbuintradayforecastdata
type Wfmbuintradaydataupdatetopicbuintradayforecastdata struct { 
	// Offered
	Offered *float32 `json:"offered,omitempty"`


	// AverageHandleTimeSeconds
	AverageHandleTimeSeconds *float32 `json:"averageHandleTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayforecastdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
