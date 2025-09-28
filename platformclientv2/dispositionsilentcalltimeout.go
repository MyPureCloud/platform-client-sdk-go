package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dispositionsilentcalltimeout
type Dispositionsilentcalltimeout struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TimeoutMs - Configured silent call timeout value.
	TimeoutMs *int `json:"timeoutMs,omitempty"`

	// TimerStartTime - Timer start time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	TimerStartTime *time.Time `json:"timerStartTime,omitempty"`

	// TimerEndTime - Timer end time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	TimerEndTime *time.Time `json:"timerEndTime,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dispositionsilentcalltimeout) SetField(field string, fieldValue interface{}) {
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

func (o Dispositionsilentcalltimeout) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "TimerStartTime","TimerEndTime", }
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
	type Alias Dispositionsilentcalltimeout
	
	TimerStartTime := new(string)
	if o.TimerStartTime != nil {
		
		*TimerStartTime = timeutil.Strftime(o.TimerStartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		TimerStartTime = nil
	}
	
	TimerEndTime := new(string)
	if o.TimerEndTime != nil {
		
		*TimerEndTime = timeutil.Strftime(o.TimerEndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		TimerEndTime = nil
	}
	
	return json.Marshal(&struct { 
		TimeoutMs *int `json:"timeoutMs,omitempty"`
		
		TimerStartTime *string `json:"timerStartTime,omitempty"`
		
		TimerEndTime *string `json:"timerEndTime,omitempty"`
		Alias
	}{ 
		TimeoutMs: o.TimeoutMs,
		
		TimerStartTime: TimerStartTime,
		
		TimerEndTime: TimerEndTime,
		Alias:    (Alias)(o),
	})
}

func (o *Dispositionsilentcalltimeout) UnmarshalJSON(b []byte) error {
	var DispositionsilentcalltimeoutMap map[string]interface{}
	err := json.Unmarshal(b, &DispositionsilentcalltimeoutMap)
	if err != nil {
		return err
	}
	
	if TimeoutMs, ok := DispositionsilentcalltimeoutMap["timeoutMs"].(float64); ok {
		TimeoutMsInt := int(TimeoutMs)
		o.TimeoutMs = &TimeoutMsInt
	}
	
	if timerStartTimeString, ok := DispositionsilentcalltimeoutMap["timerStartTime"].(string); ok {
		TimerStartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", timerStartTimeString)
		o.TimerStartTime = &TimerStartTime
	}
	
	if timerEndTimeString, ok := DispositionsilentcalltimeoutMap["timerEndTime"].(string); ok {
		TimerEndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", timerEndTimeString)
		o.TimerEndTime = &TimerEndTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dispositionsilentcalltimeout) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
