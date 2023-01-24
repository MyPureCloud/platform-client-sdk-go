package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createworkplanactivity
type Createworkplanactivity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActivityCodeId - ID of the activity code associated with this activity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// Description - Description of the activity
	Description *string `json:"description,omitempty"`

	// LengthMinutes - Length of the activity in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// StartTimeIsRelativeToShiftStart - Whether the start time of the activity is relative to the start time of the shift it belongs to
	StartTimeIsRelativeToShiftStart *bool `json:"startTimeIsRelativeToShiftStart,omitempty"`

	// FlexibleStartTime - Whether the start time of the activity is flexible
	FlexibleStartTime *bool `json:"flexibleStartTime,omitempty"`

	// EarliestStartTimeMinutes - Earliest activity start in offset minutes relative to shift start time if startTimeIsRelativeToShiftStart == true else its based on midnight. Used if flexibleStartTime == true
	EarliestStartTimeMinutes *int `json:"earliestStartTimeMinutes,omitempty"`

	// LatestStartTimeMinutes - Latest activity start in offset minutes relative to shift start time if startTimeIsRelativeToShiftStart == true else its based on midnight. Used if flexibleStartTime == true
	LatestStartTimeMinutes *int `json:"latestStartTimeMinutes,omitempty"`

	// ExactStartTimeMinutes - Exact activity start in offset minutes relative to shift start time if startTimeIsRelativeToShiftStart == true else its based on midnight. Used if flexibleStartTime == false
	ExactStartTimeMinutes *int `json:"exactStartTimeMinutes,omitempty"`

	// StartTimeIncrementMinutes - Increment in offset minutes that would contribute to different possible start times for the activity
	StartTimeIncrementMinutes *int `json:"startTimeIncrementMinutes,omitempty"`

	// CountsAsPaidTime - Whether the activity is paid
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`

	// CountsAsContiguousWorkTime - Whether the activity duration is counted towards contiguous work time
	CountsAsContiguousWorkTime *bool `json:"countsAsContiguousWorkTime,omitempty"`

	// MinimumLengthFromShiftStartMinutes - The minimum duration between shift start and shift item (e.g., break or meal) start in minutes
	MinimumLengthFromShiftStartMinutes *int `json:"minimumLengthFromShiftStartMinutes,omitempty"`

	// MinimumLengthFromShiftEndMinutes - The minimum duration between shift item (e.g., break or meal) end and shift end in minutes
	MinimumLengthFromShiftEndMinutes *int `json:"minimumLengthFromShiftEndMinutes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createworkplanactivity) SetField(field string, fieldValue interface{}) {
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

func (o Createworkplanactivity) MarshalJSON() ([]byte, error) {
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
	type Alias Createworkplanactivity
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		StartTimeIsRelativeToShiftStart *bool `json:"startTimeIsRelativeToShiftStart,omitempty"`
		
		FlexibleStartTime *bool `json:"flexibleStartTime,omitempty"`
		
		EarliestStartTimeMinutes *int `json:"earliestStartTimeMinutes,omitempty"`
		
		LatestStartTimeMinutes *int `json:"latestStartTimeMinutes,omitempty"`
		
		ExactStartTimeMinutes *int `json:"exactStartTimeMinutes,omitempty"`
		
		StartTimeIncrementMinutes *int `json:"startTimeIncrementMinutes,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		CountsAsContiguousWorkTime *bool `json:"countsAsContiguousWorkTime,omitempty"`
		
		MinimumLengthFromShiftStartMinutes *int `json:"minimumLengthFromShiftStartMinutes,omitempty"`
		
		MinimumLengthFromShiftEndMinutes *int `json:"minimumLengthFromShiftEndMinutes,omitempty"`
		Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		Description: o.Description,
		
		LengthMinutes: o.LengthMinutes,
		
		StartTimeIsRelativeToShiftStart: o.StartTimeIsRelativeToShiftStart,
		
		FlexibleStartTime: o.FlexibleStartTime,
		
		EarliestStartTimeMinutes: o.EarliestStartTimeMinutes,
		
		LatestStartTimeMinutes: o.LatestStartTimeMinutes,
		
		ExactStartTimeMinutes: o.ExactStartTimeMinutes,
		
		StartTimeIncrementMinutes: o.StartTimeIncrementMinutes,
		
		CountsAsPaidTime: o.CountsAsPaidTime,
		
		CountsAsContiguousWorkTime: o.CountsAsContiguousWorkTime,
		
		MinimumLengthFromShiftStartMinutes: o.MinimumLengthFromShiftStartMinutes,
		
		MinimumLengthFromShiftEndMinutes: o.MinimumLengthFromShiftEndMinutes,
		Alias:    (Alias)(o),
	})
}

func (o *Createworkplanactivity) UnmarshalJSON(b []byte) error {
	var CreateworkplanactivityMap map[string]interface{}
	err := json.Unmarshal(b, &CreateworkplanactivityMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := CreateworkplanactivityMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Description, ok := CreateworkplanactivityMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if LengthMinutes, ok := CreateworkplanactivityMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if StartTimeIsRelativeToShiftStart, ok := CreateworkplanactivityMap["startTimeIsRelativeToShiftStart"].(bool); ok {
		o.StartTimeIsRelativeToShiftStart = &StartTimeIsRelativeToShiftStart
	}
    
	if FlexibleStartTime, ok := CreateworkplanactivityMap["flexibleStartTime"].(bool); ok {
		o.FlexibleStartTime = &FlexibleStartTime
	}
    
	if EarliestStartTimeMinutes, ok := CreateworkplanactivityMap["earliestStartTimeMinutes"].(float64); ok {
		EarliestStartTimeMinutesInt := int(EarliestStartTimeMinutes)
		o.EarliestStartTimeMinutes = &EarliestStartTimeMinutesInt
	}
	
	if LatestStartTimeMinutes, ok := CreateworkplanactivityMap["latestStartTimeMinutes"].(float64); ok {
		LatestStartTimeMinutesInt := int(LatestStartTimeMinutes)
		o.LatestStartTimeMinutes = &LatestStartTimeMinutesInt
	}
	
	if ExactStartTimeMinutes, ok := CreateworkplanactivityMap["exactStartTimeMinutes"].(float64); ok {
		ExactStartTimeMinutesInt := int(ExactStartTimeMinutes)
		o.ExactStartTimeMinutes = &ExactStartTimeMinutesInt
	}
	
	if StartTimeIncrementMinutes, ok := CreateworkplanactivityMap["startTimeIncrementMinutes"].(float64); ok {
		StartTimeIncrementMinutesInt := int(StartTimeIncrementMinutes)
		o.StartTimeIncrementMinutes = &StartTimeIncrementMinutesInt
	}
	
	if CountsAsPaidTime, ok := CreateworkplanactivityMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    
	if CountsAsContiguousWorkTime, ok := CreateworkplanactivityMap["countsAsContiguousWorkTime"].(bool); ok {
		o.CountsAsContiguousWorkTime = &CountsAsContiguousWorkTime
	}
    
	if MinimumLengthFromShiftStartMinutes, ok := CreateworkplanactivityMap["minimumLengthFromShiftStartMinutes"].(float64); ok {
		MinimumLengthFromShiftStartMinutesInt := int(MinimumLengthFromShiftStartMinutes)
		o.MinimumLengthFromShiftStartMinutes = &MinimumLengthFromShiftStartMinutesInt
	}
	
	if MinimumLengthFromShiftEndMinutes, ok := CreateworkplanactivityMap["minimumLengthFromShiftEndMinutes"].(float64); ok {
		MinimumLengthFromShiftEndMinutesInt := int(MinimumLengthFromShiftEndMinutes)
		o.MinimumLengthFromShiftEndMinutes = &MinimumLengthFromShiftEndMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createworkplanactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
