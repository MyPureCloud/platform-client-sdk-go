package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Crossplatformchatmediapolicy
type Crossplatformchatmediapolicy struct { 
	// Actions - Actions applied when specified conditions are met
	Actions *Crossplatformpolicyactions `json:"actions,omitempty"`


	// Conditions - Conditions for when actions should be applied
	Conditions *Chatmediapolicyconditions `json:"conditions,omitempty"`

}

func (u *Crossplatformchatmediapolicy) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Crossplatformchatmediapolicy

	

	return json.Marshal(&struct { 
		Actions *Crossplatformpolicyactions `json:"actions,omitempty"`
		
		Conditions *Chatmediapolicyconditions `json:"conditions,omitempty"`
		*Alias
	}{ 
		Actions: u.Actions,
		
		Conditions: u.Conditions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Crossplatformchatmediapolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
