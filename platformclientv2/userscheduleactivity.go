package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userscheduleactivity
type Userscheduleactivity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActivityCodeId - The id for the activity code.  Look up a map of activity codes with the activities route
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// StartDate - Start time in UTC for this activity, in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`

	// LengthInMinutes - Length in minutes for this activity
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// Description - Description for this activity
	Description *string `json:"description,omitempty"`

	// CountsAsPaidTime - Whether this activity is paid
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`

	// IsDstFallback - Whether this activity spans a DST fallback
	IsDstFallback *bool `json:"isDstFallback,omitempty"`

	// TimeOffRequestId - Time off request id of this activity
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userscheduleactivity) SetField(field string, fieldValue interface{}) {
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

func (o Userscheduleactivity) MarshalJSON() ([]byte, error) {
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
	type Alias Userscheduleactivity
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		IsDstFallback *bool `json:"isDstFallback,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		StartDate: StartDate,
		
		LengthInMinutes: o.LengthInMinutes,
		
		Description: o.Description,
		
		CountsAsPaidTime: o.CountsAsPaidTime,
		
		IsDstFallback: o.IsDstFallback,
		
		TimeOffRequestId: o.TimeOffRequestId,
		Alias:    (Alias)(o),
	})
}

func (o *Userscheduleactivity) UnmarshalJSON(b []byte) error {
	var UserscheduleactivityMap map[string]interface{}
	err := json.Unmarshal(b, &UserscheduleactivityMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := UserscheduleactivityMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if startDateString, ok := UserscheduleactivityMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthInMinutes, ok := UserscheduleactivityMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if Description, ok := UserscheduleactivityMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if CountsAsPaidTime, ok := UserscheduleactivityMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    
	if IsDstFallback, ok := UserscheduleactivityMap["isDstFallback"].(bool); ok {
		o.IsDstFallback = &IsDstFallback
	}
    
	if TimeOffRequestId, ok := UserscheduleactivityMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
