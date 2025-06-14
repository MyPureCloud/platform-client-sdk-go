package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Usersettingsforchat
type Usersettingsforchat struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Muted - Whether or not to enable muting notifications
	Muted *bool `json:"muted,omitempty"`

	// MentionsOnly - Whether or not to enable notifications for mentions only
	MentionsOnly *bool `json:"mentionsOnly,omitempty"`

	// NotifyOnReactions - Whether or not to enable notifications for reactions on a user's own messages
	NotifyOnReactions *bool `json:"notifyOnReactions,omitempty"`

	// Mobile - Settings for mobile devices
	Mobile *Mobilesettings `json:"mobile,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Usersettingsforchat) SetField(field string, fieldValue interface{}) {
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

func (o Usersettingsforchat) MarshalJSON() ([]byte, error) {
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
	type Alias Usersettingsforchat
	
	return json.Marshal(&struct { 
		Muted *bool `json:"muted,omitempty"`
		
		MentionsOnly *bool `json:"mentionsOnly,omitempty"`
		
		NotifyOnReactions *bool `json:"notifyOnReactions,omitempty"`
		
		Mobile *Mobilesettings `json:"mobile,omitempty"`
		Alias
	}{ 
		Muted: o.Muted,
		
		MentionsOnly: o.MentionsOnly,
		
		NotifyOnReactions: o.NotifyOnReactions,
		
		Mobile: o.Mobile,
		Alias:    (Alias)(o),
	})
}

func (o *Usersettingsforchat) UnmarshalJSON(b []byte) error {
	var UsersettingsforchatMap map[string]interface{}
	err := json.Unmarshal(b, &UsersettingsforchatMap)
	if err != nil {
		return err
	}
	
	if Muted, ok := UsersettingsforchatMap["muted"].(bool); ok {
		o.Muted = &Muted
	}
    
	if MentionsOnly, ok := UsersettingsforchatMap["mentionsOnly"].(bool); ok {
		o.MentionsOnly = &MentionsOnly
	}
    
	if NotifyOnReactions, ok := UsersettingsforchatMap["notifyOnReactions"].(bool); ok {
		o.NotifyOnReactions = &NotifyOnReactions
	}
    
	if Mobile, ok := UsersettingsforchatMap["mobile"].(map[string]interface{}); ok {
		MobileString, _ := json.Marshal(Mobile)
		json.Unmarshal(MobileString, &o.Mobile)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Usersettingsforchat) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
