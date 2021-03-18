package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Updatenotificationsrequest
type Updatenotificationsrequest struct { 
	// Entities - The notifications to update
	Entities *[]Wfmusernotification `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updatenotificationsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
