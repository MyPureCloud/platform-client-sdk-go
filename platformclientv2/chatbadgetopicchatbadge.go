package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatbadgetopicchatbadge
type Chatbadgetopicchatbadge struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Entity
	Entity *Chatbadgetopicbadgeentity `json:"entity,omitempty"`

	// UnreadCount
	UnreadCount *int `json:"unreadCount,omitempty"`

	// LastUnreadNotificationDate
	LastUnreadNotificationDate *time.Time `json:"lastUnreadNotificationDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Chatbadgetopicchatbadge) SetField(field string, fieldValue interface{}) {
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

func (o Chatbadgetopicchatbadge) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "LastUnreadNotificationDate", }
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
	type Alias Chatbadgetopicchatbadge
	
	LastUnreadNotificationDate := new(string)
	if o.LastUnreadNotificationDate != nil {
		
		*LastUnreadNotificationDate = timeutil.Strftime(o.LastUnreadNotificationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastUnreadNotificationDate = nil
	}
	
	return json.Marshal(&struct { 
		Entity *Chatbadgetopicbadgeentity `json:"entity,omitempty"`
		
		UnreadCount *int `json:"unreadCount,omitempty"`
		
		LastUnreadNotificationDate *string `json:"lastUnreadNotificationDate,omitempty"`
		Alias
	}{ 
		Entity: o.Entity,
		
		UnreadCount: o.UnreadCount,
		
		LastUnreadNotificationDate: LastUnreadNotificationDate,
		Alias:    (Alias)(o),
	})
}

func (o *Chatbadgetopicchatbadge) UnmarshalJSON(b []byte) error {
	var ChatbadgetopicchatbadgeMap map[string]interface{}
	err := json.Unmarshal(b, &ChatbadgetopicchatbadgeMap)
	if err != nil {
		return err
	}
	
	if Entity, ok := ChatbadgetopicchatbadgeMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if UnreadCount, ok := ChatbadgetopicchatbadgeMap["unreadCount"].(float64); ok {
		UnreadCountInt := int(UnreadCount)
		o.UnreadCount = &UnreadCountInt
	}
	
	if lastUnreadNotificationDateString, ok := ChatbadgetopicchatbadgeMap["lastUnreadNotificationDate"].(string); ok {
		LastUnreadNotificationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", lastUnreadNotificationDateString)
		o.LastUnreadNotificationDate = &LastUnreadNotificationDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chatbadgetopicchatbadge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
