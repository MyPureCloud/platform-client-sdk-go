package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailmailboxinfo
type Voicemailmailboxinfo struct { 
	// UsageSizeBytes - The total number of bytes for all voicemail message audio recordings
	UsageSizeBytes *int `json:"usageSizeBytes,omitempty"`


	// TotalCount - The total number of voicemail messages
	TotalCount *int `json:"totalCount,omitempty"`


	// UnreadCount - The total number of voicemail messages marked as unread
	UnreadCount *int `json:"unreadCount,omitempty"`


	// DeletedCount - The total number of voicemail messages marked as deleted
	DeletedCount *int `json:"deletedCount,omitempty"`


	// CreatedDate - The date of the oldest voicemail message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - The date of the most recent voicemail message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailmailboxinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
