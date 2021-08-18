package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversationqueryresponse
type Analyticsconversationqueryresponse struct { 
	// Aggregations
	Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`


	// Conversations
	Conversations *[]Analyticsconversationwithoutattributes `json:"conversations,omitempty"`


	// TotalHits
	TotalHits *int `json:"totalHits,omitempty"`

}

func (u *Analyticsconversationqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsconversationqueryresponse

	

	return json.Marshal(&struct { 
		Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`
		
		Conversations *[]Analyticsconversationwithoutattributes `json:"conversations,omitempty"`
		
		TotalHits *int `json:"totalHits,omitempty"`
		*Alias
	}{ 
		Aggregations: u.Aggregations,
		
		Conversations: u.Conversations,
		
		TotalHits: u.TotalHits,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsconversationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
