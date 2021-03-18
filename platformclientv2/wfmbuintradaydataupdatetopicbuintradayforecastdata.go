package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
