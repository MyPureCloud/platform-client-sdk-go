package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmoduleruleparts
type Learningmoduleruleparts struct { 
	// Operation - The learning module rule operation
	Operation *string `json:"operation,omitempty"`


	// Selector - The learning module rule selector
	Selector *string `json:"selector,omitempty"`


	// Value - The value of rules
	Value *[]string `json:"value,omitempty"`


	// Order - The order of rules in learning module rule
	Order *int `json:"order,omitempty"`

}

func (u *Learningmoduleruleparts) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmoduleruleparts

	

	return json.Marshal(&struct { 
		Operation *string `json:"operation,omitempty"`
		
		Selector *string `json:"selector,omitempty"`
		
		Value *[]string `json:"value,omitempty"`
		
		Order *int `json:"order,omitempty"`
		*Alias
	}{ 
		Operation: u.Operation,
		
		Selector: u.Selector,
		
		Value: u.Value,
		
		Order: u.Order,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningmoduleruleparts) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
