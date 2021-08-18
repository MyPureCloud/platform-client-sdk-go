package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmmoveagentscompletetopicwfmmoveagentdata
type Wfmmoveagentscompletetopicwfmmoveagentdata struct { 
	// User
	User *Wfmmoveagentscompletetopicuserreference `json:"user,omitempty"`


	// Result
	Result *string `json:"result,omitempty"`

}

func (u *Wfmmoveagentscompletetopicwfmmoveagentdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmmoveagentscompletetopicwfmmoveagentdata

	

	return json.Marshal(&struct { 
		User *Wfmmoveagentscompletetopicuserreference `json:"user,omitempty"`
		
		Result *string `json:"result,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		Result: u.Result,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicwfmmoveagentdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
