package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Validateagentunavailabletimesrequest
type Validateagentunavailabletimesrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ValidationWeekDate - The ID of the week to validate. Must correspond to the start day of week of the business unit to which the agent belongs in the format YYYY-MM-dd. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	ValidationWeekDate *time.Time `json:"validationWeekDate,omitempty"`

	// UnavailableTimes - Proposed changes to agent's unavailable time spans to be validated
	UnavailableTimes *[]Updateunavailabletime `json:"unavailableTimes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Validateagentunavailabletimesrequest) SetField(field string, fieldValue interface{}) {
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

func (o Validateagentunavailabletimesrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "ValidationWeekDate", }

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
	type Alias Validateagentunavailabletimesrequest
	
	ValidationWeekDate := new(string)
	if o.ValidationWeekDate != nil {
		*ValidationWeekDate = timeutil.Strftime(o.ValidationWeekDate, "%Y-%m-%d")
	} else {
		ValidationWeekDate = nil
	}
	
	return json.Marshal(&struct { 
		ValidationWeekDate *string `json:"validationWeekDate,omitempty"`
		
		UnavailableTimes *[]Updateunavailabletime `json:"unavailableTimes,omitempty"`
		Alias
	}{ 
		ValidationWeekDate: ValidationWeekDate,
		
		UnavailableTimes: o.UnavailableTimes,
		Alias:    (Alias)(o),
	})
}

func (o *Validateagentunavailabletimesrequest) UnmarshalJSON(b []byte) error {
	var ValidateagentunavailabletimesrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ValidateagentunavailabletimesrequestMap)
	if err != nil {
		return err
	}
	
	if validationWeekDateString, ok := ValidateagentunavailabletimesrequestMap["validationWeekDate"].(string); ok {
		ValidationWeekDate, _ := time.Parse("2006-01-02", validationWeekDateString)
		o.ValidationWeekDate = &ValidationWeekDate
	}
	
	if UnavailableTimes, ok := ValidateagentunavailabletimesrequestMap["unavailableTimes"].([]interface{}); ok {
		UnavailableTimesString, _ := json.Marshal(UnavailableTimes)
		json.Unmarshal(UnavailableTimesString, &o.UnavailableTimes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Validateagentunavailabletimesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
