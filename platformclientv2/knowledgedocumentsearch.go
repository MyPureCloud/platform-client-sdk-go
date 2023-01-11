package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentsearch
type Knowledgedocumentsearch struct { 
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


	// Results - Documents matching the search query.
	Results *[]Knowledgedocumentsearchresult `json:"results,omitempty"`


	// Application - The client application details from which search happened.
	Application *Knowledgesearchclientapplication `json:"application,omitempty"`


	// ConversationContext - Conversation context information if the search is initiated in the context of a conversation.
	ConversationContext *Knowledgeconversationcontextresponse `json:"conversationContext,omitempty"`

}

func (o *Knowledgedocumentsearch) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentsearch
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		Results *[]Knowledgedocumentsearchresult `json:"results,omitempty"`
		
		Application *Knowledgesearchclientapplication `json:"application,omitempty"`
		
		ConversationContext *Knowledgeconversationcontextresponse `json:"conversationContext,omitempty"`
		*Alias
	}{ 
		Query: o.Query,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		SearchId: o.SearchId,
		
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		QueryType: o.QueryType,
		
		Results: o.Results,
		
		Application: o.Application,
		
		ConversationContext: o.ConversationContext,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentsearch) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentsearchMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentsearchMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgedocumentsearchMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if PageSize, ok := KnowledgedocumentsearchMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := KnowledgedocumentsearchMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if SearchId, ok := KnowledgedocumentsearchMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if Total, ok := KnowledgedocumentsearchMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PageCount, ok := KnowledgedocumentsearchMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if QueryType, ok := KnowledgedocumentsearchMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if Results, ok := KnowledgedocumentsearchMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if Application, ok := KnowledgedocumentsearchMap["application"].(map[string]interface{}); ok {
		ApplicationString, _ := json.Marshal(Application)
		json.Unmarshal(ApplicationString, &o.Application)
	}
	
	if ConversationContext, ok := KnowledgedocumentsearchMap["conversationContext"].(map[string]interface{}); ok {
		ConversationContextString, _ := json.Marshal(ConversationContext)
		json.Unmarshal(ConversationContextString, &o.ConversationContext)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsearch) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
