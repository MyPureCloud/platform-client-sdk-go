package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesourcessearchresponse
type Knowledgesourcessearchresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Query - Query to search content in the knowledge base.
	Query *string `json:"query,omitempty"`

	// SearchId - The globally unique identifier for the search.
	SearchId *string `json:"searchId,omitempty"`

	// SessionId - The sessionId for search request.
	SessionId *string `json:"sessionId,omitempty"`

	// Result - Content matching the search query.
	Result *Knowledgesearchresult `json:"result,omitempty"`

	// KnowledgeSettingId - Knowledge Setting Id used for the search request.
	KnowledgeSettingId *string `json:"knowledgeSettingId,omitempty"`

	// ConversationContext - Conversation context information if the search is initiated in the context of a conversation.
	ConversationContext *Knowledgev3conversationcontextresponse `json:"conversationContext,omitempty"`

	// Application - The client application details from which search happened.
	Application *V3knowledgesearchclientapplication `json:"application,omitempty"`

	// QueryType - The type of the query that initiates the search.
	QueryType *string `json:"queryType,omitempty"`

	// GenerationLanguage - The language used for answer generation.
	GenerationLanguage *string `json:"generationLanguage,omitempty"`

	// AnswerGeneration - Indicates if answer generation was enabled for the setting.
	AnswerGeneration *bool `json:"answerGeneration,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgesourcessearchresponse) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgesourcessearchresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgesourcessearchresponse
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		Result *Knowledgesearchresult `json:"result,omitempty"`
		
		KnowledgeSettingId *string `json:"knowledgeSettingId,omitempty"`
		
		ConversationContext *Knowledgev3conversationcontextresponse `json:"conversationContext,omitempty"`
		
		Application *V3knowledgesearchclientapplication `json:"application,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		GenerationLanguage *string `json:"generationLanguage,omitempty"`
		
		AnswerGeneration *bool `json:"answerGeneration,omitempty"`
		Alias
	}{ 
		Query: o.Query,
		
		SearchId: o.SearchId,
		
		SessionId: o.SessionId,
		
		Result: o.Result,
		
		KnowledgeSettingId: o.KnowledgeSettingId,
		
		ConversationContext: o.ConversationContext,
		
		Application: o.Application,
		
		QueryType: o.QueryType,
		
		GenerationLanguage: o.GenerationLanguage,
		
		AnswerGeneration: o.AnswerGeneration,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgesourcessearchresponse) UnmarshalJSON(b []byte) error {
	var KnowledgesourcessearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesourcessearchresponseMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgesourcessearchresponseMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if SearchId, ok := KnowledgesourcessearchresponseMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if SessionId, ok := KnowledgesourcessearchresponseMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if Result, ok := KnowledgesourcessearchresponseMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if KnowledgeSettingId, ok := KnowledgesourcessearchresponseMap["knowledgeSettingId"].(string); ok {
		o.KnowledgeSettingId = &KnowledgeSettingId
	}
    
	if ConversationContext, ok := KnowledgesourcessearchresponseMap["conversationContext"].(map[string]interface{}); ok {
		ConversationContextString, _ := json.Marshal(ConversationContext)
		json.Unmarshal(ConversationContextString, &o.ConversationContext)
	}
	
	if Application, ok := KnowledgesourcessearchresponseMap["application"].(map[string]interface{}); ok {
		ApplicationString, _ := json.Marshal(Application)
		json.Unmarshal(ApplicationString, &o.Application)
	}
	
	if QueryType, ok := KnowledgesourcessearchresponseMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if GenerationLanguage, ok := KnowledgesourcessearchresponseMap["generationLanguage"].(string); ok {
		o.GenerationLanguage = &GenerationLanguage
	}
    
	if AnswerGeneration, ok := KnowledgesourcessearchresponseMap["answerGeneration"].(bool); ok {
		o.AnswerGeneration = &AnswerGeneration
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesourcessearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
