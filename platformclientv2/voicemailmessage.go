package platformclientv2
import (
	"time"
	"encoding/json"
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
	AudioRecordingDurationSeconds *int32 `json:"audioRecordingDurationSeconds,omitempty"`


	// AudioRecordingSizeBytes - The voicemail message's audio recording size in bytes
	AudioRecordingSizeBytes *int64 `json:"audioRecordingSizeBytes,omitempty"`


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

// String returns a JSON representation of the model
func (o *Voicemailmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
