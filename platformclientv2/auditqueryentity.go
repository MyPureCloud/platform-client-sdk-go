package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryentity
type Auditqueryentity struct { 
	// Name - Name of the Entity
	Name *string `json:"name,omitempty"`


	// Actions - List of Actions
	Actions *[]string `json:"actions,omitempty"`

}

func (u *Auditqueryentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditqueryentity

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Actions *[]string `json:"actions,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Actions: u.Actions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditqueryentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
