package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradematchreviewuserresponse
type Shifttradematchreviewuserresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// WeeklyMinimumPaidMinutes - The minimum weekly paid minutes for this user per the work plan tied to the agent schedule
	WeeklyMinimumPaidMinutes *int `json:"weeklyMinimumPaidMinutes,omitempty"`

	// WeeklyMaximumPaidMinutes - The maximum weekly paid minutes for this user per the work plan tied to the agent schedule
	WeeklyMaximumPaidMinutes *int `json:"weeklyMaximumPaidMinutes,omitempty"`

	// PreTradeSchedulePaidMinutes - The paid minutes on the week schedule for this user prior to the shift trade
	PreTradeSchedulePaidMinutes *int `json:"preTradeSchedulePaidMinutes,omitempty"`

	// PostTradeSchedulePaidMinutes - The paid minutes on the week schedule for this user if the shift trade is approved
	PostTradeSchedulePaidMinutes *int `json:"postTradeSchedulePaidMinutes,omitempty"`

	// PostTradeNewShift - Preview of what the shift will look like for the opposite side of this trade after the match is approved
	PostTradeNewShift *Shifttradepreviewresponse `json:"postTradeNewShift,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Shifttradematchreviewuserresponse) SetField(field string, fieldValue interface{}) {
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

func (o Shifttradematchreviewuserresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Shifttradematchreviewuserresponse
	
	return json.Marshal(&struct { 
		WeeklyMinimumPaidMinutes *int `json:"weeklyMinimumPaidMinutes,omitempty"`
		
		WeeklyMaximumPaidMinutes *int `json:"weeklyMaximumPaidMinutes,omitempty"`
		
		PreTradeSchedulePaidMinutes *int `json:"preTradeSchedulePaidMinutes,omitempty"`
		
		PostTradeSchedulePaidMinutes *int `json:"postTradeSchedulePaidMinutes,omitempty"`
		
		PostTradeNewShift *Shifttradepreviewresponse `json:"postTradeNewShift,omitempty"`
		Alias
	}{ 
		WeeklyMinimumPaidMinutes: o.WeeklyMinimumPaidMinutes,
		
		WeeklyMaximumPaidMinutes: o.WeeklyMaximumPaidMinutes,
		
		PreTradeSchedulePaidMinutes: o.PreTradeSchedulePaidMinutes,
		
		PostTradeSchedulePaidMinutes: o.PostTradeSchedulePaidMinutes,
		
		PostTradeNewShift: o.PostTradeNewShift,
		Alias:    (Alias)(o),
	})
}

func (o *Shifttradematchreviewuserresponse) UnmarshalJSON(b []byte) error {
	var ShifttradematchreviewuserresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradematchreviewuserresponseMap)
	if err != nil {
		return err
	}
	
	if WeeklyMinimumPaidMinutes, ok := ShifttradematchreviewuserresponseMap["weeklyMinimumPaidMinutes"].(float64); ok {
		WeeklyMinimumPaidMinutesInt := int(WeeklyMinimumPaidMinutes)
		o.WeeklyMinimumPaidMinutes = &WeeklyMinimumPaidMinutesInt
	}
	
	if WeeklyMaximumPaidMinutes, ok := ShifttradematchreviewuserresponseMap["weeklyMaximumPaidMinutes"].(float64); ok {
		WeeklyMaximumPaidMinutesInt := int(WeeklyMaximumPaidMinutes)
		o.WeeklyMaximumPaidMinutes = &WeeklyMaximumPaidMinutesInt
	}
	
	if PreTradeSchedulePaidMinutes, ok := ShifttradematchreviewuserresponseMap["preTradeSchedulePaidMinutes"].(float64); ok {
		PreTradeSchedulePaidMinutesInt := int(PreTradeSchedulePaidMinutes)
		o.PreTradeSchedulePaidMinutes = &PreTradeSchedulePaidMinutesInt
	}
	
	if PostTradeSchedulePaidMinutes, ok := ShifttradematchreviewuserresponseMap["postTradeSchedulePaidMinutes"].(float64); ok {
		PostTradeSchedulePaidMinutesInt := int(PostTradeSchedulePaidMinutes)
		o.PostTradeSchedulePaidMinutes = &PostTradeSchedulePaidMinutesInt
	}
	
	if PostTradeNewShift, ok := ShifttradematchreviewuserresponseMap["postTradeNewShift"].(map[string]interface{}); ok {
		PostTradeNewShiftString, _ := json.Marshal(PostTradeNewShift)
		json.Unmarshal(PostTradeNewShiftString, &o.PostTradeNewShift)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradematchreviewuserresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
