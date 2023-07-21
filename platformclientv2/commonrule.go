package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Commonrule
type Commonrule struct { 
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

	// VarType - The type of the rule.
	VarType *string `json:"type,omitempty"`

	// InAlarm - Indicates if the rule is in alarm state.
	InAlarm *bool `json:"inAlarm,omitempty"`

	// User - The entity that created the rule.
	User *Userreference `json:"user,omitempty"`

	// Version - The current version number of the rule.
	Version *int `json:"version,omitempty"`

	// DateCreated - The creation date of the rule when the rule was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateLastModified - The timestamp of the last update to the rule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLastModified *time.Time `json:"dateLastModified,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Commonrule) SetField(field string, fieldValue interface{}) {
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

func (o Commonrule) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateLastModified", }
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
	type Alias Commonrule
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateLastModified := new(string)
	if o.DateLastModified != nil {
		
		*DateLastModified = timeutil.Strftime(o.DateLastModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Notifications *[]Alertnotification `json:"notifications,omitempty"`
		
		SendExitingAlarmNotifications *bool `json:"sendExitingAlarmNotifications,omitempty"`
		
		WaitBetweenNotificationMs *int `json:"waitBetweenNotificationMs,omitempty"`
		
		Conditions *Commonruleconditions `json:"conditions,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		InAlarm *bool `json:"inAlarm,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateLastModified *string `json:"dateLastModified,omitempty"`
		
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
		
		VarType: o.VarType,
		
		InAlarm: o.InAlarm,
		
		User: o.User,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateLastModified: DateLastModified,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Commonrule) UnmarshalJSON(b []byte) error {
	var CommonruleMap map[string]interface{}
	err := json.Unmarshal(b, &CommonruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CommonruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CommonruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := CommonruleMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Enabled, ok := CommonruleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Notifications, ok := CommonruleMap["notifications"].([]interface{}); ok {
		NotificationsString, _ := json.Marshal(Notifications)
		json.Unmarshal(NotificationsString, &o.Notifications)
	}
	
	if SendExitingAlarmNotifications, ok := CommonruleMap["sendExitingAlarmNotifications"].(bool); ok {
		o.SendExitingAlarmNotifications = &SendExitingAlarmNotifications
	}
    
	if WaitBetweenNotificationMs, ok := CommonruleMap["waitBetweenNotificationMs"].(float64); ok {
		WaitBetweenNotificationMsInt := int(WaitBetweenNotificationMs)
		o.WaitBetweenNotificationMs = &WaitBetweenNotificationMsInt
	}
	
	if Conditions, ok := CommonruleMap["conditions"].(map[string]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if VarType, ok := CommonruleMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if InAlarm, ok := CommonruleMap["inAlarm"].(bool); ok {
		o.InAlarm = &InAlarm
	}
    
	if User, ok := CommonruleMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Version, ok := CommonruleMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := CommonruleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateLastModifiedString, ok := CommonruleMap["dateLastModified"].(string); ok {
		DateLastModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastModifiedString)
		o.DateLastModified = &DateLastModified
	}
	
	if SelfUri, ok := CommonruleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Commonrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
