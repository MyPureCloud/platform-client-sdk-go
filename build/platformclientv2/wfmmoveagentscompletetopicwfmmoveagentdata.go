package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicwfmmoveagentdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
