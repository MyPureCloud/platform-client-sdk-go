package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulingperiod
type Schedulingperiod struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EarliestStartDate - The earliest date the associated activity plan can begin, in YYYY-MM-DD format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EarliestStartDate *time.Time `json:"earliestStartDate,omitempty"`

	// LatestEndDate - The latest date the associated activity plan can end, in YYYY-MM-DD format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	LatestEndDate *time.Time `json:"latestEndDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Schedulingperiod) SetField(field string, fieldValue interface{}) {
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

func (o Schedulingperiod) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "EarliestStartDate","LatestEndDate", }

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
	type Alias Schedulingperiod
	
	EarliestStartDate := new(string)
	if o.EarliestStartDate != nil {
		*EarliestStartDate = timeutil.Strftime(o.EarliestStartDate, "%Y-%m-%d")
	} else {
		EarliestStartDate = nil
	}
	
	LatestEndDate := new(string)
	if o.LatestEndDate != nil {
		*LatestEndDate = timeutil.Strftime(o.LatestEndDate, "%Y-%m-%d")
	} else {
		LatestEndDate = nil
	}
	
	return json.Marshal(&struct { 
		EarliestStartDate *string `json:"earliestStartDate,omitempty"`
		
		LatestEndDate *string `json:"latestEndDate,omitempty"`
		Alias
	}{ 
		EarliestStartDate: EarliestStartDate,
		
		LatestEndDate: LatestEndDate,
		Alias:    (Alias)(o),
	})
}

func (o *Schedulingperiod) UnmarshalJSON(b []byte) error {
	var SchedulingperiodMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulingperiodMap)
	if err != nil {
		return err
	}
	
	if earliestStartDateString, ok := SchedulingperiodMap["earliestStartDate"].(string); ok {
		EarliestStartDate, _ := time.Parse("2006-01-02", earliestStartDateString)
		o.EarliestStartDate = &EarliestStartDate
	}
	
	if latestEndDateString, ok := SchedulingperiodMap["latestEndDate"].(string); ok {
		LatestEndDate, _ := time.Parse("2006-01-02", latestEndDateString)
		o.LatestEndDate = &LatestEndDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulingperiod) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
