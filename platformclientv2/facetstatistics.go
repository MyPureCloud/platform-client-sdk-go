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

func (o *Facetstatistics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facetstatistics
	
	DateMin := new(string)
	if o.DateMin != nil {
		
		*DateMin = timeutil.Strftime(o.DateMin, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateMin = nil
	}
	
	DateMax := new(string)
	if o.DateMax != nil {
		
		*DateMax = timeutil.Strftime(o.DateMax, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Count: o.Count,
		
		Min: o.Min,
		
		Max: o.Max,
		
		Mean: o.Mean,
		
		StdDeviation: o.StdDeviation,
		
		DateMin: DateMin,
		
		DateMax: DateMax,
		Alias:    (*Alias)(o),
	})
}

func (o *Facetstatistics) UnmarshalJSON(b []byte) error {
	var FacetstatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &FacetstatisticsMap)
	if err != nil {
		return err
	}
	
	if Count, ok := FacetstatisticsMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Min, ok := FacetstatisticsMap["min"].(float64); ok {
		o.Min = &Min
	}
    
	if Max, ok := FacetstatisticsMap["max"].(float64); ok {
		o.Max = &Max
	}
    
	if Mean, ok := FacetstatisticsMap["mean"].(float64); ok {
		o.Mean = &Mean
	}
    
	if StdDeviation, ok := FacetstatisticsMap["stdDeviation"].(float64); ok {
		o.StdDeviation = &StdDeviation
	}
    
	if dateMinString, ok := FacetstatisticsMap["dateMin"].(string); ok {
		DateMin, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateMinString)
		o.DateMin = &DateMin
	}
	
	if dateMaxString, ok := FacetstatisticsMap["dateMax"].(string); ok {
		DateMax, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateMaxString)
		o.DateMax = &DateMax
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facetstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
