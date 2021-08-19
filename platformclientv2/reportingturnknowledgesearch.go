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

func (u *Reportingturnknowledgesearch) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturnknowledgesearch

	

	return json.Marshal(&struct { 
		SearchId *string `json:"searchId,omitempty"`
		
		Documents *[]Reportingturnknowledgedocument `json:"documents,omitempty"`
		
		Query *string `json:"query,omitempty"`
		*Alias
	}{ 
		SearchId: u.SearchId,
		
		Documents: u.Documents,
		
		Query: u.Query,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgesearch) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
