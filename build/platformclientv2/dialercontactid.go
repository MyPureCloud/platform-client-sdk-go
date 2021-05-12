package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactid
type Dialercontactid struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ContactListId
	ContactListId *string `json:"contactListId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercontactid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
