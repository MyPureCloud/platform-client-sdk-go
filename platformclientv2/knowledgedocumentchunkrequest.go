package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentchunkrequest
type Knowledgedocumentchunkrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Query - Query to search chunks in the knowledge base.
	Query *string `json:"query,omitempty"`

	// PageSize - Page size of the returned results.
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber - Page number of the returned results.
	PageNumber *int `json:"pageNumber,omitempty"`

	// Filter - Filter for the document chunks.
	Filter *Documentquery `json:"filter,omitempty"`

	// QueryType - The type of the query that initiates the chunks search.
	QueryType *string `json:"queryType,omitempty"`

	// PreprocessQuery - Indicates whether the chunks search query should be preprocessed.
	PreprocessQuery *bool `json:"preprocessQuery,omitempty"`

	// IncludeDraftDocuments - Indicates whether the chunk results would also include draft documents.
	IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`

	// Application - The client application details from which chunks request was sent.
	Application *Knowledgesearchclientapplication `json:"application,omitempty"`

	// ConversationContext - Conversation context information if the chunks search is initiated in the context of a conversation.
	ConversationContext *Knowledgeconversationcontext `json:"conversationContext,omitempty"`

	// ConfidenceThreshold - The confidence threshold for the chunk results. If applied, the returned results will have an equal or higher confidence than the threshold. The value should be between 0 to 1.
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentchunkrequest) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentchunkrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentchunkrequest
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Filter *Documentquery `json:"filter,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		PreprocessQuery *bool `json:"preprocessQuery,omitempty"`
		
		IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`
		
		Application *Knowledgesearchclientapplication `json:"application,omitempty"`
		
		ConversationContext *Knowledgeconversationcontext `json:"conversationContext,omitempty"`
		
		ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`
		Alias
	}{ 
		Query: o.Query,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Filter: o.Filter,
		
		QueryType: o.QueryType,
		
		PreprocessQuery: o.PreprocessQuery,
		
		IncludeDraftDocuments: o.IncludeDraftDocuments,
		
		Application: o.Application,
		
		ConversationContext: o.ConversationContext,
		
		ConfidenceThreshold: o.ConfidenceThreshold,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentchunkrequest) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentchunkrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentchunkrequestMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgedocumentchunkrequestMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if PageSize, ok := KnowledgedocumentchunkrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := KnowledgedocumentchunkrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Filter, ok := KnowledgedocumentchunkrequestMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if QueryType, ok := KnowledgedocumentchunkrequestMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if PreprocessQuery, ok := KnowledgedocumentchunkrequestMap["preprocessQuery"].(bool); ok {
		o.PreprocessQuery = &PreprocessQuery
	}
    
	if IncludeDraftDocuments, ok := KnowledgedocumentchunkrequestMap["includeDraftDocuments"].(bool); ok {
		o.IncludeDraftDocuments = &IncludeDraftDocuments
	}
    
	if Application, ok := KnowledgedocumentchunkrequestMap["application"].(map[string]interface{}); ok {
		ApplicationString, _ := json.Marshal(Application)
		json.Unmarshal(ApplicationString, &o.Application)
	}
	
	if ConversationContext, ok := KnowledgedocumentchunkrequestMap["conversationContext"].(map[string]interface{}); ok {
		ConversationContextString, _ := json.Marshal(ConversationContext)
		json.Unmarshal(ConversationContextString, &o.ConversationContext)
	}
	
	if ConfidenceThreshold, ok := KnowledgedocumentchunkrequestMap["confidenceThreshold"].(float64); ok {
		ConfidenceThresholdFloat32 := float32(ConfidenceThreshold)
		o.ConfidenceThreshold = &ConfidenceThresholdFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentchunkrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
