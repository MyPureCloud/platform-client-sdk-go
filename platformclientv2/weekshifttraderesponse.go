package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekshifttraderesponse
type Weekshifttraderesponse struct { 
	// Trade - The shift trade details
	Trade *Shifttraderesponse `json:"trade,omitempty"`


	// MatchReview - A preview of what the schedule would look like if the shift trade is approved plus any violations
	MatchReview *Shifttradematchreviewresponse `json:"matchReview,omitempty"`

}

func (u *Weekshifttraderesponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekshifttraderesponse

	

	return json.Marshal(&struct { 
		Trade *Shifttraderesponse `json:"trade,omitempty"`
		
		MatchReview *Shifttradematchreviewresponse `json:"matchReview,omitempty"`
		*Alias
	}{ 
		Trade: u.Trade,
		
		MatchReview: u.MatchReview,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Weekshifttraderesponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
