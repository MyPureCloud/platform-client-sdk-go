package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingjobsquery
type Recordingjobsquery struct { 
	// Action - Operation to perform bulk task
	Action *string `json:"action,omitempty"`


	// ActionDate - The date when the action will be performed. If the operation will cause the delete date of a recording to be older than the export date, the export date will be adjusted to the delete date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ActionDate *time.Time `json:"actionDate,omitempty"`


	// IntegrationId - IntegrationId to Access AWS S3 bucket for bulk recording exports. This field is required and used only for EXPORT action.
	IntegrationId *string `json:"integrationId,omitempty"`


	// IncludeScreenRecordings - Include Screen recordings for export action, default value = true 
	IncludeScreenRecordings *bool `json:"includeScreenRecordings,omitempty"`


	// ConversationQuery - Conversation Query. Note: After the recording is created, it might take up to 48 hours for the recording to be included in the submitted job query.
	ConversationQuery *Asyncconversationquery `json:"conversationQuery,omitempty"`

}

func (u *Recordingjobsquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingjobsquery

	
	ActionDate := new(string)
	if u.ActionDate != nil {
		
		*ActionDate = timeutil.Strftime(u.ActionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ActionDate = nil
	}
	

	return json.Marshal(&struct { 
		Action *string `json:"action,omitempty"`
		
		ActionDate *string `json:"actionDate,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		IncludeScreenRecordings *bool `json:"includeScreenRecordings,omitempty"`
		
		ConversationQuery *Asyncconversationquery `json:"conversationQuery,omitempty"`
		*Alias
	}{ 
		Action: u.Action,
		
		ActionDate: ActionDate,
		
		IntegrationId: u.IntegrationId,
		
		IncludeScreenRecordings: u.IncludeScreenRecordings,
		
		ConversationQuery: u.ConversationQuery,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Recordingjobsquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
