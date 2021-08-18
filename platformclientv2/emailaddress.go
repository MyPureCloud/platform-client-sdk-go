package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailaddress
type Emailaddress struct { 
	// Email
	Email *string `json:"email,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

func (u *Emailaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailaddress

	

	return json.Marshal(&struct { 
		Email *string `json:"email,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Email: u.Email,
		
		Name: u.Name,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Emailaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
