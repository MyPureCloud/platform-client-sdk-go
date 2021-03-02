package platformclientv2
import (
	"encoding/json"
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
	return string(j)
}
