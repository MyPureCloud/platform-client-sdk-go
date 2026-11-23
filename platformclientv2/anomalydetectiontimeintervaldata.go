package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Anomalydetectiontimeintervaldata
type Anomalydetectiontimeintervaldata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateIntervalStart - Start time of the interval, in ISO-8601 format
	DateIntervalStart *time.Time `json:"dateIntervalStart,omitempty"`

	// DateIntervalEnd - End time of the interval, in ISO-8601 format
	DateIntervalEnd *time.Time `json:"dateIntervalEnd,omitempty"`

	// ExpectedValue - Expected value for this time interval
	ExpectedValue *float64 `json:"expectedValue,omitempty"`

	// LowerBound - Lower bound of the forecasted range
	LowerBound *float64 `json:"lowerBound,omitempty"`

	// UpperBound - Upper bound of the forecasted range
	UpperBound *float64 `json:"upperBound,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Anomalydetectiontimeintervaldata) SetField(field string, fieldValue interface{}) {
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

func (o Anomalydetectiontimeintervaldata) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateIntervalStart","DateIntervalEnd", }
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
	type Alias Anomalydetectiontimeintervaldata
	
	DateIntervalStart := new(string)
	if o.DateIntervalStart != nil {
		
		*DateIntervalStart = timeutil.Strftime(o.DateIntervalStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateIntervalStart = nil
	}
	
	DateIntervalEnd := new(string)
	if o.DateIntervalEnd != nil {
		
		*DateIntervalEnd = timeutil.Strftime(o.DateIntervalEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateIntervalEnd = nil
	}
	
	return json.Marshal(&struct { 
		DateIntervalStart *string `json:"dateIntervalStart,omitempty"`
		
		DateIntervalEnd *string `json:"dateIntervalEnd,omitempty"`
		
		ExpectedValue *float64 `json:"expectedValue,omitempty"`
		
		LowerBound *float64 `json:"lowerBound,omitempty"`
		
		UpperBound *float64 `json:"upperBound,omitempty"`
		Alias
	}{ 
		DateIntervalStart: DateIntervalStart,
		
		DateIntervalEnd: DateIntervalEnd,
		
		ExpectedValue: o.ExpectedValue,
		
		LowerBound: o.LowerBound,
		
		UpperBound: o.UpperBound,
		Alias:    (Alias)(o),
	})
}

func (o *Anomalydetectiontimeintervaldata) UnmarshalJSON(b []byte) error {
	var AnomalydetectiontimeintervaldataMap map[string]interface{}
	err := json.Unmarshal(b, &AnomalydetectiontimeintervaldataMap)
	if err != nil {
		return err
	}
	
	if dateIntervalStartString, ok := AnomalydetectiontimeintervaldataMap["dateIntervalStart"].(string); ok {
		DateIntervalStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateIntervalStartString)
		o.DateIntervalStart = &DateIntervalStart
	}
	
	if dateIntervalEndString, ok := AnomalydetectiontimeintervaldataMap["dateIntervalEnd"].(string); ok {
		DateIntervalEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateIntervalEndString)
		o.DateIntervalEnd = &DateIntervalEnd
	}
	
	if ExpectedValue, ok := AnomalydetectiontimeintervaldataMap["expectedValue"].(float64); ok {
		o.ExpectedValue = &ExpectedValue
	}
    
	if LowerBound, ok := AnomalydetectiontimeintervaldataMap["lowerBound"].(float64); ok {
		o.LowerBound = &LowerBound
	}
    
	if UpperBound, ok := AnomalydetectiontimeintervaldataMap["upperBound"].(float64); ok {
		o.UpperBound = &UpperBound
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Anomalydetectiontimeintervaldata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
