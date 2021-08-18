package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailmessage
type Voicemailmessage struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Conversation - The conversation that the voicemail message is associated with
	Conversation *Conversation `json:"conversation,omitempty"`


	// Read - Whether the voicemail message is marked as read
	Read *bool `json:"read,omitempty"`


	// AudioRecordingDurationSeconds - The voicemail message's audio recording duration in seconds
	AudioRecordingDurationSeconds *int `json:"audioRecordingDurationSeconds,omitempty"`


	// AudioRecordingSizeBytes - The voicemail message's audio recording size in bytes
	AudioRecordingSizeBytes *int `json:"audioRecordingSizeBytes,omitempty"`


	// CreatedDate - The date the voicemail message was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - The date the voicemail message was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// DeletedDate - The date the voicemail message deleted property was set to true. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DeletedDate *time.Time `json:"deletedDate,omitempty"`


	// CallerAddress - The caller address
	CallerAddress *string `json:"callerAddress,omitempty"`


	// CallerName - Optionally the name of the caller that left the voicemail message if the caller was a known user
	CallerName *string `json:"callerName,omitempty"`


	// CallerUser - Optionally the user that left the voicemail message if the caller was a known user
	CallerUser *User `json:"callerUser,omitempty"`


	// Deleted - Whether the voicemail message has been marked as deleted
	Deleted *bool `json:"deleted,omitempty"`


	// Note - An optional note
	Note *string `json:"note,omitempty"`


	// User - The user that the voicemail message belongs to or null which means the voicemail message belongs to a group or queue
	User *User `json:"user,omitempty"`


	// Group - The group that the voicemail message belongs to or null which means the voicemail message belongs to a user or queue
	Group *Group `json:"group,omitempty"`


	// Queue - The queue that the voicemail message belongs to or null which means the voicemail message belongs to a user or group
	Queue *Queue `json:"queue,omitempty"`


	// CopiedFrom - Represents where this voicemail message was copied from
	CopiedFrom *Voicemailcopyrecord `json:"copiedFrom,omitempty"`


	// CopiedTo - Represents where this voicemail has been copied to
	CopiedTo *[]Voicemailcopyrecord `json:"copiedTo,omitempty"`


	// DeleteRetentionPolicy - The retention policy for this voicemail when deleted is set to true
	DeleteRetentionPolicy *Voicemailretentionpolicy `json:"deleteRetentionPolicy,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Voicemailmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailmessage

	
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
	
	DeletedDate := new(string)
	if u.DeletedDate != nil {
		
		*DeletedDate = timeutil.Strftime(u.DeletedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DeletedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Conversation *Conversation `json:"conversation,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		AudioRecordingDurationSeconds *int `json:"audioRecordingDurationSeconds,omitempty"`
		
		AudioRecordingSizeBytes *int `json:"audioRecordingSizeBytes,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		DeletedDate *string `json:"deletedDate,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		CallerName *string `json:"callerName,omitempty"`
		
		CallerUser *User `json:"callerUser,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		Note *string `json:"note,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		Queue *Queue `json:"queue,omitempty"`
		
		CopiedFrom *Voicemailcopyrecord `json:"copiedFrom,omitempty"`
		
		CopiedTo *[]Voicemailcopyrecord `json:"copiedTo,omitempty"`
		
		DeleteRetentionPolicy *Voicemailretentionpolicy `json:"deleteRetentionPolicy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Conversation: u.Conversation,
		
		Read: u.Read,
		
		AudioRecordingDurationSeconds: u.AudioRecordingDurationSeconds,
		
		AudioRecordingSizeBytes: u.AudioRecordingSizeBytes,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		DeletedDate: DeletedDate,
		
		CallerAddress: u.CallerAddress,
		
		CallerName: u.CallerName,
		
		CallerUser: u.CallerUser,
		
		Deleted: u.Deleted,
		
		Note: u.Note,
		
		User: u.User,
		
		Group: u.Group,
		
		Queue: u.Queue,
		
		CopiedFrom: u.CopiedFrom,
		
		CopiedTo: u.CopiedTo,
		
		DeleteRetentionPolicy: u.DeleteRetentionPolicy,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Voicemailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
