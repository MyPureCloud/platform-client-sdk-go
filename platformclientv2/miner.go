package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Miner) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Miner

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	ConversationsDateRangeStart := new(string)
	if u.ConversationsDateRangeStart != nil {
		*ConversationsDateRangeStart = timeutil.Strftime(u.ConversationsDateRangeStart, "%Y-%m-%d")
	} else {
		ConversationsDateRangeStart = nil
	}
	
	ConversationsDateRangeEnd := new(string)
	if u.ConversationsDateRangeEnd != nil {
		*ConversationsDateRangeEnd = timeutil.Strftime(u.ConversationsDateRangeEnd, "%Y-%m-%d")
	} else {
		ConversationsDateRangeEnd = nil
	}
	
	DateCompleted := new(string)
	if u.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(u.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	DateTriggered := new(string)
	if u.DateTriggered != nil {
		
		*DateTriggered = timeutil.Strftime(u.DateTriggered, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateTriggered = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ConversationsDateRangeStart *string `json:"conversationsDateRangeStart,omitempty"`
		
		ConversationsDateRangeEnd *string `json:"conversationsDateRangeEnd,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		ConversationDataUploaded *bool `json:"conversationDataUploaded,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		QueueIds *[]string `json:"queueIds,omitempty"`
		
		DateTriggered *string `json:"dateTriggered,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		LatestDraftVersion **Draft `json:"latestDraftVersion,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Language: u.Language,
		
		DateCreated: DateCreated,
		
		Status: u.Status,
		
		ConversationsDateRangeStart: ConversationsDateRangeStart,
		
		ConversationsDateRangeEnd: ConversationsDateRangeEnd,
		
		DateCompleted: DateCompleted,
		
		Message: u.Message,
		
		ConversationDataUploaded: u.ConversationDataUploaded,
		
		MediaType: u.MediaType,
		
		QueueIds: u.QueueIds,
		
		DateTriggered: DateTriggered,
		
		DateModified: DateModified,
		
		LatestDraftVersion: u.LatestDraftVersion,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Miner) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
