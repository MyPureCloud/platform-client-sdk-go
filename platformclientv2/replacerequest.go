package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Replacerequest
type Replacerequest struct { 
	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// AuthToken
	AuthToken *string `json:"authToken,omitempty"`

}

func (u *Replacerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Replacerequest

	

	return json.Marshal(&struct { 
		ChangeNumber *int `json:"changeNumber,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		AuthToken *string `json:"authToken,omitempty"`
		*Alias
	}{ 
		ChangeNumber: u.ChangeNumber,
		
		Name: u.Name,
		
		AuthToken: u.AuthToken,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Replacerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
