package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastmetadata
type Forecastmetadata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateCreated - Forecast creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// ForecastId - Forecast ID
	ForecastId *string `json:"forecastId,omitempty"`

	// IntervalLengthInMinutes - Interval length
	IntervalLengthInMinutes *string `json:"intervalLengthInMinutes,omitempty"`

	// Source - Forecast source
	Source *string `json:"source,omitempty"`

	// DateStart - Forecast start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// TimeZone - Timezone of the business unit
	TimeZone *string `json:"timeZone,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Forecastmetadata) SetField(field string, fieldValue interface{}) {
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

func (o Forecastmetadata) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateStart", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
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
	type Alias Forecastmetadata
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ForecastId *string `json:"forecastId,omitempty"`
		
		IntervalLengthInMinutes *string `json:"intervalLengthInMinutes,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		Alias
	}{ 
		DateCreated: DateCreated,
		
		ForecastId: o.ForecastId,
		
		IntervalLengthInMinutes: o.IntervalLengthInMinutes,
		
		Source: o.Source,
		
		DateStart: DateStart,
		
		TimeZone: o.TimeZone,
		Alias:    (Alias)(o),
	})
}

func (o *Forecastmetadata) UnmarshalJSON(b []byte) error {
	var ForecastmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ForecastmetadataMap)
	if err != nil {
		return err
	}
	
	if dateCreatedString, ok := ForecastmetadataMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ForecastId, ok := ForecastmetadataMap["forecastId"].(string); ok {
		o.ForecastId = &ForecastId
	}
    
	if IntervalLengthInMinutes, ok := ForecastmetadataMap["intervalLengthInMinutes"].(string); ok {
		o.IntervalLengthInMinutes = &IntervalLengthInMinutes
	}
    
	if Source, ok := ForecastmetadataMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if dateStartString, ok := ForecastmetadataMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if TimeZone, ok := ForecastmetadataMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Forecastmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
