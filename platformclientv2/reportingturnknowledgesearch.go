package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturnknowledgesearch
type Reportingturnknowledgesearch struct { 
	// SearchId - The ID of this knowledge search.
	SearchId *string `json:"searchId,omitempty"`


	// Documents - The list of search documents captured during this reporting turn.
	Documents *[]Reportingturnknowledgedocument `json:"documents,omitempty"`


	// Query - The search query that was used to search the Knowledge Base documents for a matching question.
	Query *string `json:"query,omitempty"`

}

func (o *Reportingturnknowledgesearch) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturnknowledgesearch
	
	return json.Marshal(&struct { 
		SearchId *string `json:"searchId,omitempty"`
		
		Documents *[]Reportingturnknowledgedocument `json:"documents,omitempty"`
		
		Query *string `json:"query,omitempty"`
		*Alias
	}{ 
		SearchId: o.SearchId,
		
		Documents: o.Documents,
		
		Query: o.Query,
		Alias:    (*Alias)(o),
	})
}

func (o *Reportingturnknowledgesearch) UnmarshalJSON(b []byte) error {
	var ReportingturnknowledgesearchMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingturnknowledgesearchMap)
	if err != nil {
		return err
	}
	
	if SearchId, ok := ReportingturnknowledgesearchMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
	
	if Documents, ok := ReportingturnknowledgesearchMap["documents"].([]interface{}); ok {
		DocumentsString, _ := json.Marshal(Documents)
		json.Unmarshal(DocumentsString, &o.Documents)
	}
	
	if Query, ok := ReportingturnknowledgesearchMap["query"].(string); ok {
		o.Query = &Query
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgesearch) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
