package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregatequeryrequestpredicate
type Developmentactivityaggregatequeryrequestpredicate struct { 
	// Dimension - Each predicates specifies a dimension.
	Dimension *string `json:"dimension,omitempty"`


	// Value - Corresponding value for dimensions in predicates. If the dimensions is type, Valid Values: Informational, Coaching
	Value *string `json:"value,omitempty"`

}

func (u *Developmentactivityaggregatequeryrequestpredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivityaggregatequeryrequestpredicate

	

	return json.Marshal(&struct { 
		Dimension *string `json:"dimension,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Dimension: u.Dimension,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryrequestpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
