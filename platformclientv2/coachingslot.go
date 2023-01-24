package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingslot
type Coachingslot struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateStart - Start date and time of scheduled coaching appointment slot. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// LengthInMinutes - Length of coaching appointment slot in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// StaffingDifference - Difference between scheduled and forecast headcount for this slot after scheduling the coaching appointment
	StaffingDifference *float64 `json:"staffingDifference,omitempty"`

	// DifferenceRating - Rating based on the staffing difference for scheduled slot
	DifferenceRating *string `json:"differenceRating,omitempty"`

	// WfmSchedule - Workforce Management schedule information associated with the slot
	WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Coachingslot) SetField(field string, fieldValue interface{}) {
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

func (o Coachingslot) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Coachingslot
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		StaffingDifference *float64 `json:"staffingDifference,omitempty"`
		
		DifferenceRating *string `json:"differenceRating,omitempty"`
		
		WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`
		Alias
	}{ 
		DateStart: DateStart,
		
		LengthInMinutes: o.LengthInMinutes,
		
		StaffingDifference: o.StaffingDifference,
		
		DifferenceRating: o.DifferenceRating,
		
		WfmSchedule: o.WfmSchedule,
		Alias:    (Alias)(o),
	})
}

func (o *Coachingslot) UnmarshalJSON(b []byte) error {
	var CoachingslotMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingslotMap)
	if err != nil {
		return err
	}
	
	if dateStartString, ok := CoachingslotMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthInMinutes, ok := CoachingslotMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if StaffingDifference, ok := CoachingslotMap["staffingDifference"].(float64); ok {
		o.StaffingDifference = &StaffingDifference
	}
    
	if DifferenceRating, ok := CoachingslotMap["differenceRating"].(string); ok {
		o.DifferenceRating = &DifferenceRating
	}
    
	if WfmSchedule, ok := CoachingslotMap["wfmSchedule"].(map[string]interface{}); ok {
		WfmScheduleString, _ := json.Marshal(WfmSchedule)
		json.Unmarshal(WfmScheduleString, &o.WfmSchedule)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
