package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatbadgetopicchatbadge
type Chatbadgetopicchatbadge struct { 
	// Entity
	Entity *Chatbadgetopicbadgeentity `json:"entity,omitempty"`


	// UnreadCount
	UnreadCount *int `json:"unreadCount,omitempty"`


	// LastUnreadNotificationDate
	LastUnreadNotificationDate *time.Time `json:"lastUnreadNotificationDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Chatbadgetopicchatbadge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
