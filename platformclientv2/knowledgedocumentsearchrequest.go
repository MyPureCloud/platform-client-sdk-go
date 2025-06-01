package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentsearchrequest
type Knowledgedocumentsearchrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Query - Query to search content in the knowledge base. Maximum of 30 records per query can be fetched.
	Query *string `json:"query,omitempty"`

	// PageSize - Page size of the returned results.
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber - Page number of the returned results.
	PageNumber *int `json:"pageNumber,omitempty"`

	// SearchId - The globally unique identifier for the search.
	SearchId *string `json:"searchId,omitempty"`

	// Total - The total number of documents matching the query.
	Total *int `json:"total,omitempty"`

	// PageCount - Number of pages returned in the result calculated according to the pageSize and the total
	PageCount *int `json:"pageCount,omitempty"`

	// QueryType - The type of the query that initiates the search.
	QueryType *string `json:"queryType,omitempty"`

	// IncludeDraftDocuments - Indicates whether the search results would also include draft documents.
	IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`

	// Interval - Retrieves the documents created/modified/published in specified date and time range.
	Interval *Documentqueryinterval `json:"interval,omitempty"`

	// Filter - Filter for the document search.
	Filter *Documentquery `json:"filter,omitempty"`

	// SortOrder - The sort order for search results.
	SortOrder *string `json:"sortOrder,omitempty"`

	// SortBy - The field in the documents that you want to sort the search results by.
	SortBy *string `json:"sortBy,omitempty"`

	// Application - The client application details from which search request was sent.
	Application *Knowledgesearchclientapplication `json:"application,omitempty"`

	// ConversationContext - Conversation context information if the search is initiated in the context of a conversation.
	ConversationContext *Knowledgeconversationcontext `json:"conversationContext,omitempty"`

	// ConfidenceThreshold - The confidence threshold for the search results. If applied, the returned results will have an equal or higher confidence than the threshold. The value should be between 0 to 1.
	ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`

	// AnswerHighlightTopResults - The number of articles to be sent for answer-highlighting. Can range from 1-5.
	AnswerHighlightTopResults *int `json:"answerHighlightTopResults,omitempty"`

	// AnswerMode - Allows extracted answers from an article (AnswerHighlight) and/or AI-generated answers (AnswerGeneration). Default mode: AnswerHighlight. Use this property with answerHighlightTopResults.
	AnswerMode *[]string `json:"answerMode,omitempty"`

	// PreprocessQuery - Indicates whether the search query should be preprocessed.
	PreprocessQuery *bool `json:"preprocessQuery,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentsearchrequest) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentsearchrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentsearchrequest
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`
		
		Interval *Documentqueryinterval `json:"interval,omitempty"`
		
		Filter *Documentquery `json:"filter,omitempty"`
		
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		Application *Knowledgesearchclientapplication `json:"application,omitempty"`
		
		ConversationContext *Knowledgeconversationcontext `json:"conversationContext,omitempty"`
		
		ConfidenceThreshold *float32 `json:"confidenceThreshold,omitempty"`
		
		AnswerHighlightTopResults *int `json:"answerHighlightTopResults,omitempty"`
		
		AnswerMode *[]string `json:"answerMode,omitempty"`
		
		PreprocessQuery *bool `json:"preprocessQuery,omitempty"`
		Alias
	}{ 
		Query: o.Query,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		SearchId: o.SearchId,
		
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		QueryType: o.QueryType,
		
		IncludeDraftDocuments: o.IncludeDraftDocuments,
		
		Interval: o.Interval,
		
		Filter: o.Filter,
		
		SortOrder: o.SortOrder,
		
		SortBy: o.SortBy,
		
		Application: o.Application,
		
		ConversationContext: o.ConversationContext,
		
		ConfidenceThreshold: o.ConfidenceThreshold,
		
		AnswerHighlightTopResults: o.AnswerHighlightTopResults,
		
		AnswerMode: o.AnswerMode,
		
		PreprocessQuery: o.PreprocessQuery,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentsearchrequest) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentsearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentsearchrequestMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgedocumentsearchrequestMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if PageSize, ok := KnowledgedocumentsearchrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := KnowledgedocumentsearchrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if SearchId, ok := KnowledgedocumentsearchrequestMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if Total, ok := KnowledgedocumentsearchrequestMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PageCount, ok := KnowledgedocumentsearchrequestMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if QueryType, ok := KnowledgedocumentsearchrequestMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if IncludeDraftDocuments, ok := KnowledgedocumentsearchrequestMap["includeDraftDocuments"].(bool); ok {
		o.IncludeDraftDocuments = &IncludeDraftDocuments
	}
    
	if Interval, ok := KnowledgedocumentsearchrequestMap["interval"].(map[string]interface{}); ok {
		IntervalString, _ := json.Marshal(Interval)
		json.Unmarshal(IntervalString, &o.Interval)
	}
	
	if Filter, ok := KnowledgedocumentsearchrequestMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if SortOrder, ok := KnowledgedocumentsearchrequestMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
    
	if SortBy, ok := KnowledgedocumentsearchrequestMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
    
	if Application, ok := KnowledgedocumentsearchrequestMap["application"].(map[string]interface{}); ok {
		ApplicationString, _ := json.Marshal(Application)
		json.Unmarshal(ApplicationString, &o.Application)
	}
	
	if ConversationContext, ok := KnowledgedocumentsearchrequestMap["conversationContext"].(map[string]interface{}); ok {
		ConversationContextString, _ := json.Marshal(ConversationContext)
		json.Unmarshal(ConversationContextString, &o.ConversationContext)
	}
	
	if ConfidenceThreshold, ok := KnowledgedocumentsearchrequestMap["confidenceThreshold"].(float64); ok {
		ConfidenceThresholdFloat32 := float32(ConfidenceThreshold)
		o.ConfidenceThreshold = &ConfidenceThresholdFloat32
	}
	
	if AnswerHighlightTopResults, ok := KnowledgedocumentsearchrequestMap["answerHighlightTopResults"].(float64); ok {
		AnswerHighlightTopResultsInt := int(AnswerHighlightTopResults)
		o.AnswerHighlightTopResults = &AnswerHighlightTopResultsInt
	}
	
	if AnswerMode, ok := KnowledgedocumentsearchrequestMap["answerMode"].([]interface{}); ok {
		AnswerModeString, _ := json.Marshal(AnswerMode)
		json.Unmarshal(AnswerModeString, &o.AnswerMode)
	}
	
	if PreprocessQuery, ok := KnowledgedocumentsearchrequestMap["preprocessQuery"].(bool); ok {
		o.PreprocessQuery = &PreprocessQuery
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
