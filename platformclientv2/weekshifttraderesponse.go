package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Weekshifttraderesponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
