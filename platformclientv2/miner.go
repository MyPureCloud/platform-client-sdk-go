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

func (o *Miner) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Miner
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	ConversationsDateRangeStart := new(string)
	if o.ConversationsDateRangeStart != nil {
		*ConversationsDateRangeStart = timeutil.Strftime(o.ConversationsDateRangeStart, "%Y-%m-%d")
	} else {
		ConversationsDateRangeStart = nil
	}
	
	ConversationsDateRangeEnd := new(string)
	if o.ConversationsDateRangeEnd != nil {
		*ConversationsDateRangeEnd = timeutil.Strftime(o.ConversationsDateRangeEnd, "%Y-%m-%d")
	} else {
		ConversationsDateRangeEnd = nil
	}
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	DateTriggered := new(string)
	if o.DateTriggered != nil {
		
		*DateTriggered = timeutil.Strftime(o.DateTriggered, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateTriggered = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		Language: o.Language,
		
		DateCreated: DateCreated,
		
		Status: o.Status,
		
		ConversationsDateRangeStart: ConversationsDateRangeStart,
		
		ConversationsDateRangeEnd: ConversationsDateRangeEnd,
		
		DateCompleted: DateCompleted,
		
		Message: o.Message,
		
		ConversationDataUploaded: o.ConversationDataUploaded,
		
		MediaType: o.MediaType,
		
		QueueIds: o.QueueIds,
		
		DateTriggered: DateTriggered,
		
		DateModified: DateModified,
		
		LatestDraftVersion: o.LatestDraftVersion,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Miner) UnmarshalJSON(b []byte) error {
	var MinerMap map[string]interface{}
	err := json.Unmarshal(b, &MinerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MinerMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := MinerMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Language, ok := MinerMap["language"].(string); ok {
		o.Language = &Language
	}
	
	if dateCreatedString, ok := MinerMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if Status, ok := MinerMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if conversationsDateRangeStartString, ok := MinerMap["conversationsDateRangeStart"].(string); ok {
		ConversationsDateRangeStart, _ := time.Parse("2006-01-02", conversationsDateRangeStartString)
		o.ConversationsDateRangeStart = &ConversationsDateRangeStart
	}
	
	if conversationsDateRangeEndString, ok := MinerMap["conversationsDateRangeEnd"].(string); ok {
		ConversationsDateRangeEnd, _ := time.Parse("2006-01-02", conversationsDateRangeEndString)
		o.ConversationsDateRangeEnd = &ConversationsDateRangeEnd
	}
	
	if dateCompletedString, ok := MinerMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if Message, ok := MinerMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if ConversationDataUploaded, ok := MinerMap["conversationDataUploaded"].(bool); ok {
		o.ConversationDataUploaded = &ConversationDataUploaded
	}
	
	if MediaType, ok := MinerMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if QueueIds, ok := MinerMap["queueIds"].([]interface{}); ok {
		QueueIdsString, _ := json.Marshal(QueueIds)
		json.Unmarshal(QueueIdsString, &o.QueueIds)
	}
	
	if dateTriggeredString, ok := MinerMap["dateTriggered"].(string); ok {
		DateTriggered, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateTriggeredString)
		o.DateTriggered = &DateTriggered
	}
	
	if dateModifiedString, ok := MinerMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if LatestDraftVersion, ok := MinerMap["latestDraftVersion"].(map[string]interface{}); ok {
		LatestDraftVersionString, _ := json.Marshal(LatestDraftVersion)
		json.Unmarshal(LatestDraftVersionString, &o.LatestDraftVersion)
	}
	
	if SelfUri, ok := MinerMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Miner) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
