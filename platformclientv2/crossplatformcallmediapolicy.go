package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Crossplatformcallmediapolicy
type Crossplatformcallmediapolicy struct { 
	// Actions - Actions applied when specified conditions are met
	Actions *Crossplatformpolicyactions `json:"actions,omitempty"`


	// Conditions - Conditions for when actions should be applied
	Conditions *Callmediapolicyconditions `json:"conditions,omitempty"`

}

func (u *Crossplatformcallmediapolicy) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Crossplatformcallmediapolicy

	

	return json.Marshal(&struct { 
		Actions *Crossplatformpolicyactions `json:"actions,omitempty"`
		
		Conditions *Callmediapolicyconditions `json:"conditions,omitempty"`
		*Alias
	}{ 
		Actions: u.Actions,
		
		Conditions: u.Conditions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Crossplatformcallmediapolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
