package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentsearchrequest
type Knowledgedocumentsearchrequest struct { 
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

}

func (o *Knowledgedocumentsearchrequest) MarshalJSON() ([]byte, error) {
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
		
		IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`
		
		Interval *Documentqueryinterval `json:"interval,omitempty"`
		
		Filter *Documentquery `json:"filter,omitempty"`
		
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		Application *Knowledgesearchclientapplication `json:"application,omitempty"`
		*Alias
	}{ 
		Query: o.Query,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		SearchId: o.SearchId,
		
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		IncludeDraftDocuments: o.IncludeDraftDocuments,
		
		Interval: o.Interval,
		
		Filter: o.Filter,
		
		SortOrder: o.SortOrder,
		
		SortBy: o.SortBy,
		
		Application: o.Application,
		Alias:    (*Alias)(o),
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
