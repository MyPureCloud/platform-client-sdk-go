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

func (o *Voicemailmailboxinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailmailboxinfo
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		UsageSizeBytes: o.UsageSizeBytes,
		
		TotalCount: o.TotalCount,
		
		UnreadCount: o.UnreadCount,
		
		DeletedCount: o.DeletedCount,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Voicemailmailboxinfo) UnmarshalJSON(b []byte) error {
	var VoicemailmailboxinfoMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailmailboxinfoMap)
	if err != nil {
		return err
	}
	
	if UsageSizeBytes, ok := VoicemailmailboxinfoMap["usageSizeBytes"].(float64); ok {
		UsageSizeBytesInt := int(UsageSizeBytes)
		o.UsageSizeBytes = &UsageSizeBytesInt
	}
	
	if TotalCount, ok := VoicemailmailboxinfoMap["totalCount"].(float64); ok {
		TotalCountInt := int(TotalCount)
		o.TotalCount = &TotalCountInt
	}
	
	if UnreadCount, ok := VoicemailmailboxinfoMap["unreadCount"].(float64); ok {
		UnreadCountInt := int(UnreadCount)
		o.UnreadCount = &UnreadCountInt
	}
	
	if DeletedCount, ok := VoicemailmailboxinfoMap["deletedCount"].(float64); ok {
		DeletedCountInt := int(DeletedCount)
		o.DeletedCount = &DeletedCountInt
	}
	
	if createdDateString, ok := VoicemailmailboxinfoMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := VoicemailmailboxinfoMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailmailboxinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
