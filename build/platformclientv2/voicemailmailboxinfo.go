package platformclientv2
import (
	"time"
	"encoding/json"
)

// Voicemailmailboxinfo
type Voicemailmailboxinfo struct { 
	// UsageSizeBytes - The total number of bytes for all voicemail message audio recordings
	UsageSizeBytes *int64 `json:"usageSizeBytes,omitempty"`


	// TotalCount - The total number of voicemail messages
	TotalCount *int32 `json:"totalCount,omitempty"`


	// UnreadCount - The total number of voicemail messages marked as unread
	UnreadCount *int32 `json:"unreadCount,omitempty"`


	// DeletedCount - The total number of voicemail messages marked as deleted
	DeletedCount *int32 `json:"deletedCount,omitempty"`


	// CreatedDate - The date of the oldest voicemail message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - The date of the most recent voicemail message. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailmailboxinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
