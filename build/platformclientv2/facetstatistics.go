package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Facetstatistics
type Facetstatistics struct { 
	// Count
	Count *int `json:"count,omitempty"`


	// Min
	Min *float64 `json:"min,omitempty"`


	// Max
	Max *float64 `json:"max,omitempty"`


	// Mean
	Mean *float64 `json:"mean,omitempty"`


	// StdDeviation
	StdDeviation *float64 `json:"stdDeviation,omitempty"`


	// DateMin - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateMin *time.Time `json:"dateMin,omitempty"`


	// DateMax - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateMax *time.Time `json:"dateMax,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facetstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
