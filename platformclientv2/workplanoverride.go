package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanoverride
type Workplanoverride struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartDate - The start date in yyyy-MM-dd format of the work plan override. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartDate *time.Time `json:"startDate,omitempty"`

	// WeekCount - Number of weeks for the work plan override
	WeekCount *int `json:"weekCount,omitempty"`

	// WorkPlan - The work plan reference associated with this override
	WorkPlan *Workplanreference `json:"workPlan,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workplanoverride) SetField(field string, fieldValue interface{}) {
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

func (o Workplanoverride) MarshalJSON() ([]byte, error) {
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
	type Alias Workplanoverride
	
	StartDate := new(string)
	if o.StartDate != nil {
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%d")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		WorkPlan *Workplanreference `json:"workPlan,omitempty"`
		Alias
	}{ 
		StartDate: StartDate,
		
		WeekCount: o.WeekCount,
		
		WorkPlan: o.WorkPlan,
		Alias:    (Alias)(o),
	})
}

func (o *Workplanoverride) UnmarshalJSON(b []byte) error {
	var WorkplanoverrideMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanoverrideMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := WorkplanoverrideMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02", startDateString)
		o.StartDate = &StartDate
	}
	
	if WeekCount, ok := WorkplanoverrideMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if WorkPlan, ok := WorkplanoverrideMap["workPlan"].(map[string]interface{}); ok {
		WorkPlanString, _ := json.Marshal(WorkPlan)
		json.Unmarshal(WorkPlanString, &o.WorkPlan)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanoverride) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
