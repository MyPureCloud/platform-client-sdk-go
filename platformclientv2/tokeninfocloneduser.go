package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Tokeninfocloneduser
type Tokeninfocloneduser struct { 
	// Id - User id of the original native user
	Id *string `json:"id,omitempty"`


	// Organization - Organization of the original native user
	Organization *Entity `json:"organization,omitempty"`

}

func (u *Tokeninfocloneduser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Tokeninfocloneduser

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Organization *Entity `json:"organization,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Organization: u.Organization,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Tokeninfocloneduser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
