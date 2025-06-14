package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentview
type Knowledgedocumentview struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DocumentVariationId - The variation of the viewed document.
	DocumentVariationId *string `json:"documentVariationId,omitempty"`

	// DocumentVersionId - The version of the viewed document.
	DocumentVersionId *string `json:"documentVersionId,omitempty"`

	// SearchId - The search that surfaced the viewed document.
	SearchId *string `json:"searchId,omitempty"`

	// QueryType - The type of the query that surfaced the document.
	QueryType *string `json:"queryType,omitempty"`

	// SurfacingMethod - The method how knowledge was surfaced. Article: Full article was shown. Snippet: A snippet from the article was shown. Highlight: A highlighted answer in a snippet was shown.Generative: A generated answer in a snippet was shown.
	SurfacingMethod *string `json:"surfacingMethod,omitempty"`

	// Application - The client application from which the document was viewed.
	Application *Knowledgesearchclientapplication `json:"application,omitempty"`

	// SessionId - The unique identifier of the knowledge session in which the document was viewed.
	SessionId *string `json:"sessionId,omitempty"`

	// ConversationContext - Conversation context information if the document was viewed in the context of a conversation.
	ConversationContext *Knowledgeconversationcontext `json:"conversationContext,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentview) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentview) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentview
	
	return json.Marshal(&struct { 
		DocumentVariationId *string `json:"documentVariationId,omitempty"`
		
		DocumentVersionId *string `json:"documentVersionId,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		SurfacingMethod *string `json:"surfacingMethod,omitempty"`
		
		Application *Knowledgesearchclientapplication `json:"application,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		ConversationContext *Knowledgeconversationcontext `json:"conversationContext,omitempty"`
		Alias
	}{ 
		DocumentVariationId: o.DocumentVariationId,
		
		DocumentVersionId: o.DocumentVersionId,
		
		SearchId: o.SearchId,
		
		QueryType: o.QueryType,
		
		SurfacingMethod: o.SurfacingMethod,
		
		Application: o.Application,
		
		SessionId: o.SessionId,
		
		ConversationContext: o.ConversationContext,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentview) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentviewMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentviewMap)
	if err != nil {
		return err
	}
	
	if DocumentVariationId, ok := KnowledgedocumentviewMap["documentVariationId"].(string); ok {
		o.DocumentVariationId = &DocumentVariationId
	}
    
	if DocumentVersionId, ok := KnowledgedocumentviewMap["documentVersionId"].(string); ok {
		o.DocumentVersionId = &DocumentVersionId
	}
    
	if SearchId, ok := KnowledgedocumentviewMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if QueryType, ok := KnowledgedocumentviewMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if SurfacingMethod, ok := KnowledgedocumentviewMap["surfacingMethod"].(string); ok {
		o.SurfacingMethod = &SurfacingMethod
	}
    
	if Application, ok := KnowledgedocumentviewMap["application"].(map[string]interface{}); ok {
		ApplicationString, _ := json.Marshal(Application)
		json.Unmarshal(ApplicationString, &o.Application)
	}
	
	if SessionId, ok := KnowledgedocumentviewMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if ConversationContext, ok := KnowledgedocumentviewMap["conversationContext"].(map[string]interface{}); ok {
		ConversationContextString, _ := json.Marshal(ConversationContext)
		json.Unmarshal(ConversationContextString, &o.ConversationContext)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
