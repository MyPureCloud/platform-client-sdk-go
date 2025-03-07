package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messageevent - Message event element.  Examples include: system messages, typing indicators, cobrowse offerings.
type Messageevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventType - Type of this event element
	EventType *string `json:"eventType,omitempty"`

	// CoBrowse - CoBrowse event.
	CoBrowse *Eventcobrowse `json:"coBrowse,omitempty"`

	// Typing - Typing event.
	Typing *Eventtyping `json:"typing,omitempty"`

	// Presence - Presence event.
	Presence *Eventpresence `json:"presence,omitempty"`

	// Video - Video event.
	Video *Eventvideo `json:"video,omitempty"`

	// Reactions - A list of reactions to a message.
	Reactions *[]Contentreaction `json:"reactions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messageevent) SetField(field string, fieldValue interface{}) {
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

func (o Messageevent) MarshalJSON() ([]byte, error) {
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
	type Alias Messageevent
	
	return json.Marshal(&struct { 
		EventType *string `json:"eventType,omitempty"`
		
		CoBrowse *Eventcobrowse `json:"coBrowse,omitempty"`
		
		Typing *Eventtyping `json:"typing,omitempty"`
		
		Presence *Eventpresence `json:"presence,omitempty"`
		
		Video *Eventvideo `json:"video,omitempty"`
		
		Reactions *[]Contentreaction `json:"reactions,omitempty"`
		Alias
	}{ 
		EventType: o.EventType,
		
		CoBrowse: o.CoBrowse,
		
		Typing: o.Typing,
		
		Presence: o.Presence,
		
		Video: o.Video,
		
		Reactions: o.Reactions,
		Alias:    (Alias)(o),
	})
}

func (o *Messageevent) UnmarshalJSON(b []byte) error {
	var MessageeventMap map[string]interface{}
	err := json.Unmarshal(b, &MessageeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := MessageeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if CoBrowse, ok := MessageeventMap["coBrowse"].(map[string]interface{}); ok {
		CoBrowseString, _ := json.Marshal(CoBrowse)
		json.Unmarshal(CoBrowseString, &o.CoBrowse)
	}
	
	if Typing, ok := MessageeventMap["typing"].(map[string]interface{}); ok {
		TypingString, _ := json.Marshal(Typing)
		json.Unmarshal(TypingString, &o.Typing)
	}
	
	if Presence, ok := MessageeventMap["presence"].(map[string]interface{}); ok {
		PresenceString, _ := json.Marshal(Presence)
		json.Unmarshal(PresenceString, &o.Presence)
	}
	
	if Video, ok := MessageeventMap["video"].(map[string]interface{}); ok {
		VideoString, _ := json.Marshal(Video)
		json.Unmarshal(VideoString, &o.Video)
	}
	
	if Reactions, ok := MessageeventMap["reactions"].([]interface{}); ok {
		ReactionsString, _ := json.Marshal(Reactions)
		json.Unmarshal(ReactionsString, &o.Reactions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messageevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
