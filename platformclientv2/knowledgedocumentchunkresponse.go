package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentchunkresponse
type Knowledgedocumentchunkresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Query - Query to search chunks in the knowledge base.
	Query *string `json:"query,omitempty"`

	// Total - The total number of chunks matching the query.
	Total *int `json:"total,omitempty"`

	// PageCount - Number of pages returned in the result calculated according to the pageSize and the total
	PageCount *int `json:"pageCount,omitempty"`

	// PageSize - Page size of the returned results.
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber - Page number of the returned results.
	PageNumber *int `json:"pageNumber,omitempty"`

	// QueryType - The type of the query that initiates the chunks search.
	QueryType *string `json:"queryType,omitempty"`

	// SearchId - The globally unique identifier for the chunks search.
	SearchId *string `json:"searchId,omitempty"`

	// PreprocessQuery - Indicates whether the chunks search query should be preprocessed.
	PreprocessQuery *bool `json:"preprocessQuery,omitempty"`

	// ConfidenceThreshold - The confidence threshold for the chunk results. If applied, the returned results will have an equal or higher chunk confidence than the threshold.
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`

	// Results - Chunks matching the search query.
	Results *[]Documentchunkblock `json:"results,omitempty"`

	// Application - The client application details from which chunks search happened.
	Application *Knowledgesearchclientapplication `json:"application,omitempty"`

	// ConversationContext - Conversation context information if the chunks search is initiated in the context of a conversation.
	ConversationContext *Knowledgeconversationcontextresponse `json:"conversationContext,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentchunkresponse) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentchunkresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentchunkresponse
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		PreprocessQuery *bool `json:"preprocessQuery,omitempty"`
		
		ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`
		
		Results *[]Documentchunkblock `json:"results,omitempty"`
		
		Application *Knowledgesearchclientapplication `json:"application,omitempty"`
		
		ConversationContext *Knowledgeconversationcontextresponse `json:"conversationContext,omitempty"`
		Alias
	}{ 
		Query: o.Query,
		
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		QueryType: o.QueryType,
		
		SearchId: o.SearchId,
		
		PreprocessQuery: o.PreprocessQuery,
		
		ConfidenceThreshold: o.ConfidenceThreshold,
		
		Results: o.Results,
		
		Application: o.Application,
		
		ConversationContext: o.ConversationContext,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentchunkresponse) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentchunkresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentchunkresponseMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgedocumentchunkresponseMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if Total, ok := KnowledgedocumentchunkresponseMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PageCount, ok := KnowledgedocumentchunkresponseMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if PageSize, ok := KnowledgedocumentchunkresponseMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := KnowledgedocumentchunkresponseMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if QueryType, ok := KnowledgedocumentchunkresponseMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if SearchId, ok := KnowledgedocumentchunkresponseMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if PreprocessQuery, ok := KnowledgedocumentchunkresponseMap["preprocessQuery"].(bool); ok {
		o.PreprocessQuery = &PreprocessQuery
	}
    
	if ConfidenceThreshold, ok := KnowledgedocumentchunkresponseMap["confidenceThreshold"].(float64); ok {
		ConfidenceThresholdFloat32 := float32(ConfidenceThreshold)
		o.ConfidenceThreshold = &ConfidenceThresholdFloat32
	}
	
	if Results, ok := KnowledgedocumentchunkresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if Application, ok := KnowledgedocumentchunkresponseMap["application"].(map[string]interface{}); ok {
		ApplicationString, _ := json.Marshal(Application)
		json.Unmarshal(ApplicationString, &o.Application)
	}
	
	if ConversationContext, ok := KnowledgedocumentchunkresponseMap["conversationContext"].(map[string]interface{}); ok {
		ConversationContextString, _ := json.Marshal(ConversationContext)
		json.Unmarshal(ConversationContextString, &o.ConversationContext)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentchunkresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
