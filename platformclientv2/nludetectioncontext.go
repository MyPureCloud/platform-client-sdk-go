package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectioncontext
type Nludetectioncontext struct { 
	// Intent - Restrict detection to this intent.
	Intent *Contextintent `json:"intent,omitempty"`


	// Entity - Use this entity to restrict detection.
	Entity *Contextentity `json:"entity,omitempty"`

}

func (u *Nludetectioncontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludetectioncontext

	

	return json.Marshal(&struct { 
		Intent *Contextintent `json:"intent,omitempty"`
		
		Entity *Contextentity `json:"entity,omitempty"`
		*Alias
	}{ 
		Intent: u.Intent,
		
		Entity: u.Entity,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nludetectioncontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
