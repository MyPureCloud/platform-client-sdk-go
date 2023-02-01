package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletimeoffrange
type Availabletimeoffrange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TimeOffLimit - The time off limit
	TimeOffLimit *Timeofflimitreference `json:"timeOffLimit,omitempty"`

	// StartDate - Start date of the requested date range. The end date is determined by the size of interval list. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartDate *time.Time `json:"startDate,omitempty"`

	// Granularity - Granularity choice for time off limit
	Granularity *string `json:"granularity,omitempty"`

	// AvailableMinutesPerInterval - The list of available time off values in minutes per granularity interval
	AvailableMinutesPerInterval *[]int `json:"availableMinutesPerInterval,omitempty"`

	// WaitlistedRequestsPerInterval - The current number of waitlisted time off requests for every interval per granularity
	WaitlistedRequestsPerInterval *[]int `json:"waitlistedRequestsPerInterval,omitempty"`

	// WaitlistEnabled - Whether the time off request can be waitlisted
	WaitlistEnabled *bool `json:"waitlistEnabled,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Availabletimeoffrange) SetField(field string, fieldValue interface{}) {
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

func (o Availabletimeoffrange) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "StartDate", }

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
	type Alias Availabletimeoffrange
	
	StartDate := new(string)
	if o.StartDate != nil {
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%d")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		TimeOffLimit *Timeofflimitreference `json:"timeOffLimit,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		AvailableMinutesPerInterval *[]int `json:"availableMinutesPerInterval,omitempty"`
		
		WaitlistedRequestsPerInterval *[]int `json:"waitlistedRequestsPerInterval,omitempty"`
		
		WaitlistEnabled *bool `json:"waitlistEnabled,omitempty"`
		Alias
	}{ 
		TimeOffLimit: o.TimeOffLimit,
		
		StartDate: StartDate,
		
		Granularity: o.Granularity,
		
		AvailableMinutesPerInterval: o.AvailableMinutesPerInterval,
		
		WaitlistedRequestsPerInterval: o.WaitlistedRequestsPerInterval,
		
		WaitlistEnabled: o.WaitlistEnabled,
		Alias:    (Alias)(o),
	})
}

func (o *Availabletimeoffrange) UnmarshalJSON(b []byte) error {
	var AvailabletimeoffrangeMap map[string]interface{}
	err := json.Unmarshal(b, &AvailabletimeoffrangeMap)
	if err != nil {
		return err
	}
	
	if TimeOffLimit, ok := AvailabletimeoffrangeMap["timeOffLimit"].(map[string]interface{}); ok {
		TimeOffLimitString, _ := json.Marshal(TimeOffLimit)
		json.Unmarshal(TimeOffLimitString, &o.TimeOffLimit)
	}
	
	if startDateString, ok := AvailabletimeoffrangeMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02", startDateString)
		o.StartDate = &StartDate
	}
	
	if Granularity, ok := AvailabletimeoffrangeMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if AvailableMinutesPerInterval, ok := AvailabletimeoffrangeMap["availableMinutesPerInterval"].([]interface{}); ok {
		AvailableMinutesPerIntervalString, _ := json.Marshal(AvailableMinutesPerInterval)
		json.Unmarshal(AvailableMinutesPerIntervalString, &o.AvailableMinutesPerInterval)
	}
	
	if WaitlistedRequestsPerInterval, ok := AvailabletimeoffrangeMap["waitlistedRequestsPerInterval"].([]interface{}); ok {
		WaitlistedRequestsPerIntervalString, _ := json.Marshal(WaitlistedRequestsPerInterval)
		json.Unmarshal(WaitlistedRequestsPerIntervalString, &o.WaitlistedRequestsPerInterval)
	}
	
	if WaitlistEnabled, ok := AvailabletimeoffrangeMap["waitlistEnabled"].(bool); ok {
		o.WaitlistEnabled = &WaitlistEnabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Availabletimeoffrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
