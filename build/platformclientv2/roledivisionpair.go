package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Roledivisionpair
type Roledivisionpair struct { 
	// RoleId - The ID of the role
	RoleId *string `json:"roleId,omitempty"`


	// DivisionId - The ID of the division
	DivisionId *string `json:"divisionId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Roledivisionpair) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
