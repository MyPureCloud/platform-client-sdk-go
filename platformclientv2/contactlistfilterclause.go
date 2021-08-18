package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistfilterclause
type Contactlistfilterclause struct { 
	// FilterType - How to join predicates together.
	FilterType *string `json:"filterType,omitempty"`


	// Predicates - Conditions to filter the contacts by.
	Predicates *[]Contactlistfilterpredicate `json:"predicates,omitempty"`

}

func (u *Contactlistfilterclause) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistfilterclause

	

	return json.Marshal(&struct { 
		FilterType *string `json:"filterType,omitempty"`
		
		Predicates *[]Contactlistfilterpredicate `json:"predicates,omitempty"`
		*Alias
	}{ 
		FilterType: u.FilterType,
		
		Predicates: u.Predicates,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contactlistfilterclause) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
