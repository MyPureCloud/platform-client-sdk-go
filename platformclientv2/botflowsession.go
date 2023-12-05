package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Botflowsession
type Botflowsession struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the bot session.
	Id *string `json:"id,omitempty"`

	// Flow - The flow associated to this bot session.
	Flow *Entity `json:"flow,omitempty"`

	// Channel - Channel-specific information that describes the message channel/provider.
	Channel *Botchannel `json:"channel,omitempty"`

	// Language - The initial language of operation for the session.
	Language *string `json:"language,omitempty"`

	// EndLanguage - The language of the session at the time the session ended
	EndLanguage *string `json:"endLanguage,omitempty"`

	// BotResult - The reason for session termination.
	BotResult *string `json:"botResult,omitempty"`

	// BotResultCategory - The category of result for the session.
	BotResultCategory *string `json:"botResultCategory,omitempty"`

	// DateCreated - Timestamp indicating when the session was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// Conversation - The conversation details, across potentially multiple Bot Flow sessions.
	Conversation *Addressableentityref `json:"conversation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Botflowsession) SetField(field string, fieldValue interface{}) {
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

func (o Botflowsession) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias Botflowsession
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Flow *Entity `json:"flow,omitempty"`
		
		Channel *Botchannel `json:"channel,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		EndLanguage *string `json:"endLanguage,omitempty"`
		
		BotResult *string `json:"botResult,omitempty"`
		
		BotResultCategory *string `json:"botResultCategory,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Flow: o.Flow,
		
		Channel: o.Channel,
		
		Language: o.Language,
		
		EndLanguage: o.EndLanguage,
		
		BotResult: o.BotResult,
		
		BotResultCategory: o.BotResultCategory,
		
		DateCreated: DateCreated,
		
		Conversation: o.Conversation,
		Alias:    (Alias)(o),
	})
}

func (o *Botflowsession) UnmarshalJSON(b []byte) error {
	var BotflowsessionMap map[string]interface{}
	err := json.Unmarshal(b, &BotflowsessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BotflowsessionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Flow, ok := BotflowsessionMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if Channel, ok := BotflowsessionMap["channel"].(map[string]interface{}); ok {
		ChannelString, _ := json.Marshal(Channel)
		json.Unmarshal(ChannelString, &o.Channel)
	}
	
	if Language, ok := BotflowsessionMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if EndLanguage, ok := BotflowsessionMap["endLanguage"].(string); ok {
		o.EndLanguage = &EndLanguage
	}
    
	if BotResult, ok := BotflowsessionMap["botResult"].(string); ok {
		o.BotResult = &BotResult
	}
    
	if BotResultCategory, ok := BotflowsessionMap["botResultCategory"].(string); ok {
		o.BotResultCategory = &BotResultCategory
	}
    
	if dateCreatedString, ok := BotflowsessionMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if Conversation, ok := BotflowsessionMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Botflowsession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
