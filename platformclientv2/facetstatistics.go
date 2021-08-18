package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Facetstatistics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facetstatistics

	
	DateMin := new(string)
	if u.DateMin != nil {
		
		*DateMin = timeutil.Strftime(u.DateMin, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateMin = nil
	}
	
	DateMax := new(string)
	if u.DateMax != nil {
		
		*DateMax = timeutil.Strftime(u.DateMax, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateMax = nil
	}
	

	return json.Marshal(&struct { 
		Count *int `json:"count,omitempty"`
		
		Min *float64 `json:"min,omitempty"`
		
		Max *float64 `json:"max,omitempty"`
		
		Mean *float64 `json:"mean,omitempty"`
		
		StdDeviation *float64 `json:"stdDeviation,omitempty"`
		
		DateMin *string `json:"dateMin,omitempty"`
		
		DateMax *string `json:"dateMax,omitempty"`
		*Alias
	}{ 
		Count: u.Count,
		
		Min: u.Min,
		
		Max: u.Max,
		
		Mean: u.Mean,
		
		StdDeviation: u.StdDeviation,
		
		DateMin: DateMin,
		
		DateMax: DateMax,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Facetstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
