package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Facetstatistics
type Facetstatistics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Facetstatistics) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Facetstatistics) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateMin","DateMax", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		Alias
	}{ 
		Count: o.Count,
		
		Min: o.Min,
		
		Max: o.Max,
		
		Mean: o.Mean,
		
		StdDeviation: o.StdDeviation,
		
		DateMin: DateMin,
		
		DateMax: DateMax,
		Alias:    (Alias)(o),
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
