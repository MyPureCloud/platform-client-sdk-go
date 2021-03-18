package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Chatsettings
type Chatsettings struct { 
	// MessageRetentionPeriodDays - Retention time for messages in days
	MessageRetentionPeriodDays *int `json:"messageRetentionPeriodDays,omitempty"`

}

// String returns a JSON representation of the model
func (o *Chatsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
