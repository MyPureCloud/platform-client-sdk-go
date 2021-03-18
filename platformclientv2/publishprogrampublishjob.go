package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Publishprogrampublishjob
type Publishprogrampublishjob struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`

}

// String returns a JSON representation of the model
func (o *Publishprogrampublishjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
