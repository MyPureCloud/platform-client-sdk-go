package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulegenerationwarning
type Schedulegenerationwarning struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UserId - ID of the user in the warning
	UserId *string `json:"userId,omitempty"`

	// UserNotLicensed - Whether the user does not have the appropriate license to be scheduled
	UserNotLicensed *bool `json:"userNotLicensed,omitempty"`

	// UnableToMeetMaxDays - Whether the number of scheduled days exceeded the maximum days to schedule defined in the agent work plan
	UnableToMeetMaxDays *bool `json:"unableToMeetMaxDays,omitempty"`

	// UnableToScheduleRequiredDays - Days indicated as required to work in agent work plan where no viable shift was found to schedule
	UnableToScheduleRequiredDays *[]string `json:"unableToScheduleRequiredDays,omitempty"`

	// UnableToMeetMinPaidForTheWeek - Whether the schedule did not meet the minimum paid time for the week defined in the agent work plan
	UnableToMeetMinPaidForTheWeek *bool `json:"unableToMeetMinPaidForTheWeek,omitempty"`

	// UnableToMeetMaxPaidForTheWeek - Whether the schedule exceeded the maximum paid time for the week defined in the agent work plan
	UnableToMeetMaxPaidForTheWeek *bool `json:"unableToMeetMaxPaidForTheWeek,omitempty"`

	// NoNeedDays - Days agent was scheduled but there was no need to meet. The scheduled days have no effect on service levels
	NoNeedDays *[]string `json:"noNeedDays,omitempty"`

	// ShiftsTooCloseTogether - Whether the schedule did not meet the minimum time between shifts defined in the agent work plan
	ShiftsTooCloseTogether *bool `json:"shiftsTooCloseTogether,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Schedulegenerationwarning) SetField(field string, fieldValue interface{}) {
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

func (o Schedulegenerationwarning) MarshalJSON() ([]byte, error) {
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
	type Alias Schedulegenerationwarning
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		UserNotLicensed *bool `json:"userNotLicensed,omitempty"`
		
		UnableToMeetMaxDays *bool `json:"unableToMeetMaxDays,omitempty"`
		
		UnableToScheduleRequiredDays *[]string `json:"unableToScheduleRequiredDays,omitempty"`
		
		UnableToMeetMinPaidForTheWeek *bool `json:"unableToMeetMinPaidForTheWeek,omitempty"`
		
		UnableToMeetMaxPaidForTheWeek *bool `json:"unableToMeetMaxPaidForTheWeek,omitempty"`
		
		NoNeedDays *[]string `json:"noNeedDays,omitempty"`
		
		ShiftsTooCloseTogether *bool `json:"shiftsTooCloseTogether,omitempty"`
		Alias
	}{ 
		UserId: o.UserId,
		
		UserNotLicensed: o.UserNotLicensed,
		
		UnableToMeetMaxDays: o.UnableToMeetMaxDays,
		
		UnableToScheduleRequiredDays: o.UnableToScheduleRequiredDays,
		
		UnableToMeetMinPaidForTheWeek: o.UnableToMeetMinPaidForTheWeek,
		
		UnableToMeetMaxPaidForTheWeek: o.UnableToMeetMaxPaidForTheWeek,
		
		NoNeedDays: o.NoNeedDays,
		
		ShiftsTooCloseTogether: o.ShiftsTooCloseTogether,
		Alias:    (Alias)(o),
	})
}

func (o *Schedulegenerationwarning) UnmarshalJSON(b []byte) error {
	var SchedulegenerationwarningMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulegenerationwarningMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := SchedulegenerationwarningMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if UserNotLicensed, ok := SchedulegenerationwarningMap["userNotLicensed"].(bool); ok {
		o.UserNotLicensed = &UserNotLicensed
	}
    
	if UnableToMeetMaxDays, ok := SchedulegenerationwarningMap["unableToMeetMaxDays"].(bool); ok {
		o.UnableToMeetMaxDays = &UnableToMeetMaxDays
	}
    
	if UnableToScheduleRequiredDays, ok := SchedulegenerationwarningMap["unableToScheduleRequiredDays"].([]interface{}); ok {
		UnableToScheduleRequiredDaysString, _ := json.Marshal(UnableToScheduleRequiredDays)
		json.Unmarshal(UnableToScheduleRequiredDaysString, &o.UnableToScheduleRequiredDays)
	}
	
	if UnableToMeetMinPaidForTheWeek, ok := SchedulegenerationwarningMap["unableToMeetMinPaidForTheWeek"].(bool); ok {
		o.UnableToMeetMinPaidForTheWeek = &UnableToMeetMinPaidForTheWeek
	}
    
	if UnableToMeetMaxPaidForTheWeek, ok := SchedulegenerationwarningMap["unableToMeetMaxPaidForTheWeek"].(bool); ok {
		o.UnableToMeetMaxPaidForTheWeek = &UnableToMeetMaxPaidForTheWeek
	}
    
	if NoNeedDays, ok := SchedulegenerationwarningMap["noNeedDays"].([]interface{}); ok {
		NoNeedDaysString, _ := json.Marshal(NoNeedDays)
		json.Unmarshal(NoNeedDaysString, &o.NoNeedDays)
	}
	
	if ShiftsTooCloseTogether, ok := SchedulegenerationwarningMap["shiftsTooCloseTogether"].(bool); ok {
		o.ShiftsTooCloseTogether = &ShiftsTooCloseTogether
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulegenerationwarning) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
