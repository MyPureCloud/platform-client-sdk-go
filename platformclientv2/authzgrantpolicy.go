package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Authzgrantpolicy
type Authzgrantpolicy struct { 
	// Actions
	Actions *[]string `json:"actions,omitempty"`


	// Condition
	Condition *string `json:"condition,omitempty"`


	// Domain
	Domain *string `json:"domain,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`

}

func (u *Authzgrantpolicy) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Authzgrantpolicy

	

	return json.Marshal(&struct { 
		Actions *[]string `json:"actions,omitempty"`
		
		Condition *string `json:"condition,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		*Alias
	}{ 
		Actions: u.Actions,
		
		Condition: u.Condition,
		
		Domain: u.Domain,
		
		EntityName: u.EntityName,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Authzgrantpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
