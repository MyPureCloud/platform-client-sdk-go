package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createadmintimeoffrequest
type Createadmintimeoffrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Status - The status of this time off request
	Status *string `json:"status,omitempty"`

	// Users - A set of IDs for users to associate with this time off request
	Users *[]Userreference `json:"users,omitempty"`

	// ActivityCodeId - The ID of the activity code associated with this time off request. Activity code must be of the TimeOff category
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// Notes - Notes about the time off request
	Notes *string `json:"notes,omitempty"`

	// FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.
	FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`

	// PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests.
	PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`

	// DailyDurationMinutes - The daily duration of this time off request in minutes
	DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`

	// DurationMinutes - Daily durations for each day of this time off request in minutes
	DurationMinutes *[]int `json:"durationMinutes,omitempty"`

	// PayableMinutes - Payable minutes for each day of this time off request
	PayableMinutes *[]int `json:"payableMinutes,omitempty"`

	// Paid - Whether this is a paid time off request
	Paid *bool `json:"paid,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createadmintimeoffrequest) SetField(field string, fieldValue interface{}) {
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

func (o Createadmintimeoffrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createadmintimeoffrequest
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		Users *[]Userreference `json:"users,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		
		PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`
		
		DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`
		
		DurationMinutes *[]int `json:"durationMinutes,omitempty"`
		
		PayableMinutes *[]int `json:"payableMinutes,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		Alias
	}{ 
		Status: o.Status,
		
		Users: o.Users,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Notes: o.Notes,
		
		FullDayManagementUnitDates: o.FullDayManagementUnitDates,
		
		PartialDayStartDateTimes: o.PartialDayStartDateTimes,
		
		DailyDurationMinutes: o.DailyDurationMinutes,
		
		DurationMinutes: o.DurationMinutes,
		
		PayableMinutes: o.PayableMinutes,
		
		Paid: o.Paid,
		Alias:    (Alias)(o),
	})
}

func (o *Createadmintimeoffrequest) UnmarshalJSON(b []byte) error {
	var CreateadmintimeoffrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateadmintimeoffrequestMap)
	if err != nil {
		return err
	}
	
	if Status, ok := CreateadmintimeoffrequestMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Users, ok := CreateadmintimeoffrequestMap["users"].([]interface{}); ok {
		UsersString, _ := json.Marshal(Users)
		json.Unmarshal(UsersString, &o.Users)
	}
	
	if ActivityCodeId, ok := CreateadmintimeoffrequestMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Notes, ok := CreateadmintimeoffrequestMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if FullDayManagementUnitDates, ok := CreateadmintimeoffrequestMap["fullDayManagementUnitDates"].([]interface{}); ok {
		FullDayManagementUnitDatesString, _ := json.Marshal(FullDayManagementUnitDates)
		json.Unmarshal(FullDayManagementUnitDatesString, &o.FullDayManagementUnitDates)
	}
	
	if PartialDayStartDateTimes, ok := CreateadmintimeoffrequestMap["partialDayStartDateTimes"].([]interface{}); ok {
		PartialDayStartDateTimesString, _ := json.Marshal(PartialDayStartDateTimes)
		json.Unmarshal(PartialDayStartDateTimesString, &o.PartialDayStartDateTimes)
	}
	
	if DailyDurationMinutes, ok := CreateadmintimeoffrequestMap["dailyDurationMinutes"].(float64); ok {
		DailyDurationMinutesInt := int(DailyDurationMinutes)
		o.DailyDurationMinutes = &DailyDurationMinutesInt
	}
	
	if DurationMinutes, ok := CreateadmintimeoffrequestMap["durationMinutes"].([]interface{}); ok {
		DurationMinutesString, _ := json.Marshal(DurationMinutes)
		json.Unmarshal(DurationMinutesString, &o.DurationMinutes)
	}
	
	if PayableMinutes, ok := CreateadmintimeoffrequestMap["payableMinutes"].([]interface{}); ok {
		PayableMinutesString, _ := json.Marshal(PayableMinutes)
		json.Unmarshal(PayableMinutesString, &o.PayableMinutes)
	}
	
	if Paid, ok := CreateadmintimeoffrequestMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createadmintimeoffrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
