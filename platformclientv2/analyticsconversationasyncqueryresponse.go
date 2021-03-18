package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversationasyncqueryresponse
type Analyticsconversationasyncqueryresponse struct { 
	// Cursor - Optional cursor to indicate where to resume the results
	Cursor *string `json:"cursor,omitempty"`


	// DataAvailabilityDate - Data available up to at least this datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DataAvailabilityDate *time.Time `json:"dataAvailabilityDate,omitempty"`


	// Conversations
	Conversations *[]Analyticsconversation `json:"conversations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationasyncqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
