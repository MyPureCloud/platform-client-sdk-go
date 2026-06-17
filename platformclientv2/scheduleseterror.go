package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Scheduleseterror
type Scheduleseterror struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ErrorCode - Error code that indicates why schedule set optimization failed. At least one of workPlans or workPlanRotations is set if there is an error during optimization
	ErrorCode *string `json:"errorCode,omitempty"`

	// WorkPlans - Work plans involved in the optimization failure
	WorkPlans *[]Workplanreference `json:"workPlans,omitempty"`

	// WorkPlanRotations - Work plan rotations involved in the optimization failure
	WorkPlanRotations *[]Workplanrotationreference `json:"workPlanRotations,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Scheduleseterror) SetField(field string, fieldValue interface{}) {
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

func (o Scheduleseterror) MarshalJSON() ([]byte, error) {
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
	type Alias Scheduleseterror
	
	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		WorkPlans *[]Workplanreference `json:"workPlans,omitempty"`
		
		WorkPlanRotations *[]Workplanrotationreference `json:"workPlanRotations,omitempty"`
		Alias
	}{ 
		ErrorCode: o.ErrorCode,
		
		WorkPlans: o.WorkPlans,
		
		WorkPlanRotations: o.WorkPlanRotations,
		Alias:    (Alias)(o),
	})
}

func (o *Scheduleseterror) UnmarshalJSON(b []byte) error {
	var ScheduleseterrorMap map[string]interface{}
	err := json.Unmarshal(b, &ScheduleseterrorMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := ScheduleseterrorMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if WorkPlans, ok := ScheduleseterrorMap["workPlans"].([]interface{}); ok {
		WorkPlansString, _ := json.Marshal(WorkPlans)
		json.Unmarshal(WorkPlansString, &o.WorkPlans)
	}
	
	if WorkPlanRotations, ok := ScheduleseterrorMap["workPlanRotations"].([]interface{}); ok {
		WorkPlanRotationsString, _ := json.Marshal(WorkPlanRotations)
		json.Unmarshal(WorkPlanRotationsString, &o.WorkPlanRotations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scheduleseterror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
