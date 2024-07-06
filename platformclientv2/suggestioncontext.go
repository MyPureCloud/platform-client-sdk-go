package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Suggestioncontext
type Suggestioncontext struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Queue - The queue used to assign the interaction to the user, if any.
	Queue *Addressableentityref `json:"queue,omitempty"`

	// MediaType - The media type of the conversation in which the suggestion event was raised.
	MediaType *string `json:"mediaType,omitempty"`

	// User - The agent participant who received the raised suggestion, if any.
	User *Userreference `json:"user,omitempty"`

	// ExternalContact - The external contact of the end-user participant, if any.
	ExternalContact *Addressableentityref `json:"externalContact,omitempty"`

	// Utterance - The utterance in the voice conversation, after which the suggestion was raised, if any.
	Utterance *Entity `json:"utterance,omitempty"`

	// Message - The message in the digital conversation, after which the suggestion was raised, if any.
	Message *Addressableentityref `json:"message,omitempty"`

	// QueryStatement - The query statement used when generating the suggestion, if any.
	QueryStatement *string `json:"queryStatement,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Suggestioncontext) SetField(field string, fieldValue interface{}) {
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

func (o Suggestioncontext) MarshalJSON() ([]byte, error) {
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
	type Alias Suggestioncontext
	
	return json.Marshal(&struct { 
		Queue *Addressableentityref `json:"queue,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		ExternalContact *Addressableentityref `json:"externalContact,omitempty"`
		
		Utterance *Entity `json:"utterance,omitempty"`
		
		Message *Addressableentityref `json:"message,omitempty"`
		
		QueryStatement *string `json:"queryStatement,omitempty"`
		Alias
	}{ 
		Queue: o.Queue,
		
		MediaType: o.MediaType,
		
		User: o.User,
		
		ExternalContact: o.ExternalContact,
		
		Utterance: o.Utterance,
		
		Message: o.Message,
		
		QueryStatement: o.QueryStatement,
		Alias:    (Alias)(o),
	})
}

func (o *Suggestioncontext) UnmarshalJSON(b []byte) error {
	var SuggestioncontextMap map[string]interface{}
	err := json.Unmarshal(b, &SuggestioncontextMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := SuggestioncontextMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if MediaType, ok := SuggestioncontextMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if User, ok := SuggestioncontextMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if ExternalContact, ok := SuggestioncontextMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if Utterance, ok := SuggestioncontextMap["utterance"].(map[string]interface{}); ok {
		UtteranceString, _ := json.Marshal(Utterance)
		json.Unmarshal(UtteranceString, &o.Utterance)
	}
	
	if Message, ok := SuggestioncontextMap["message"].(map[string]interface{}); ok {
		MessageString, _ := json.Marshal(Message)
		json.Unmarshal(MessageString, &o.Message)
	}
	
	if QueryStatement, ok := SuggestioncontextMap["queryStatement"].(string); ok {
		o.QueryStatement = &QueryStatement
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Suggestioncontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
