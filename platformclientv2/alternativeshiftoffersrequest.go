package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Alternativeshiftoffersrequest
type Alternativeshiftoffersrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Schedule - The existing schedule being used to find alternative shift offers
	Schedule *Alternativeshiftschedulelookup `json:"schedule,omitempty"`

	// QueryWeekDate - The start date for the week in this schedule in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	QueryWeekDate *time.Time `json:"queryWeekDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Alternativeshiftoffersrequest) SetField(field string, fieldValue interface{}) {
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

func (o Alternativeshiftoffersrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "QueryWeekDate", }

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
	type Alias Alternativeshiftoffersrequest
	
	QueryWeekDate := new(string)
	if o.QueryWeekDate != nil {
		*QueryWeekDate = timeutil.Strftime(o.QueryWeekDate, "%Y-%m-%d")
	} else {
		QueryWeekDate = nil
	}
	
	return json.Marshal(&struct { 
		Schedule *Alternativeshiftschedulelookup `json:"schedule,omitempty"`
		
		QueryWeekDate *string `json:"queryWeekDate,omitempty"`
		Alias
	}{ 
		Schedule: o.Schedule,
		
		QueryWeekDate: QueryWeekDate,
		Alias:    (Alias)(o),
	})
}

func (o *Alternativeshiftoffersrequest) UnmarshalJSON(b []byte) error {
	var AlternativeshiftoffersrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AlternativeshiftoffersrequestMap)
	if err != nil {
		return err
	}
	
	if Schedule, ok := AlternativeshiftoffersrequestMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if queryWeekDateString, ok := AlternativeshiftoffersrequestMap["queryWeekDate"].(string); ok {
		QueryWeekDate, _ := time.Parse("2006-01-02", queryWeekDateString)
		o.QueryWeekDate = &QueryWeekDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Alternativeshiftoffersrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
