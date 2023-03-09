package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentpossibleworkshiftsresponse
type Agentpossibleworkshiftsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// WeekStartDate - Start date of requested effective work plan. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekStartDate *time.Time `json:"weekStartDate,omitempty"`

	// Pattern - Each element is the ID of an effective work plan for a specific week
	Pattern *[]int `json:"pattern,omitempty"`

	// WeeklyPossibleWorkShifts - Each element is a weekly effective work plan that can be used for multiple weeks
	WeeklyPossibleWorkShifts *[]Possibleworkshiftsforweek `json:"weeklyPossibleWorkShifts,omitempty"`

	// SchedulerIntervalLengthMinutes - Number of minutes in each interval in the intervalScheduleProbabilities
	SchedulerIntervalLengthMinutes *int `json:"schedulerIntervalLengthMinutes,omitempty"`

	// TimeZone - The time zone of the business unit
	TimeZone *string `json:"timeZone,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentpossibleworkshiftsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Agentpossibleworkshiftsresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "WeekStartDate", }

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
	type Alias Agentpossibleworkshiftsresponse
	
	WeekStartDate := new(string)
	if o.WeekStartDate != nil {
		*WeekStartDate = timeutil.Strftime(o.WeekStartDate, "%Y-%m-%d")
	} else {
		WeekStartDate = nil
	}
	
	return json.Marshal(&struct { 
		WeekStartDate *string `json:"weekStartDate,omitempty"`
		
		Pattern *[]int `json:"pattern,omitempty"`
		
		WeeklyPossibleWorkShifts *[]Possibleworkshiftsforweek `json:"weeklyPossibleWorkShifts,omitempty"`
		
		SchedulerIntervalLengthMinutes *int `json:"schedulerIntervalLengthMinutes,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		Alias
	}{ 
		WeekStartDate: WeekStartDate,
		
		Pattern: o.Pattern,
		
		WeeklyPossibleWorkShifts: o.WeeklyPossibleWorkShifts,
		
		SchedulerIntervalLengthMinutes: o.SchedulerIntervalLengthMinutes,
		
		TimeZone: o.TimeZone,
		Alias:    (Alias)(o),
	})
}

func (o *Agentpossibleworkshiftsresponse) UnmarshalJSON(b []byte) error {
	var AgentpossibleworkshiftsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AgentpossibleworkshiftsresponseMap)
	if err != nil {
		return err
	}
	
	if weekStartDateString, ok := AgentpossibleworkshiftsresponseMap["weekStartDate"].(string); ok {
		WeekStartDate, _ := time.Parse("2006-01-02", weekStartDateString)
		o.WeekStartDate = &WeekStartDate
	}
	
	if Pattern, ok := AgentpossibleworkshiftsresponseMap["pattern"].([]interface{}); ok {
		PatternString, _ := json.Marshal(Pattern)
		json.Unmarshal(PatternString, &o.Pattern)
	}
	
	if WeeklyPossibleWorkShifts, ok := AgentpossibleworkshiftsresponseMap["weeklyPossibleWorkShifts"].([]interface{}); ok {
		WeeklyPossibleWorkShiftsString, _ := json.Marshal(WeeklyPossibleWorkShifts)
		json.Unmarshal(WeeklyPossibleWorkShiftsString, &o.WeeklyPossibleWorkShifts)
	}
	
	if SchedulerIntervalLengthMinutes, ok := AgentpossibleworkshiftsresponseMap["schedulerIntervalLengthMinutes"].(float64); ok {
		SchedulerIntervalLengthMinutesInt := int(SchedulerIntervalLengthMinutes)
		o.SchedulerIntervalLengthMinutes = &SchedulerIntervalLengthMinutesInt
	}
	
	if TimeZone, ok := AgentpossibleworkshiftsresponseMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentpossibleworkshiftsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
