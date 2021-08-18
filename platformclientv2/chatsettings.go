package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatsettings
type Chatsettings struct { 
	// MessageRetentionPeriodDays - Retention time for messages in days
	MessageRetentionPeriodDays *int `json:"messageRetentionPeriodDays,omitempty"`

}

func (u *Chatsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Chatsettings

	

	return json.Marshal(&struct { 
		MessageRetentionPeriodDays *int `json:"messageRetentionPeriodDays,omitempty"`
		*Alias
	}{ 
		MessageRetentionPeriodDays: u.MessageRetentionPeriodDays,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Chatsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
