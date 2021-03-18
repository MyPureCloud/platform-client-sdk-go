package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Historicaladherenceactuals) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
