package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeofflimitvalues
type Timeofflimitvalues struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// LimitMinutes - Time-off limit values in minutes per granularity interval
	LimitMinutes *[]int `json:"limitMinutes,omitempty"`

	// AllocatedMinutes - Allocated time-off minutes per granularity interval
	AllocatedMinutes *[]int `json:"allocatedMinutes,omitempty"`

	// WaitlistedMinutes - Waitlisted time-off minutes per granularity interval
	WaitlistedMinutes *[]int `json:"waitlistedMinutes,omitempty"`

	// WaitlistedRequests - The current number of waitlisted time-off requests per granularity interval
	WaitlistedRequests *[]int `json:"waitlistedRequests,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timeofflimitvalues) SetField(field string, fieldValue interface{}) {
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

func (o Timeofflimitvalues) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Timeofflimitvalues
	
	return json.Marshal(&struct { 
		LimitMinutes *[]int `json:"limitMinutes,omitempty"`
		
		AllocatedMinutes *[]int `json:"allocatedMinutes,omitempty"`
		
		WaitlistedMinutes *[]int `json:"waitlistedMinutes,omitempty"`
		
		WaitlistedRequests *[]int `json:"waitlistedRequests,omitempty"`
		Alias
	}{ 
		LimitMinutes: o.LimitMinutes,
		
		AllocatedMinutes: o.AllocatedMinutes,
		
		WaitlistedMinutes: o.WaitlistedMinutes,
		
		WaitlistedRequests: o.WaitlistedRequests,
		Alias:    (Alias)(o),
	})
}

func (o *Timeofflimitvalues) UnmarshalJSON(b []byte) error {
	var TimeofflimitvaluesMap map[string]interface{}
	err := json.Unmarshal(b, &TimeofflimitvaluesMap)
	if err != nil {
		return err
	}
	
	if LimitMinutes, ok := TimeofflimitvaluesMap["limitMinutes"].([]interface{}); ok {
		LimitMinutesString, _ := json.Marshal(LimitMinutes)
		json.Unmarshal(LimitMinutesString, &o.LimitMinutes)
	}
	
	if AllocatedMinutes, ok := TimeofflimitvaluesMap["allocatedMinutes"].([]interface{}); ok {
		AllocatedMinutesString, _ := json.Marshal(AllocatedMinutes)
		json.Unmarshal(AllocatedMinutesString, &o.AllocatedMinutes)
	}
	
	if WaitlistedMinutes, ok := TimeofflimitvaluesMap["waitlistedMinutes"].([]interface{}); ok {
		WaitlistedMinutesString, _ := json.Marshal(WaitlistedMinutes)
		json.Unmarshal(WaitlistedMinutesString, &o.WaitlistedMinutes)
	}
	
	if WaitlistedRequests, ok := TimeofflimitvaluesMap["waitlistedRequests"].([]interface{}); ok {
		WaitlistedRequestsString, _ := json.Marshal(WaitlistedRequests)
		json.Unmarshal(WaitlistedRequestsString, &o.WaitlistedRequests)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeofflimitvalues) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
