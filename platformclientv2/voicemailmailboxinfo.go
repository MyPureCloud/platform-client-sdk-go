package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Voicemailmailboxinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailmailboxinfo

	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		UsageSizeBytes *int `json:"usageSizeBytes,omitempty"`
		
		TotalCount *int `json:"totalCount,omitempty"`
		
		UnreadCount *int `json:"unreadCount,omitempty"`
		
		DeletedCount *int `json:"deletedCount,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		UsageSizeBytes: u.UsageSizeBytes,
		
		TotalCount: u.TotalCount,
		
		UnreadCount: u.UnreadCount,
		
		DeletedCount: u.DeletedCount,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Voicemailmailboxinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
