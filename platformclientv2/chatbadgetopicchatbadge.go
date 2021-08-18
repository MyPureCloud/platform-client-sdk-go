package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Chatbadgetopicchatbadge) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Chatbadgetopicchatbadge

	
	LastUnreadNotificationDate := new(string)
	if u.LastUnreadNotificationDate != nil {
		
		*LastUnreadNotificationDate = timeutil.Strftime(u.LastUnreadNotificationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastUnreadNotificationDate = nil
	}
	

	return json.Marshal(&struct { 
		Entity *Chatbadgetopicbadgeentity `json:"entity,omitempty"`
		
		UnreadCount *int `json:"unreadCount,omitempty"`
		
		LastUnreadNotificationDate *string `json:"lastUnreadNotificationDate,omitempty"`
		*Alias
	}{ 
		Entity: u.Entity,
		
		UnreadCount: u.UnreadCount,
		
		LastUnreadNotificationDate: LastUnreadNotificationDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Chatbadgetopicchatbadge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
