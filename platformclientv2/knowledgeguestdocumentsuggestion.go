package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeguestdocumentsuggestion
type Knowledgeguestdocumentsuggestion struct { 
	// Query - Query to get autocomplete suggestions for the matching knowledge documents.
	Query *string `json:"query,omitempty"`


	// PageSize - Page size of the returned results.
	PageSize *int `json:"pageSize,omitempty"`


	// SessionId - Session ID of the guest suggestions.
	SessionId *string `json:"sessionId,omitempty"`


	// Results - Suggestions matching the query.
	Results *[]Knowledgeguestdocumentsuggestionresult `json:"results,omitempty"`

}

func (o *Knowledgeguestdocumentsuggestion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeguestdocumentsuggestion
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		Results *[]Knowledgeguestdocumentsuggestionresult `json:"results,omitempty"`
		*Alias
	}{ 
		Query: o.Query,
		
		PageSize: o.PageSize,
		
		SessionId: o.SessionId,
		
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeguestdocumentsuggestion) UnmarshalJSON(b []byte) error {
	var KnowledgeguestdocumentsuggestionMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeguestdocumentsuggestionMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgeguestdocumentsuggestionMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if PageSize, ok := KnowledgeguestdocumentsuggestionMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if SessionId, ok := KnowledgeguestdocumentsuggestionMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if Results, ok := KnowledgeguestdocumentsuggestionMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentsuggestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
