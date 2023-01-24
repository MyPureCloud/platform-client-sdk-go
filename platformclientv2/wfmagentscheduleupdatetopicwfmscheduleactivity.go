package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmscheduleactivity
type Wfmagentscheduleupdatetopicwfmscheduleactivity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`

	// CountsAsPaidTime
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`

	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// TimeOffRequestId
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmagentscheduleupdatetopicwfmscheduleactivity) SetField(field string, fieldValue interface{}) {
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

func (o Wfmagentscheduleupdatetopicwfmscheduleactivity) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmagentscheduleupdatetopicwfmscheduleactivity
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		Description *string `json:"description,omitempty"`
		Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		StartDate: StartDate,
		
		CountsAsPaidTime: o.CountsAsPaidTime,
		
		LengthInMinutes: o.LengthInMinutes,
		
		TimeOffRequestId: o.TimeOffRequestId,
		
		Description: o.Description,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmagentscheduleupdatetopicwfmscheduleactivity) UnmarshalJSON(b []byte) error {
	var WfmagentscheduleupdatetopicwfmscheduleactivityMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentscheduleupdatetopicwfmscheduleactivityMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if startDateString, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if CountsAsPaidTime, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    
	if LengthInMinutes, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if TimeOffRequestId, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    
	if Description, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["description"].(string); ok {
		o.Description = &Description
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
