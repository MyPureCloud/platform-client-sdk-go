package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Waitlistposition
type Waitlistposition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TimeOffRequest - The time off request for this wait list position
	TimeOffRequest *Timeoffrequestreference `json:"timeOffRequest,omitempty"`

	// TimeOffLimit - The time off limit for which time off request is waitlisted
	TimeOffLimit *Timeofflimitreference `json:"timeOffLimit,omitempty"`

	// Date - The date to which this wait list position applies, as defined by the time zone of the business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	Date *time.Time `json:"date,omitempty"`

	// WaitlistPosition - The time off request's position in the waitlist on the date. 1 means time off is the first in the waitlist
	WaitlistPosition *int `json:"waitlistPosition,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Waitlistposition) SetField(field string, fieldValue interface{}) {
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

func (o Waitlistposition) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "Date", }

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
	type Alias Waitlistposition
	
	Date := new(string)
	if o.Date != nil {
		*Date = timeutil.Strftime(o.Date, "%Y-%m-%d")
	} else {
		Date = nil
	}
	
	return json.Marshal(&struct { 
		TimeOffRequest *Timeoffrequestreference `json:"timeOffRequest,omitempty"`
		
		TimeOffLimit *Timeofflimitreference `json:"timeOffLimit,omitempty"`
		
		Date *string `json:"date,omitempty"`
		
		WaitlistPosition *int `json:"waitlistPosition,omitempty"`
		Alias
	}{ 
		TimeOffRequest: o.TimeOffRequest,
		
		TimeOffLimit: o.TimeOffLimit,
		
		Date: Date,
		
		WaitlistPosition: o.WaitlistPosition,
		Alias:    (Alias)(o),
	})
}

func (o *Waitlistposition) UnmarshalJSON(b []byte) error {
	var WaitlistpositionMap map[string]interface{}
	err := json.Unmarshal(b, &WaitlistpositionMap)
	if err != nil {
		return err
	}
	
	if TimeOffRequest, ok := WaitlistpositionMap["timeOffRequest"].(map[string]interface{}); ok {
		TimeOffRequestString, _ := json.Marshal(TimeOffRequest)
		json.Unmarshal(TimeOffRequestString, &o.TimeOffRequest)
	}
	
	if TimeOffLimit, ok := WaitlistpositionMap["timeOffLimit"].(map[string]interface{}); ok {
		TimeOffLimitString, _ := json.Marshal(TimeOffLimit)
		json.Unmarshal(TimeOffLimitString, &o.TimeOffLimit)
	}
	
	if dateString, ok := WaitlistpositionMap["date"].(string); ok {
		Date, _ := time.Parse("2006-01-02", dateString)
		o.Date = &Date
	}
	
	if WaitlistPosition, ok := WaitlistpositionMap["waitlistPosition"].(float64); ok {
		WaitlistPositionInt := int(WaitlistPosition)
		o.WaitlistPosition = &WaitlistPositionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Waitlistposition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
