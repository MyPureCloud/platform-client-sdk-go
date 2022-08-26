package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentsuggestion
type Knowledgedocumentsuggestion struct { 
	// Query - Query to get autocomplete suggestions for the matching knowledge documents.
	Query *string `json:"query,omitempty"`


	// PageSize - Page size of the returned results.
	PageSize *int `json:"pageSize,omitempty"`


	// Results - Documents matching to the autocomplete suggestions query.
	Results *[]Knowledgedocumentsuggestionresult `json:"results,omitempty"`

}

func (o *Knowledgedocumentsuggestion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentsuggestion
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		Results *[]Knowledgedocumentsuggestionresult `json:"results,omitempty"`
		*Alias
	}{ 
		Query: o.Query,
		
		PageSize: o.PageSize,
		
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentsuggestion) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentsuggestionMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentsuggestionMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgedocumentsuggestionMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if PageSize, ok := KnowledgedocumentsuggestionMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if Results, ok := KnowledgedocumentsuggestionMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsuggestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
