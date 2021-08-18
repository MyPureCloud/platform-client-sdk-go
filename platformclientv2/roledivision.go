package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Roledivision
type Roledivision struct { 
	// RoleId - Role to be associated with the given division which forms a grant
	RoleId *string `json:"roleId,omitempty"`


	// DivisionId - Division associated with the given role which forms a grant
	DivisionId *string `json:"divisionId,omitempty"`

}

func (u *Roledivision) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Roledivision

	

	return json.Marshal(&struct { 
		RoleId *string `json:"roleId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		*Alias
	}{ 
		RoleId: u.RoleId,
		
		DivisionId: u.DivisionId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Roledivision) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
