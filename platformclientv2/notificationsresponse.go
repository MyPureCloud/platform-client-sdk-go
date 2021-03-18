package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Notificationsresponse
type Notificationsresponse struct { 
	// Entities
	Entities *[]Wfmusernotification `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Notificationsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
