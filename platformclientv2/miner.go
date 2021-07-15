package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Miner
type Miner struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Chat Corpus Name.
	Name *string `json:"name,omitempty"`


	// Language - Language Localization code.
	Language *string `json:"language,omitempty"`


	// DateCreated - Date when the miner was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// Status - Status of the miner.
	Status *string `json:"status,omitempty"`


	// ConversationsDateRangeStart - Date from which the conversations need to be taken for mining. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	ConversationsDateRangeStart *time.Time `json:"conversationsDateRangeStart,omitempty"`


	// ConversationsDateRangeEnd - Date till which the conversations need to be taken for mining. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	ConversationsDateRangeEnd *time.Time `json:"conversationsDateRangeEnd,omitempty"`


	// DateCompleted - Date when the mining process was completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// Message - Mining message if present.
	Message *string `json:"message,omitempty"`


	// ConversationDataUploaded - Flag to indicate whether data file to be mined was uploaded.
	ConversationDataUploaded *bool `json:"conversationDataUploaded,omitempty"`


	// MediaType - Media type for filtering conversations.
	MediaType *string `json:"mediaType,omitempty"`


	// QueueIds - List of queue IDs for filtering conversations.
	QueueIds *[]string `json:"queueIds,omitempty"`


	// DateTriggered - Date when the miner started execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateTriggered *time.Time `json:"dateTriggered,omitempty"`


	// DateModified - Date when the miner was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// LatestDraftVersion - Latest draft details of the miner.
	LatestDraftVersion **Draft `json:"latestDraftVersion,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Miner) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
