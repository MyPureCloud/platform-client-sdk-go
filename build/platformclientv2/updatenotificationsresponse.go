package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Updatenotificationsresponse
type Updatenotificationsresponse struct { 
	// Entities
	Entities *[]Updatenotificationresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updatenotificationsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
