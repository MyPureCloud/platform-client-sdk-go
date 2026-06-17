package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulebidscheduledshift
type Schedulebidscheduledshift struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// WorkPlanShiftId - The ID of the work plan shift that was used in schedule generation
	WorkPlanShiftId *string `json:"workPlanShiftId,omitempty"`

	// WorkPlanId - The ID of the work plan from which the shift comes
	WorkPlanId *string `json:"workPlanId,omitempty"`

	// StartDate - The start date of the scheduled shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`

	// LengthMinutes - The length of the shift in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// Activities - The activities associated with this shift
	Activities *[]Schedulebidscheduledactivity `json:"activities,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Schedulebidscheduledshift) SetField(field string, fieldValue interface{}) {
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

func (o Schedulebidscheduledshift) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate", }
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
	type Alias Schedulebidscheduledshift
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		WorkPlanShiftId *string `json:"workPlanShiftId,omitempty"`
		
		WorkPlanId *string `json:"workPlanId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Activities *[]Schedulebidscheduledactivity `json:"activities,omitempty"`
		Alias
	}{ 
		WorkPlanShiftId: o.WorkPlanShiftId,
		
		WorkPlanId: o.WorkPlanId,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Activities: o.Activities,
		Alias:    (Alias)(o),
	})
}

func (o *Schedulebidscheduledshift) UnmarshalJSON(b []byte) error {
	var SchedulebidscheduledshiftMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulebidscheduledshiftMap)
	if err != nil {
		return err
	}
	
	if WorkPlanShiftId, ok := SchedulebidscheduledshiftMap["workPlanShiftId"].(string); ok {
		o.WorkPlanShiftId = &WorkPlanShiftId
	}
    
	if WorkPlanId, ok := SchedulebidscheduledshiftMap["workPlanId"].(string); ok {
		o.WorkPlanId = &WorkPlanId
	}
    
	if startDateString, ok := SchedulebidscheduledshiftMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := SchedulebidscheduledshiftMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Activities, ok := SchedulebidscheduledshiftMap["activities"].([]interface{}); ok {
		ActivitiesString, _ := json.Marshal(Activities)
		json.Unmarshal(ActivitiesString, &o.Activities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulebidscheduledshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
