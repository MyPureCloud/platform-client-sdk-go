package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Modifiableruleproperties
type Modifiableruleproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - Name of the rule
	Name *string `json:"name,omitempty"`

	// Description - The description of the rule.
	Description *string `json:"description,omitempty"`

	// Enabled - Indicates if the rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Notifications - The alert notification types to trigger when alarm state changes as well as the users they will be sent to.
	Notifications *[]Alertnotification `json:"notifications,omitempty"`

	// SendExitingAlarmNotifications - Indicates if the alert will send a notification when it is closed.
	SendExitingAlarmNotifications *bool `json:"sendExitingAlarmNotifications,omitempty"`

	// WaitBetweenNotificationMs - The amount of time in milliseconds to wait between notification.
	WaitBetweenNotificationMs *int `json:"waitBetweenNotificationMs,omitempty"`

	// Conditions - The set of metric conditions that would trigger an alert.
	Conditions *Commonruleconditions `json:"conditions,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Modifiableruleproperties) SetField(field string, fieldValue interface{}) {
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

func (o Modifiableruleproperties) MarshalJSON() ([]byte, error) {
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
	type Alias Modifiableruleproperties
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Notifications *[]Alertnotification `json:"notifications,omitempty"`
		
		SendExitingAlarmNotifications *bool `json:"sendExitingAlarmNotifications,omitempty"`
		
		WaitBetweenNotificationMs *int `json:"waitBetweenNotificationMs,omitempty"`
		
		Conditions *Commonruleconditions `json:"conditions,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Enabled: o.Enabled,
		
		Notifications: o.Notifications,
		
		SendExitingAlarmNotifications: o.SendExitingAlarmNotifications,
		
		WaitBetweenNotificationMs: o.WaitBetweenNotificationMs,
		
		Conditions: o.Conditions,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Modifiableruleproperties) UnmarshalJSON(b []byte) error {
	var ModifiablerulepropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &ModifiablerulepropertiesMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ModifiablerulepropertiesMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ModifiablerulepropertiesMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ModifiablerulepropertiesMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Enabled, ok := ModifiablerulepropertiesMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Notifications, ok := ModifiablerulepropertiesMap["notifications"].([]interface{}); ok {
		NotificationsString, _ := json.Marshal(Notifications)
		json.Unmarshal(NotificationsString, &o.Notifications)
	}
	
	if SendExitingAlarmNotifications, ok := ModifiablerulepropertiesMap["sendExitingAlarmNotifications"].(bool); ok {
		o.SendExitingAlarmNotifications = &SendExitingAlarmNotifications
	}
    
	if WaitBetweenNotificationMs, ok := ModifiablerulepropertiesMap["waitBetweenNotificationMs"].(float64); ok {
		WaitBetweenNotificationMsInt := int(WaitBetweenNotificationMs)
		o.WaitBetweenNotificationMs = &WaitBetweenNotificationMsInt
	}
	
	if Conditions, ok := ModifiablerulepropertiesMap["conditions"].(map[string]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if SelfUri, ok := ModifiablerulepropertiesMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Modifiableruleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
