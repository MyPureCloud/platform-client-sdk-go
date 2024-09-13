package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentsanswerfilter
type Knowledgedocumentsanswerfilter struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Query - The search query.
	Query *string `json:"query,omitempty"`

	// Language - The language of the documents.
	Language *string `json:"language,omitempty"`

	// AppType - The appType
	AppType *string `json:"appType,omitempty"`

	// QueryType - The query type
	QueryType *string `json:"queryType,omitempty"`

	// SearchId - The search id.
	SearchId *string `json:"searchId,omitempty"`

	// InsertHighlightIntoVariationContent - If specified - insert highlight data into the variation content.
	InsertHighlightIntoVariationContent *bool `json:"insertHighlightIntoVariationContent,omitempty"`

	// AnswerMode - Allows extracted answers from an article (AnswerHighlight) and/or AI-generated answers (AnswerGeneration). Default mode: AnswerHighlight
	AnswerMode *[]string `json:"answerMode,omitempty"`

	// VariationIds - The variation Ids to answer.
	VariationIds *[]string `json:"variationIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentsanswerfilter) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentsanswerfilter) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentsanswerfilter
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		AppType *string `json:"appType,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		InsertHighlightIntoVariationContent *bool `json:"insertHighlightIntoVariationContent,omitempty"`
		
		AnswerMode *[]string `json:"answerMode,omitempty"`
		
		VariationIds *[]string `json:"variationIds,omitempty"`
		Alias
	}{ 
		Query: o.Query,
		
		Language: o.Language,
		
		AppType: o.AppType,
		
		QueryType: o.QueryType,
		
		SearchId: o.SearchId,
		
		InsertHighlightIntoVariationContent: o.InsertHighlightIntoVariationContent,
		
		AnswerMode: o.AnswerMode,
		
		VariationIds: o.VariationIds,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentsanswerfilter) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentsanswerfilterMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentsanswerfilterMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgedocumentsanswerfilterMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if Language, ok := KnowledgedocumentsanswerfilterMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if AppType, ok := KnowledgedocumentsanswerfilterMap["appType"].(string); ok {
		o.AppType = &AppType
	}
    
	if QueryType, ok := KnowledgedocumentsanswerfilterMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if SearchId, ok := KnowledgedocumentsanswerfilterMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if InsertHighlightIntoVariationContent, ok := KnowledgedocumentsanswerfilterMap["insertHighlightIntoVariationContent"].(bool); ok {
		o.InsertHighlightIntoVariationContent = &InsertHighlightIntoVariationContent
	}
    
	if AnswerMode, ok := KnowledgedocumentsanswerfilterMap["answerMode"].([]interface{}); ok {
		AnswerModeString, _ := json.Marshal(AnswerMode)
		json.Unmarshal(AnswerModeString, &o.AnswerMode)
	}
	
	if VariationIds, ok := KnowledgedocumentsanswerfilterMap["variationIds"].([]interface{}); ok {
		VariationIdsString, _ := json.Marshal(VariationIds)
		json.Unmarshal(VariationIdsString, &o.VariationIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsanswerfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
