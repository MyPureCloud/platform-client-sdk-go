package platformclientv2
import (
	"time"
	"encoding/json"
)

// Orphanupdaterequest
type Orphanupdaterequest struct { 
	// ArchiveDate - The orphan recording's archive date. Must be greater than 1 day from now if set. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ArchiveDate *time.Time `json:"archiveDate,omitempty"`


	// DeleteDate - The orphan recording's delete date. Must be greater than archiveDate if set, otherwise one day from now. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DeleteDate *time.Time `json:"deleteDate,omitempty"`


	// ConversationId - A conversation Id that this orphan's recording is to be attached to. If not present, the conversationId will be deduced from the recording media.
	ConversationId *string `json:"conversationId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Orphanupdaterequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
