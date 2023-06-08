package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Touchpointresponse
type Touchpointresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConversationId - ID of conversation.
	ConversationId *string `json:"conversationId,omitempty"`

	// AgentId - ID of agent.
	AgentId *string `json:"agentId,omitempty"`

	// AssociatedValue - The value attributed to this touchpoint.
	AssociatedValue *float32 `json:"associatedValue,omitempty"`

	// MediaType - Media Type of the touchpoint; allowed values are Email, Message and Voice.
	MediaType *string `json:"mediaType,omitempty"`

	// State - Outcome Attribution Touchpoint status.
	State *string `json:"state,omitempty"`

	// Message - Additional information on the state of the touchpoint entity, populated when the touchpoint has an error.
	Message *string `json:"message,omitempty"`

	// CreatedDate - Date conversation happened. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Touchpointresponse) SetField(field string, fieldValue interface{}) {
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

func (o Touchpointresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate", }
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
	type Alias Touchpointresponse
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		AgentId *string `json:"agentId,omitempty"`
		
		AssociatedValue *float32 `json:"associatedValue,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		Alias
	}{ 
		ConversationId: o.ConversationId,
		
		AgentId: o.AgentId,
		
		AssociatedValue: o.AssociatedValue,
		
		MediaType: o.MediaType,
		
		State: o.State,
		
		Message: o.Message,
		
		CreatedDate: CreatedDate,
		Alias:    (Alias)(o),
	})
}

func (o *Touchpointresponse) UnmarshalJSON(b []byte) error {
	var TouchpointresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TouchpointresponseMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := TouchpointresponseMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if AgentId, ok := TouchpointresponseMap["agentId"].(string); ok {
		o.AgentId = &AgentId
	}
    
	if AssociatedValue, ok := TouchpointresponseMap["associatedValue"].(float64); ok {
		AssociatedValueFloat32 := float32(AssociatedValue)
		o.AssociatedValue = &AssociatedValueFloat32
	}
    
	if MediaType, ok := TouchpointresponseMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if State, ok := TouchpointresponseMap["state"].(string); ok {
		o.State = &State
	}
    
	if Message, ok := TouchpointresponseMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if createdDateString, ok := TouchpointresponseMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Touchpointresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
