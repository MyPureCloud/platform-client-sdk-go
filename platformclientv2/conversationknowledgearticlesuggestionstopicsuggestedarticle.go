package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationknowledgearticlesuggestionstopicsuggestedarticle
type Conversationknowledgearticlesuggestionstopicsuggestedarticle struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Title
	Title *string `json:"title,omitempty"`

	// Uri
	Uri *string `json:"uri,omitempty"`

	// Snippets
	Snippets *[]string `json:"snippets,omitempty"`

	// Metadata
	Metadata *map[string]string `json:"metadata,omitempty"`

	// DocumentId
	DocumentId *string `json:"documentId,omitempty"`

	// VersionId
	VersionId *string `json:"versionId,omitempty"`

	// VariationIds
	VariationIds *[]string `json:"variationIds,omitempty"`

	// KnowledgeAnswer
	KnowledgeAnswer *Conversationknowledgearticlesuggestionstopicknowledgeanswer `json:"knowledgeAnswer,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationknowledgearticlesuggestionstopicsuggestedarticle) SetField(field string, fieldValue interface{}) {
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

func (o Conversationknowledgearticlesuggestionstopicsuggestedarticle) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationknowledgearticlesuggestionstopicsuggestedarticle
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Uri *string `json:"uri,omitempty"`
		
		Snippets *[]string `json:"snippets,omitempty"`
		
		Metadata *map[string]string `json:"metadata,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		
		VersionId *string `json:"versionId,omitempty"`
		
		VariationIds *[]string `json:"variationIds,omitempty"`
		
		KnowledgeAnswer *Conversationknowledgearticlesuggestionstopicknowledgeanswer `json:"knowledgeAnswer,omitempty"`
		Alias
	}{ 
		Title: o.Title,
		
		Uri: o.Uri,
		
		Snippets: o.Snippets,
		
		Metadata: o.Metadata,
		
		DocumentId: o.DocumentId,
		
		VersionId: o.VersionId,
		
		VariationIds: o.VariationIds,
		
		KnowledgeAnswer: o.KnowledgeAnswer,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationknowledgearticlesuggestionstopicsuggestedarticle) UnmarshalJSON(b []byte) error {
	var ConversationknowledgearticlesuggestionstopicsuggestedarticleMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationknowledgearticlesuggestionstopicsuggestedarticleMap)
	if err != nil {
		return err
	}
	
	if Title, ok := ConversationknowledgearticlesuggestionstopicsuggestedarticleMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Uri, ok := ConversationknowledgearticlesuggestionstopicsuggestedarticleMap["uri"].(string); ok {
		o.Uri = &Uri
	}
    
	if Snippets, ok := ConversationknowledgearticlesuggestionstopicsuggestedarticleMap["snippets"].([]interface{}); ok {
		SnippetsString, _ := json.Marshal(Snippets)
		json.Unmarshal(SnippetsString, &o.Snippets)
	}
	
	if Metadata, ok := ConversationknowledgearticlesuggestionstopicsuggestedarticleMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if DocumentId, ok := ConversationknowledgearticlesuggestionstopicsuggestedarticleMap["documentId"].(string); ok {
		o.DocumentId = &DocumentId
	}
    
	if VersionId, ok := ConversationknowledgearticlesuggestionstopicsuggestedarticleMap["versionId"].(string); ok {
		o.VersionId = &VersionId
	}
    
	if VariationIds, ok := ConversationknowledgearticlesuggestionstopicsuggestedarticleMap["variationIds"].([]interface{}); ok {
		VariationIdsString, _ := json.Marshal(VariationIds)
		json.Unmarshal(VariationIdsString, &o.VariationIds)
	}
	
	if KnowledgeAnswer, ok := ConversationknowledgearticlesuggestionstopicsuggestedarticleMap["knowledgeAnswer"].(map[string]interface{}); ok {
		KnowledgeAnswerString, _ := json.Marshal(KnowledgeAnswer)
		json.Unmarshal(KnowledgeAnswerString, &o.KnowledgeAnswer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationknowledgearticlesuggestionstopicsuggestedarticle) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
