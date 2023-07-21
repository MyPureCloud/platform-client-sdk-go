package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Commonalert
type Commonalert struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// User - The user who created the rule that triggered the alert.
	User *Userreference `json:"user,omitempty"`

	// Rule - The properties of the rule that triggered the alert.
	Rule *Alertruleproperties `json:"rule,omitempty"`

	// Notifications - The collection of notification methods and the ids of users who were notified by those methods.
	Notifications *[]Alertnotification `json:"notifications,omitempty"`

	// DateStart - The timestamp of when the alert was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - The timestamp of when the alert ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// Active - Indicates if an alert is currently active.
	Active *bool `json:"active,omitempty"`

	// Unread - Indicates if an alert has not been read.
	Unread *bool `json:"unread,omitempty"`

	// WaitBetweenNotificationMs - The amount of time to wait between notification. Time is in milliseconds.
	WaitBetweenNotificationMs *int `json:"waitBetweenNotificationMs,omitempty"`

	// Muted - Flag indicating if the alert is in a muted state.
	Muted *bool `json:"muted,omitempty"`

	// Snoozed - Flag indicating if the alert is in a snoozed state.
	Snoozed *bool `json:"snoozed,omitempty"`

	// DateMutedUntil - Timestamp of when the mute status of the alert should end. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateMutedUntil *time.Time `json:"dateMutedUntil,omitempty"`

	// DateSnoozedUntil - Timestamp of when the snooze status of the alert should end. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateSnoozedUntil *time.Time `json:"dateSnoozedUntil,omitempty"`

	// Conditions - The conditions that make up the rule.
	Conditions *Commonruleconditions `json:"conditions,omitempty"`

	// ConversationId - The id of the conversation instance that caused the alert to trigger.
	ConversationId *string `json:"conversationId,omitempty"`

	// RuleUri
	RuleUri *string `json:"ruleUri,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Commonalert) SetField(field string, fieldValue interface{}) {
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

func (o Commonalert) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart","DateEnd","DateMutedUntil","DateSnoozedUntil", }
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
	type Alias Commonalert
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnd = nil
	}
	
	DateMutedUntil := new(string)
	if o.DateMutedUntil != nil {
		
		*DateMutedUntil = timeutil.Strftime(o.DateMutedUntil, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateMutedUntil = nil
	}
	
	DateSnoozedUntil := new(string)
	if o.DateSnoozedUntil != nil {
		
		*DateSnoozedUntil = timeutil.Strftime(o.DateSnoozedUntil, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateSnoozedUntil = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		Rule *Alertruleproperties `json:"rule,omitempty"`
		
		Notifications *[]Alertnotification `json:"notifications,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		Unread *bool `json:"unread,omitempty"`
		
		WaitBetweenNotificationMs *int `json:"waitBetweenNotificationMs,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Snoozed *bool `json:"snoozed,omitempty"`
		
		DateMutedUntil *string `json:"dateMutedUntil,omitempty"`
		
		DateSnoozedUntil *string `json:"dateSnoozedUntil,omitempty"`
		
		Conditions *Commonruleconditions `json:"conditions,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		RuleUri *string `json:"ruleUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		User: o.User,
		
		Rule: o.Rule,
		
		Notifications: o.Notifications,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		Active: o.Active,
		
		Unread: o.Unread,
		
		WaitBetweenNotificationMs: o.WaitBetweenNotificationMs,
		
		Muted: o.Muted,
		
		Snoozed: o.Snoozed,
		
		DateMutedUntil: DateMutedUntil,
		
		DateSnoozedUntil: DateSnoozedUntil,
		
		Conditions: o.Conditions,
		
		ConversationId: o.ConversationId,
		
		RuleUri: o.RuleUri,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Commonalert) UnmarshalJSON(b []byte) error {
	var CommonalertMap map[string]interface{}
	err := json.Unmarshal(b, &CommonalertMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CommonalertMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CommonalertMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if User, ok := CommonalertMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Rule, ok := CommonalertMap["rule"].(map[string]interface{}); ok {
		RuleString, _ := json.Marshal(Rule)
		json.Unmarshal(RuleString, &o.Rule)
	}
	
	if Notifications, ok := CommonalertMap["notifications"].([]interface{}); ok {
		NotificationsString, _ := json.Marshal(Notifications)
		json.Unmarshal(NotificationsString, &o.Notifications)
	}
	
	if dateStartString, ok := CommonalertMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := CommonalertMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if Active, ok := CommonalertMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if Unread, ok := CommonalertMap["unread"].(bool); ok {
		o.Unread = &Unread
	}
    
	if WaitBetweenNotificationMs, ok := CommonalertMap["waitBetweenNotificationMs"].(float64); ok {
		WaitBetweenNotificationMsInt := int(WaitBetweenNotificationMs)
		o.WaitBetweenNotificationMs = &WaitBetweenNotificationMsInt
	}
	
	if Muted, ok := CommonalertMap["muted"].(bool); ok {
		o.Muted = &Muted
	}
    
	if Snoozed, ok := CommonalertMap["snoozed"].(bool); ok {
		o.Snoozed = &Snoozed
	}
    
	if dateMutedUntilString, ok := CommonalertMap["dateMutedUntil"].(string); ok {
		DateMutedUntil, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateMutedUntilString)
		o.DateMutedUntil = &DateMutedUntil
	}
	
	if dateSnoozedUntilString, ok := CommonalertMap["dateSnoozedUntil"].(string); ok {
		DateSnoozedUntil, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateSnoozedUntilString)
		o.DateSnoozedUntil = &DateSnoozedUntil
	}
	
	if Conditions, ok := CommonalertMap["conditions"].(map[string]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if ConversationId, ok := CommonalertMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if RuleUri, ok := CommonalertMap["ruleUri"].(string); ok {
		o.RuleUri = &RuleUri
	}
    
	if SelfUri, ok := CommonalertMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Commonalert) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
