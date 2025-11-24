package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestionevent
type Conversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestionevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// SuggestionId
	SuggestionId *string `json:"suggestionId,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// TriggerType
	TriggerType *string `json:"triggerType,omitempty"`

	// EngagementType
	EngagementType *string `json:"engagementType,omitempty"`

	// Feedback
	Feedback *Conversationthirdpartyactionsuggestionstopicsuggestionfeedback `json:"feedback,omitempty"`

	// ActiveIntent
	ActiveIntent *Conversationthirdpartyactionsuggestionstopicsuggestedintent `json:"activeIntent,omitempty"`

	// Context
	Context *Conversationthirdpartyactionsuggestionstopicsuggestioncontext `json:"context,omitempty"`

	// ThirdPartySuggestion
	ThirdPartySuggestion *Conversationthirdpartyactionsuggestionstopicthirdpartysuggestion `json:"thirdPartySuggestion,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestionevent) SetField(field string, fieldValue interface{}) {
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

func (o Conversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestionevent) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EventTime", }
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
	type Alias Conversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestionevent
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		EventTime *string `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		SuggestionId *string `json:"suggestionId,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		TriggerType *string `json:"triggerType,omitempty"`
		
		EngagementType *string `json:"engagementType,omitempty"`
		
		Feedback *Conversationthirdpartyactionsuggestionstopicsuggestionfeedback `json:"feedback,omitempty"`
		
		ActiveIntent *Conversationthirdpartyactionsuggestionstopicsuggestedintent `json:"activeIntent,omitempty"`
		
		Context *Conversationthirdpartyactionsuggestionstopicsuggestioncontext `json:"context,omitempty"`
		
		ThirdPartySuggestion *Conversationthirdpartyactionsuggestionstopicthirdpartysuggestion `json:"thirdPartySuggestion,omitempty"`
		Alias
	}{ 
		EventTime: EventTime,
		
		ConversationId: o.ConversationId,
		
		SuggestionId: o.SuggestionId,
		
		State: o.State,
		
		TriggerType: o.TriggerType,
		
		EngagementType: o.EngagementType,
		
		Feedback: o.Feedback,
		
		ActiveIntent: o.ActiveIntent,
		
		Context: o.Context,
		
		ThirdPartySuggestion: o.ThirdPartySuggestion,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestionevent) UnmarshalJSON(b []byte) error {
	var ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap)
	if err != nil {
		return err
	}
	
	if eventTimeString, ok := ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if ConversationId, ok := ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if SuggestionId, ok := ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap["suggestionId"].(string); ok {
		o.SuggestionId = &SuggestionId
	}
    
	if State, ok := ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap["state"].(string); ok {
		o.State = &State
	}
    
	if TriggerType, ok := ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap["triggerType"].(string); ok {
		o.TriggerType = &TriggerType
	}
    
	if EngagementType, ok := ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap["engagementType"].(string); ok {
		o.EngagementType = &EngagementType
	}
    
	if Feedback, ok := ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap["feedback"].(map[string]interface{}); ok {
		FeedbackString, _ := json.Marshal(Feedback)
		json.Unmarshal(FeedbackString, &o.Feedback)
	}
	
	if ActiveIntent, ok := ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap["activeIntent"].(map[string]interface{}); ok {
		ActiveIntentString, _ := json.Marshal(ActiveIntent)
		json.Unmarshal(ActiveIntentString, &o.ActiveIntent)
	}
	
	if Context, ok := ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	
	if ThirdPartySuggestion, ok := ConversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestioneventMap["thirdPartySuggestion"].(map[string]interface{}); ok {
		ThirdPartySuggestionString, _ := json.Marshal(ThirdPartySuggestion)
		json.Unmarshal(ThirdPartySuggestionString, &o.ThirdPartySuggestion)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationthirdpartyactionsuggestionstopicthirdpartyactionsuggestionevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
