package platformclientv2
import (
	"encoding/json"
)

// Wfmintradaydataupdatetopicintradayforecastdata
type Wfmintradaydataupdatetopicintradayforecastdata struct { 
	// Offered
	Offered *float32 `json:"offered,omitempty"`


	// AverageTalkTimeSeconds
	AverageTalkTimeSeconds *float32 `json:"averageTalkTimeSeconds,omitempty"`


	// AverageAfterCallWorkSeconds
	AverageAfterCallWorkSeconds *float32 `json:"averageAfterCallWorkSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayforecastdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
