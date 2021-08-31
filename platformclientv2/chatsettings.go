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

func (o *Chatsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Chatsettings
	
	return json.Marshal(&struct { 
		MessageRetentionPeriodDays *int `json:"messageRetentionPeriodDays,omitempty"`
		*Alias
	}{ 
		MessageRetentionPeriodDays: o.MessageRetentionPeriodDays,
		Alias:    (*Alias)(o),
	})
}

func (o *Chatsettings) UnmarshalJSON(b []byte) error {
	var ChatsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ChatsettingsMap)
	if err != nil {
		return err
	}
	
	if MessageRetentionPeriodDays, ok := ChatsettingsMap["messageRetentionPeriodDays"].(float64); ok {
		MessageRetentionPeriodDaysInt := int(MessageRetentionPeriodDays)
		o.MessageRetentionPeriodDays = &MessageRetentionPeriodDaysInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chatsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
