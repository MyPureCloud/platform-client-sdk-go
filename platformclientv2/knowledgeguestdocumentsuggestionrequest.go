package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeguestdocumentsuggestionrequest
type Knowledgeguestdocumentsuggestionrequest struct { 
	// Query - Query to get autocomplete suggestions for the matching knowledge documents.
	Query *string `json:"query,omitempty"`


	// PageSize - Page size of the returned results.
	PageSize *int `json:"pageSize,omitempty"`


	// IncludeDraftDocuments - Indicates whether the suggestion results would also include draft documents.
	IncludeDraftDocuments *bool `json:"includeDraftDocuments,omitempty"`

}

func (o *Knowledgeguestdocumentsuggestionrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeguestdocumentsuggestionrequest
	
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

func (o *Knowledgeguestdocumentsuggestionrequest) UnmarshalJSON(b []byte) error {
	var KnowledgeguestdocumentsuggestionrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeguestdocumentsuggestionrequestMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgeguestdocumentsuggestionrequestMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if PageSize, ok := KnowledgeguestdocumentsuggestionrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if IncludeDraftDocuments, ok := KnowledgeguestdocumentsuggestionrequestMap["includeDraftDocuments"].(bool); ok {
		o.IncludeDraftDocuments = &IncludeDraftDocuments
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentsuggestionrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
