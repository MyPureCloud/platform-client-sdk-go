package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Servicelevel
type Servicelevel struct { 
	// Percentage - The desired Service Level. A value between 0 and 1.
	Percentage *float64 `json:"percentage,omitempty"`


	// DurationMs - Service Level target in milliseconds.
	DurationMs *int `json:"durationMs,omitempty"`

}

// String returns a JSON representation of the model
func (o *Servicelevel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
