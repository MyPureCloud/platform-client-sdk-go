package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentsuggestionrequest
type Knowledgedocumentsuggestionrequest struct { 
	// Query - Query to get autocomplete suggestions for the matching knowledge documents.
	Query *string `json:"query,omitempty"`


	// PageSize - Page size of the returned results.
	PageSize *int `json:"pageSize,omitempty"`


	// IncludeDraftDocuments - Indicates whether the suggestion results would also include draft documents.
	IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`

}

func (o *Knowledgedocumentsuggestionrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentsuggestionrequest
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`
		*Alias
	}{ 
		Query: o.Query,
		
		PageSize: o.PageSize,
		
		IncludeDraftDocuments: o.IncludeDraftDocuments,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentsuggestionrequest) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentsuggestionrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentsuggestionrequestMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgedocumentsuggestionrequestMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if PageSize, ok := KnowledgedocumentsuggestionrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if IncludeDraftDocuments, ok := KnowledgedocumentsuggestionrequestMap["includeDraftDocuments"].(bool); ok {
		o.IncludeDraftDocuments = &IncludeDraftDocuments
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsuggestionrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
