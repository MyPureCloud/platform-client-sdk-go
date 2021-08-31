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

func (o *Chatbadgetopicchatbadge) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Chatbadgetopicchatbadge
	
	LastUnreadNotificationDate := new(string)
	if o.LastUnreadNotificationDate != nil {
		
		*LastUnreadNotificationDate = timeutil.Strftime(o.LastUnreadNotificationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastUnreadNotificationDate = nil
	}
	
	return json.Marshal(&struct { 
		Entity *Chatbadgetopicbadgeentity `json:"entity,omitempty"`
		
		UnreadCount *int `json:"unreadCount,omitempty"`
		
		LastUnreadNotificationDate *string `json:"lastUnreadNotificationDate,omitempty"`
		*Alias
	}{ 
		Entity: o.Entity,
		
		UnreadCount: o.UnreadCount,
		
		LastUnreadNotificationDate: LastUnreadNotificationDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Chatbadgetopicchatbadge) UnmarshalJSON(b []byte) error {
	var ChatbadgetopicchatbadgeMap map[string]interface{}
	err := json.Unmarshal(b, &ChatbadgetopicchatbadgeMap)
	if err != nil {
		return err
	}
	
	if Entity, ok := ChatbadgetopicchatbadgeMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if UnreadCount, ok := ChatbadgetopicchatbadgeMap["unreadCount"].(float64); ok {
		UnreadCountInt := int(UnreadCount)
		o.UnreadCount = &UnreadCountInt
	}
	
	if lastUnreadNotificationDateString, ok := ChatbadgetopicchatbadgeMap["lastUnreadNotificationDate"].(string); ok {
		LastUnreadNotificationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", lastUnreadNotificationDateString)
		o.LastUnreadNotificationDate = &LastUnreadNotificationDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chatbadgetopicchatbadge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
