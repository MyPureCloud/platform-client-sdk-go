package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationproperties
type Conversationproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// IsWaiting - Indicates filtering for waiting
	IsWaiting *bool `json:"isWaiting,omitempty"`

	// IsActive - Indicates filtering for active
	IsActive *bool `json:"isActive,omitempty"`

	// IsAcd - Indicates filtering for Acd
	IsAcd *bool `json:"isAcd,omitempty"`

	// IsPreferred - Indicates filtering for Preferred Agent Routing
	IsPreferred *bool `json:"isPreferred,omitempty"`

	// IsScreenshare - Indicates filtering for screenshare
	IsScreenshare *bool `json:"isScreenshare,omitempty"`

	// IsCobrowse - Indicates filtering for Cobrowse
	IsCobrowse *bool `json:"isCobrowse,omitempty"`

	// IsVoicemail - Indicates filtering for Voice mail
	IsVoicemail *bool `json:"isVoicemail,omitempty"`

	// IsFlagged - Indicates filtering for flagged
	IsFlagged *bool `json:"isFlagged,omitempty"`

	// IsMonitored - Indicates filtering for monitored
	IsMonitored *bool `json:"isMonitored,omitempty"`

	// FilterWrapUpNotes - Indicates filtering for WrapUpNotes
	FilterWrapUpNotes *bool `json:"filterWrapUpNotes,omitempty"`

	// MatchAll - Indicates comparison operation, TRUE indicates filters will use AND logic, FALSE indicates OR logic
	MatchAll *bool `json:"matchAll,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationproperties) SetField(field string, fieldValue interface{}) {
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

func (o Conversationproperties) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationproperties
	
	return json.Marshal(&struct { 
		IsWaiting *bool `json:"isWaiting,omitempty"`
		
		IsActive *bool `json:"isActive,omitempty"`
		
		IsAcd *bool `json:"isAcd,omitempty"`
		
		IsPreferred *bool `json:"isPreferred,omitempty"`
		
		IsScreenshare *bool `json:"isScreenshare,omitempty"`
		
		IsCobrowse *bool `json:"isCobrowse,omitempty"`
		
		IsVoicemail *bool `json:"isVoicemail,omitempty"`
		
		IsFlagged *bool `json:"isFlagged,omitempty"`
		
		IsMonitored *bool `json:"isMonitored,omitempty"`
		
		FilterWrapUpNotes *bool `json:"filterWrapUpNotes,omitempty"`
		
		MatchAll *bool `json:"matchAll,omitempty"`
		Alias
	}{ 
		IsWaiting: o.IsWaiting,
		
		IsActive: o.IsActive,
		
		IsAcd: o.IsAcd,
		
		IsPreferred: o.IsPreferred,
		
		IsScreenshare: o.IsScreenshare,
		
		IsCobrowse: o.IsCobrowse,
		
		IsVoicemail: o.IsVoicemail,
		
		IsFlagged: o.IsFlagged,
		
		IsMonitored: o.IsMonitored,
		
		FilterWrapUpNotes: o.FilterWrapUpNotes,
		
		MatchAll: o.MatchAll,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationproperties) UnmarshalJSON(b []byte) error {
	var ConversationpropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationpropertiesMap)
	if err != nil {
		return err
	}
	
	if IsWaiting, ok := ConversationpropertiesMap["isWaiting"].(bool); ok {
		o.IsWaiting = &IsWaiting
	}
    
	if IsActive, ok := ConversationpropertiesMap["isActive"].(bool); ok {
		o.IsActive = &IsActive
	}
    
	if IsAcd, ok := ConversationpropertiesMap["isAcd"].(bool); ok {
		o.IsAcd = &IsAcd
	}
    
	if IsPreferred, ok := ConversationpropertiesMap["isPreferred"].(bool); ok {
		o.IsPreferred = &IsPreferred
	}
    
	if IsScreenshare, ok := ConversationpropertiesMap["isScreenshare"].(bool); ok {
		o.IsScreenshare = &IsScreenshare
	}
    
	if IsCobrowse, ok := ConversationpropertiesMap["isCobrowse"].(bool); ok {
		o.IsCobrowse = &IsCobrowse
	}
    
	if IsVoicemail, ok := ConversationpropertiesMap["isVoicemail"].(bool); ok {
		o.IsVoicemail = &IsVoicemail
	}
    
	if IsFlagged, ok := ConversationpropertiesMap["isFlagged"].(bool); ok {
		o.IsFlagged = &IsFlagged
	}
    
	if IsMonitored, ok := ConversationpropertiesMap["isMonitored"].(bool); ok {
		o.IsMonitored = &IsMonitored
	}
    
	if FilterWrapUpNotes, ok := ConversationpropertiesMap["filterWrapUpNotes"].(bool); ok {
		o.FilterWrapUpNotes = &FilterWrapUpNotes
	}
    
	if MatchAll, ok := ConversationpropertiesMap["matchAll"].(bool); ok {
		o.MatchAll = &MatchAll
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
