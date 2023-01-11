package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentguestsearch
type Knowledgedocumentguestsearch struct { 
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


	// SessionId - Session ID of the search.
	SessionId *string `json:"sessionId,omitempty"`


	// Results - Documents that matched the search query.
	Results *[]Knowledgedocumentguestsearchresult `json:"results,omitempty"`

}

func (o *Knowledgedocumentguestsearch) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentguestsearch
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		SearchId *string `json:"searchId,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		QueryType *string `json:"queryType,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		Results *[]Knowledgedocumentguestsearchresult `json:"results,omitempty"`
		*Alias
	}{ 
		Query: o.Query,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		SearchId: o.SearchId,
		
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		QueryType: o.QueryType,
		
		SessionId: o.SessionId,
		
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentguestsearch) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentguestsearchMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentguestsearchMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgedocumentguestsearchMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if PageSize, ok := KnowledgedocumentguestsearchMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := KnowledgedocumentguestsearchMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if SearchId, ok := KnowledgedocumentguestsearchMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if Total, ok := KnowledgedocumentguestsearchMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PageCount, ok := KnowledgedocumentguestsearchMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if QueryType, ok := KnowledgedocumentguestsearchMap["queryType"].(string); ok {
		o.QueryType = &QueryType
	}
    
	if SessionId, ok := KnowledgedocumentguestsearchMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if Results, ok := KnowledgedocumentguestsearchMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentguestsearch) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
