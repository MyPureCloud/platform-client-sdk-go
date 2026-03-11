package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesourcessearchrequest
type Knowledgesourcessearchrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Query - Input query to search content on the knowledge setting.
	Query *string `json:"query,omitempty"`

	// KnowledgeSettingId - Knowledge Setting Id to use for search request.
	KnowledgeSettingId *string `json:"knowledgeSettingId,omitempty"`

	// Application - The client application details from which search requested.
	Application *V3knowledgesearchclientapplication `json:"application,omitempty"`

	// ConversationContext - Conversation context information if the search is initiated in the context of a conversation.
	ConversationContext *Knowledgev3conversationcontext `json:"conversationContext,omitempty"`

	// SessionId - The session id for search request.
	SessionId *string `json:"sessionId,omitempty"`

	// QueryType - The type of the query that initiates the search.
	QueryType *string `json:"queryType,omitempty"`

	// GenerationLanguage - The language to use for answer generation.
	GenerationLanguage *string `json:"generationLanguage,omitempty"`

	// ConversationTurns - List of conversation turns to use for stateful search.
	ConversationTurns *[]Knowledgeconversationturn `json:"conversationTurns,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgesourcessearchrequest) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgesourcessearchrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgesourcessearchrequest
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		KnowledgeSettingId *string `json:"knowledgeSettingId,omitempty"`
		
		Application *V3knowledgesearchclientapplication `json:"application,omitempty"`
		
		ConversationContext *Knowledgev3conversationcontext `json:"conversationContext,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		GenerationLanguage *string `json:"generationLanguage,omitempty"`
		
		ConversationTurns *[]Knowledgeconversationturn `json:"conversationTurns,omitempty"`
		Alias
	}{ 
		Query: o.Query,
		
		KnowledgeSettingId: o.KnowledgeSettingId,
		
		Application: o.Application,
		
		ConversationContext: o.ConversationContext,
		
		SessionId: o.SessionId,
		
		QueryType: o.QueryType,
		
		GenerationLanguage: o.GenerationLanguage,
		
		ConversationTurns: o.ConversationTurns,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgesourcessearchrequest) UnmarshalJSON(b []byte) error {
	var KnowledgesourcessearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesourcessearchrequestMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgesourcessearchrequestMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if KnowledgeSettingId, ok := KnowledgesourcessearchrequestMap["knowledgeSettingId"].(string); ok {
		o.KnowledgeSettingId = &KnowledgeSettingId
	}
    
	if Application, ok := KnowledgesourcessearchrequestMap["application"].(map[string]interface{}); ok {
		ApplicationString, _ := json.Marshal(Application)
		json.Unmarshal(ApplicationString, &o.Application)
	}
	
	if ConversationContext, ok := KnowledgesourcessearchrequestMap["conversationContext"].(map[string]interface{}); ok {
		ConversationContextString, _ := json.Marshal(ConversationContext)
		json.Unmarshal(ConversationContextString, &o.ConversationContext)
	}
	
	if SessionId, ok := KnowledgesourcessearchrequestMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if QueryType, ok := KnowledgesourcessearchrequestMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if GenerationLanguage, ok := KnowledgesourcessearchrequestMap["generationLanguage"].(string); ok {
		o.GenerationLanguage = &GenerationLanguage
	}
    
	if ConversationTurns, ok := KnowledgesourcessearchrequestMap["conversationTurns"].([]interface{}); ok {
		ConversationTurnsString, _ := json.Marshal(ConversationTurns)
		json.Unmarshal(ConversationTurnsString, &o.ConversationTurns)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesourcessearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
