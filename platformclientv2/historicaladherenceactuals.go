package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicaladherenceactuals
type Historicaladherenceactuals struct { 
	// ActualActivityCategory - Activity in which the user is actually engaged
	ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`


	// StartOffsetSeconds - Actual start offset in seconds relative to query start time
	StartOffsetSeconds *int `json:"startOffsetSeconds,omitempty"`


	// EndOffsetSeconds - Actual end offset in seconds relative to query start time
	EndOffsetSeconds *int `json:"endOffsetSeconds,omitempty"`

}

func (u *Historicaladherenceactuals) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicaladherenceactuals

	

	return json.Marshal(&struct { 
		ActualActivityCategory *string `json:"actualActivityCategory,omitempty"`
		
		StartOffsetSeconds *int `json:"startOffsetSeconds,omitempty"`
		
		EndOffsetSeconds *int `json:"endOffsetSeconds,omitempty"`
		*Alias
	}{ 
		ActualActivityCategory: u.ActualActivityCategory,
		
		StartOffsetSeconds: u.StartOffsetSeconds,
		
		EndOffsetSeconds: u.EndOffsetSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Historicaladherenceactuals) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
