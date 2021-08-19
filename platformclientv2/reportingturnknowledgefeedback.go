package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturnknowledgefeedback
type Reportingturnknowledgefeedback struct { 
	// SearchId - The ID of the original knowledge search that this feedback relates to.
	SearchId *string `json:"searchId,omitempty"`


	// Rating - The feedback rating for the search (1.0 - 5.0). 1 = Negative, 5 = Positive.
	Rating *int `json:"rating,omitempty"`


	// Documents - The list of search documents that the feedback applies to.
	Documents *[]Reportingturnknowledgedocument `json:"documents,omitempty"`

}

func (u *Reportingturnknowledgefeedback) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturnknowledgefeedback

	

	return json.Marshal(&struct { 
		SearchId *string `json:"searchId,omitempty"`
		
		Rating *int `json:"rating,omitempty"`
		
		Documents *[]Reportingturnknowledgedocument `json:"documents,omitempty"`
		*Alias
	}{ 
		SearchId: u.SearchId,
		
		Rating: u.Rating,
		
		Documents: u.Documents,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgefeedback) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
