package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Suggestion
type Suggestion struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Conversation - The conversation that the suggestions correspond to.
	Conversation *Addressableentityref `json:"conversation,omitempty"`

	// Assistant - The assistant that was used to provide the suggestions.
	Assistant *Addressableentityref `json:"assistant,omitempty"`

	// VarType - The type of the documents for which the suggestion is.
	VarType *string `json:"type,omitempty"`

	// Faq - The Faq from the knowledgebase that was provided as the suggestion.
	Faq *Faq `json:"faq,omitempty"`

	// Article - The article from the knowledgebase that was provided as the suggestion.
	Article *Article `json:"article,omitempty"`

	// DateCreated - Date when the suggestion was created. For example: yyyy-MM-ddTHH:mm:ss.SSZ. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// AnswerRecordId - The ID of the knowledge search that provided the suggestion.
	AnswerRecordId *string `json:"answerRecordId,omitempty"`

	// TriggerType - The trigger type of the suggestion.
	TriggerType *string `json:"triggerType,omitempty"`

	// Context - The conversation context in which the suggestion was raised.
	Context *Suggestioncontext `json:"context,omitempty"`

	// State - The state of the suggestion.
	State *string `json:"state,omitempty"`

	// KnowledgeSearch - The suggested knowledge search result that was provided as the suggestion.
	KnowledgeSearch *Suggestionknowledgesearch `json:"knowledgeSearch,omitempty"`

	// KnowledgeArticle - The suggested knowledge article that was provided as the suggestion.
	KnowledgeArticle *Suggestionknowledgearticle `json:"knowledgeArticle,omitempty"`

	// CannedResponse - The suggested canned response that was provided as the suggestion.
	CannedResponse *Suggestioncannedresponse `json:"cannedResponse,omitempty"`

	// Script - The suggested script that was provided as the suggestion.
	Script *Suggestionscript `json:"script,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Suggestion) SetField(field string, fieldValue interface{}) {
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

func (o Suggestion) MarshalJSON() ([]byte, error) {
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
	type Alias Suggestion
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		
		Assistant *Addressableentityref `json:"assistant,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Faq *Faq `json:"faq,omitempty"`
		
		Article *Article `json:"article,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		AnswerRecordId *string `json:"answerRecordId,omitempty"`
		
		TriggerType *string `json:"triggerType,omitempty"`
		
		Context *Suggestioncontext `json:"context,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		KnowledgeSearch *Suggestionknowledgesearch `json:"knowledgeSearch,omitempty"`
		
		KnowledgeArticle *Suggestionknowledgearticle `json:"knowledgeArticle,omitempty"`
		
		CannedResponse *Suggestioncannedresponse `json:"cannedResponse,omitempty"`
		
		Script *Suggestionscript `json:"script,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Conversation: o.Conversation,
		
		Assistant: o.Assistant,
		
		VarType: o.VarType,
		
		Faq: o.Faq,
		
		Article: o.Article,
		
		DateCreated: DateCreated,
		
		AnswerRecordId: o.AnswerRecordId,
		
		TriggerType: o.TriggerType,
		
		Context: o.Context,
		
		State: o.State,
		
		KnowledgeSearch: o.KnowledgeSearch,
		
		KnowledgeArticle: o.KnowledgeArticle,
		
		CannedResponse: o.CannedResponse,
		
		Script: o.Script,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Suggestion) UnmarshalJSON(b []byte) error {
	var SuggestionMap map[string]interface{}
	err := json.Unmarshal(b, &SuggestionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SuggestionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Conversation, ok := SuggestionMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Assistant, ok := SuggestionMap["assistant"].(map[string]interface{}); ok {
		AssistantString, _ := json.Marshal(Assistant)
		json.Unmarshal(AssistantString, &o.Assistant)
	}
	
	if VarType, ok := SuggestionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Faq, ok := SuggestionMap["faq"].(map[string]interface{}); ok {
		FaqString, _ := json.Marshal(Faq)
		json.Unmarshal(FaqString, &o.Faq)
	}
	
	if Article, ok := SuggestionMap["article"].(map[string]interface{}); ok {
		ArticleString, _ := json.Marshal(Article)
		json.Unmarshal(ArticleString, &o.Article)
	}
	
	if dateCreatedString, ok := SuggestionMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if AnswerRecordId, ok := SuggestionMap["answerRecordId"].(string); ok {
		o.AnswerRecordId = &AnswerRecordId
	}
    
	if TriggerType, ok := SuggestionMap["triggerType"].(string); ok {
		o.TriggerType = &TriggerType
	}
    
	if Context, ok := SuggestionMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	
	if State, ok := SuggestionMap["state"].(string); ok {
		o.State = &State
	}
    
	if KnowledgeSearch, ok := SuggestionMap["knowledgeSearch"].(map[string]interface{}); ok {
		KnowledgeSearchString, _ := json.Marshal(KnowledgeSearch)
		json.Unmarshal(KnowledgeSearchString, &o.KnowledgeSearch)
	}
	
	if KnowledgeArticle, ok := SuggestionMap["knowledgeArticle"].(map[string]interface{}); ok {
		KnowledgeArticleString, _ := json.Marshal(KnowledgeArticle)
		json.Unmarshal(KnowledgeArticleString, &o.KnowledgeArticle)
	}
	
	if CannedResponse, ok := SuggestionMap["cannedResponse"].(map[string]interface{}); ok {
		CannedResponseString, _ := json.Marshal(CannedResponse)
		json.Unmarshal(CannedResponseString, &o.CannedResponse)
	}
	
	if Script, ok := SuggestionMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if SelfUri, ok := SuggestionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Suggestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
