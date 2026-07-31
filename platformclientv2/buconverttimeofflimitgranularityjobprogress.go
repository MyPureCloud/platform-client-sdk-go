package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buconverttimeofflimitgranularityjobprogress
type Buconverttimeofflimitgranularityjobprogress struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateEarliestComplete - Earliest date completed for time-off limit granularity conversion. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEarliestComplete *time.Time `json:"dateEarliestComplete,omitempty"`

	// DateLatestComplete - Latest date completed for time-off limit granularity conversion. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateLatestComplete *time.Time `json:"dateLatestComplete,omitempty"`

	// NumberOfDaysComplete - Number of days completed for time-off limit granularity conversion
	NumberOfDaysComplete *int `json:"numberOfDaysComplete,omitempty"`

	// PercentageComplete - Percentage completed for time-off limit granularity conversion
	PercentageComplete *int `json:"percentageComplete,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buconverttimeofflimitgranularityjobprogress) SetField(field string, fieldValue interface{}) {
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

func (o Buconverttimeofflimitgranularityjobprogress) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateEarliestComplete","DateLatestComplete", }

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
	type Alias Buconverttimeofflimitgranularityjobprogress
	
	DateEarliestComplete := new(string)
	if o.DateEarliestComplete != nil {
		*DateEarliestComplete = timeutil.Strftime(o.DateEarliestComplete, "%Y-%m-%d")
	} else {
		DateEarliestComplete = nil
	}
	
	DateLatestComplete := new(string)
	if o.DateLatestComplete != nil {
		*DateLatestComplete = timeutil.Strftime(o.DateLatestComplete, "%Y-%m-%d")
	} else {
		DateLatestComplete = nil
	}
	
	return json.Marshal(&struct { 
		DateEarliestComplete *string `json:"dateEarliestComplete,omitempty"`
		
		DateLatestComplete *string `json:"dateLatestComplete,omitempty"`
		
		NumberOfDaysComplete *int `json:"numberOfDaysComplete,omitempty"`
		
		PercentageComplete *int `json:"percentageComplete,omitempty"`
		Alias
	}{ 
		DateEarliestComplete: DateEarliestComplete,
		
		DateLatestComplete: DateLatestComplete,
		
		NumberOfDaysComplete: o.NumberOfDaysComplete,
		
		PercentageComplete: o.PercentageComplete,
		Alias:    (Alias)(o),
	})
}

func (o *Buconverttimeofflimitgranularityjobprogress) UnmarshalJSON(b []byte) error {
	var BuconverttimeofflimitgranularityjobprogressMap map[string]interface{}
	err := json.Unmarshal(b, &BuconverttimeofflimitgranularityjobprogressMap)
	if err != nil {
		return err
	}
	
	if dateEarliestCompleteString, ok := BuconverttimeofflimitgranularityjobprogressMap["dateEarliestComplete"].(string); ok {
		DateEarliestComplete, _ := time.Parse("2006-01-02", dateEarliestCompleteString)
		o.DateEarliestComplete = &DateEarliestComplete
	}
	
	if dateLatestCompleteString, ok := BuconverttimeofflimitgranularityjobprogressMap["dateLatestComplete"].(string); ok {
		DateLatestComplete, _ := time.Parse("2006-01-02", dateLatestCompleteString)
		o.DateLatestComplete = &DateLatestComplete
	}
	
	if NumberOfDaysComplete, ok := BuconverttimeofflimitgranularityjobprogressMap["numberOfDaysComplete"].(float64); ok {
		NumberOfDaysCompleteInt := int(NumberOfDaysComplete)
		o.NumberOfDaysComplete = &NumberOfDaysCompleteInt
	}
	
	if PercentageComplete, ok := BuconverttimeofflimitgranularityjobprogressMap["percentageComplete"].(float64); ok {
		PercentageCompleteInt := int(PercentageComplete)
		o.PercentageComplete = &PercentageCompleteInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buconverttimeofflimitgranularityjobprogress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
