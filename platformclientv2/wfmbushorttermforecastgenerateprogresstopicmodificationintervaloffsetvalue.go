package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastgenerateprogresstopicmodificationintervaloffsetvalue
type Wfmbushorttermforecastgenerateprogresstopicmodificationintervaloffsetvalue struct { 
	// IntervalIndex
	IntervalIndex *int `json:"intervalIndex,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`

}

func (u *Wfmbushorttermforecastgenerateprogresstopicmodificationintervaloffsetvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastgenerateprogresstopicmodificationintervaloffsetvalue

	

	return json.Marshal(&struct { 
		IntervalIndex *int `json:"intervalIndex,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		*Alias
	}{ 
		IntervalIndex: u.IntervalIndex,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastgenerateprogresstopicmodificationintervaloffsetvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
