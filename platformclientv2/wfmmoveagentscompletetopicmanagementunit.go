package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmmoveagentscompletetopicmanagementunit
type Wfmmoveagentscompletetopicmanagementunit struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmmoveagentscompletetopicmanagementunit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
