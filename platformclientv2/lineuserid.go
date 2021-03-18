package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Lineuserid - Channel-specific User ID for Line accounts
type Lineuserid struct { 
	// UserId - The unique channel-specific userId for the user
	UserId *string `json:"userId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Lineuserid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
