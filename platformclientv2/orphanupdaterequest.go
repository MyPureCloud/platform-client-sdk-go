package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Orphanupdaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Orphanupdaterequest

	
	ArchiveDate := new(string)
	if u.ArchiveDate != nil {
		
		*ArchiveDate = timeutil.Strftime(u.ArchiveDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ArchiveDate = nil
	}
	
	DeleteDate := new(string)
	if u.DeleteDate != nil {
		
		*DeleteDate = timeutil.Strftime(u.DeleteDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DeleteDate = nil
	}
	

	return json.Marshal(&struct { 
		ArchiveDate *string `json:"archiveDate,omitempty"`
		
		DeleteDate *string `json:"deleteDate,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		*Alias
	}{ 
		ArchiveDate: ArchiveDate,
		
		DeleteDate: DeleteDate,
		
		ConversationId: u.ConversationId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Orphanupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
