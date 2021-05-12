package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Createuser
type Createuser struct { 
	// Name - User's full name
	Name *string `json:"name,omitempty"`


	// Department
	Department *string `json:"department,omitempty"`


	// Email - User's email and username
	Email *string `json:"email,omitempty"`


	// Addresses - Email addresses and phone numbers for this user
	Addresses *[]Contact `json:"addresses,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// Password - User's password
	Password *string `json:"password,omitempty"`


	// DivisionId - The division to which this user will belong
	DivisionId *string `json:"divisionId,omitempty"`


	// State - Optional initialized state of the user. If not specified, state will be Active if invites are sent, otherwise Inactive.
	State *string `json:"state,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
