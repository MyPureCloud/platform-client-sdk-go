package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentscheduleactivity
type Buagentscheduleactivity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartDate - The start date/time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`

	// LengthMinutes - The length of this activity in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// Description - The description of this activity
	Description *string `json:"description,omitempty"`

	// ActivityCodeId - The ID of the activity code associated with this activity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// Paid - Whether this activity is paid
	Paid *bool `json:"paid,omitempty"`

	// PayableMinutes - Payable minutes for this activity
	PayableMinutes *int `json:"payableMinutes,omitempty"`

	// TimeOffRequestId - The ID of the time off request associated with this activity, if applicable
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`

	// TimeOffRequestSyncVersion - The sync version of the partial day time off request for which the scheduled activity is associated, if applicable
	TimeOffRequestSyncVersion *int `json:"timeOffRequestSyncVersion,omitempty"`

	// ExternalActivityId - The ID of the external activity associated with this activity, if applicable
	ExternalActivityId *string `json:"externalActivityId,omitempty"`

	// ExternalActivityType - The type of the external activity associated with this activity, if applicable
	ExternalActivityType *string `json:"externalActivityType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buagentscheduleactivity) SetField(field string, fieldValue interface{}) {
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

func (o Buagentscheduleactivity) MarshalJSON() ([]byte, error) {
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
	type Alias Buagentscheduleactivity
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		PayableMinutes *int `json:"payableMinutes,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		TimeOffRequestSyncVersion *int `json:"timeOffRequestSyncVersion,omitempty"`
		
		ExternalActivityId *string `json:"externalActivityId,omitempty"`
		
		ExternalActivityType *string `json:"externalActivityType,omitempty"`
		Alias
	}{ 
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Description: o.Description,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Paid: o.Paid,
		
		PayableMinutes: o.PayableMinutes,
		
		TimeOffRequestId: o.TimeOffRequestId,
		
		TimeOffRequestSyncVersion: o.TimeOffRequestSyncVersion,
		
		ExternalActivityId: o.ExternalActivityId,
		
		ExternalActivityType: o.ExternalActivityType,
		Alias:    (Alias)(o),
	})
}

func (o *Buagentscheduleactivity) UnmarshalJSON(b []byte) error {
	var BuagentscheduleactivityMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentscheduleactivityMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := BuagentscheduleactivityMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := BuagentscheduleactivityMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Description, ok := BuagentscheduleactivityMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ActivityCodeId, ok := BuagentscheduleactivityMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Paid, ok := BuagentscheduleactivityMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
    
	if PayableMinutes, ok := BuagentscheduleactivityMap["payableMinutes"].(float64); ok {
		PayableMinutesInt := int(PayableMinutes)
		o.PayableMinutes = &PayableMinutesInt
	}
	
	if TimeOffRequestId, ok := BuagentscheduleactivityMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    
	if TimeOffRequestSyncVersion, ok := BuagentscheduleactivityMap["timeOffRequestSyncVersion"].(float64); ok {
		TimeOffRequestSyncVersionInt := int(TimeOffRequestSyncVersion)
		o.TimeOffRequestSyncVersion = &TimeOffRequestSyncVersionInt
	}
	
	if ExternalActivityId, ok := BuagentscheduleactivityMap["externalActivityId"].(string); ok {
		o.ExternalActivityId = &ExternalActivityId
	}
    
	if ExternalActivityType, ok := BuagentscheduleactivityMap["externalActivityType"].(string); ok {
		o.ExternalActivityType = &ExternalActivityType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
