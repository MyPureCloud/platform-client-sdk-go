package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Unreadstatus
type Unreadstatus struct { 
	// Unread - Sets if the alert is read or unread.
	Unread *bool `json:"unread,omitempty"`

}

func (o *Unreadstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Unreadstatus
	
	return json.Marshal(&struct { 
		Unread *bool `json:"unread,omitempty"`
		*Alias
	}{ 
		Unread: o.Unread,
		Alias:    (*Alias)(o),
	})
}

func (o *Unreadstatus) UnmarshalJSON(b []byte) error {
	var UnreadstatusMap map[string]interface{}
	err := json.Unmarshal(b, &UnreadstatusMap)
	if err != nil {
		return err
	}
	
	if Unread, ok := UnreadstatusMap["unread"].(bool); ok {
		o.Unread = &Unread
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Unreadstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
