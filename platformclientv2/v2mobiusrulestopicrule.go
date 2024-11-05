package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusrulestopicrule
type V2mobiusrulestopicrule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// UserId
	UserId *string `json:"userId,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// Notifications
	Notifications *[]V2mobiusrulestopicalertnotification `json:"notifications,omitempty"`

	// Conditions
	Conditions *V2mobiusrulestopiccondition `json:"conditions,omitempty"`

	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

	// InAlarm
	InAlarm *bool `json:"inAlarm,omitempty"`

	// Action
	Action *string `json:"action,omitempty"`

	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// SendExitingAlarmNotification
	SendExitingAlarmNotification *bool `json:"sendExitingAlarmNotification,omitempty"`

	// WaitBetweenNotificationMs
	WaitBetweenNotificationMs *int `json:"waitBetweenNotificationMs,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2mobiusrulestopicrule) SetField(field string, fieldValue interface{}) {
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

func (o V2mobiusrulestopicrule) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias V2mobiusrulestopicrule
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Notifications *[]V2mobiusrulestopicalertnotification `json:"notifications,omitempty"`
		
		Conditions *V2mobiusrulestopiccondition `json:"conditions,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		InAlarm *bool `json:"inAlarm,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		SendExitingAlarmNotification *bool `json:"sendExitingAlarmNotification,omitempty"`
		
		WaitBetweenNotificationMs *int `json:"waitBetweenNotificationMs,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		UserId: o.UserId,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Notifications: o.Notifications,
		
		Conditions: o.Conditions,
		
		Enabled: o.Enabled,
		
		InAlarm: o.InAlarm,
		
		Action: o.Action,
		
		DateCreated: DateCreated,
		
		SendExitingAlarmNotification: o.SendExitingAlarmNotification,
		
		WaitBetweenNotificationMs: o.WaitBetweenNotificationMs,
		Alias:    (Alias)(o),
	})
}

func (o *V2mobiusrulestopicrule) UnmarshalJSON(b []byte) error {
	var V2mobiusrulestopicruleMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusrulestopicruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2mobiusrulestopicruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if UserId, ok := V2mobiusrulestopicruleMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Name, ok := V2mobiusrulestopicruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := V2mobiusrulestopicruleMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Notifications, ok := V2mobiusrulestopicruleMap["notifications"].([]interface{}); ok {
		NotificationsString, _ := json.Marshal(Notifications)
		json.Unmarshal(NotificationsString, &o.Notifications)
	}
	
	if Conditions, ok := V2mobiusrulestopicruleMap["conditions"].(map[string]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if Enabled, ok := V2mobiusrulestopicruleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if InAlarm, ok := V2mobiusrulestopicruleMap["inAlarm"].(bool); ok {
		o.InAlarm = &InAlarm
	}
    
	if Action, ok := V2mobiusrulestopicruleMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if dateCreatedString, ok := V2mobiusrulestopicruleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if SendExitingAlarmNotification, ok := V2mobiusrulestopicruleMap["sendExitingAlarmNotification"].(bool); ok {
		o.SendExitingAlarmNotification = &SendExitingAlarmNotification
	}
    
	if WaitBetweenNotificationMs, ok := V2mobiusrulestopicruleMap["waitBetweenNotificationMs"].(float64); ok {
		WaitBetweenNotificationMsInt := int(WaitBetweenNotificationMs)
		o.WaitBetweenNotificationMs = &WaitBetweenNotificationMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusrulestopicrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
