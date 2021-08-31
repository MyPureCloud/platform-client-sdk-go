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

func (o *Weekshifttraderesponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekshifttraderesponse
	
	return json.Marshal(&struct { 
		Trade *Shifttraderesponse `json:"trade,omitempty"`
		
		MatchReview *Shifttradematchreviewresponse `json:"matchReview,omitempty"`
		*Alias
	}{ 
		Trade: o.Trade,
		
		MatchReview: o.MatchReview,
		Alias:    (*Alias)(o),
	})
}

func (o *Weekshifttraderesponse) UnmarshalJSON(b []byte) error {
	var WeekshifttraderesponseMap map[string]interface{}
	err := json.Unmarshal(b, &WeekshifttraderesponseMap)
	if err != nil {
		return err
	}
	
	if Trade, ok := WeekshifttraderesponseMap["trade"].(map[string]interface{}); ok {
		TradeString, _ := json.Marshal(Trade)
		json.Unmarshal(TradeString, &o.Trade)
	}
	
	if MatchReview, ok := WeekshifttraderesponseMap["matchReview"].(map[string]interface{}); ok {
		MatchReviewString, _ := json.Marshal(MatchReview)
		json.Unmarshal(MatchReviewString, &o.MatchReview)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Weekshifttraderesponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
