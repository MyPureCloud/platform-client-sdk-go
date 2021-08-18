package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ring
type Ring struct { 
	// ExpansionCriteria
	ExpansionCriteria *[]Expansioncriterium `json:"expansionCriteria,omitempty"`


	// Actions
	Actions *Actions `json:"actions,omitempty"`

}

func (u *Ring) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ring

	

	return json.Marshal(&struct { 
		ExpansionCriteria *[]Expansioncriterium `json:"expansionCriteria,omitempty"`
		
		Actions *Actions `json:"actions,omitempty"`
		*Alias
	}{ 
		ExpansionCriteria: u.ExpansionCriteria,
		
		Actions: u.Actions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Ring) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
