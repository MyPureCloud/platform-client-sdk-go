package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Recording
type Recording struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// Path
	Path *string `json:"path,omitempty"`


	// StartTime - The start time of the recording. Null when there is no playable media.
	StartTime *string `json:"startTime,omitempty"`


	// EndTime - The end time of the recording. Null when there is no playable media.
	EndTime *string `json:"endTime,omitempty"`


	// Media - The type of media that the recording is. At the moment that could be audio, chat, or email.
	Media *string `json:"media,omitempty"`


	// Annotations - Annotations that belong to the recording.
	Annotations *[]Annotation `json:"annotations,omitempty"`


	// Transcript - Represents a chat transcript
	Transcript *[]Chatmessage `json:"transcript,omitempty"`


	// EmailTranscript - Represents an email transcript
	EmailTranscript *[]Recordingemailmessage `json:"emailTranscript,omitempty"`


	// MessagingTranscript - Represents a messaging transcript
	MessagingTranscript *[]Recordingmessagingmessage `json:"messagingTranscript,omitempty"`


	// FileState - Represents the current file state for a recording. Examples: Uploading, Archived, etc
	FileState *string `json:"fileState,omitempty"`


	// RestoreExpirationTime - The amount of time a restored recording will remain restored before being archived again. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RestoreExpirationTime *time.Time `json:"restoreExpirationTime,omitempty"`


	// MediaUris - The different mediaUris for the recording. Null when there is no playable media.
	MediaUris *map[string]Mediaresult `json:"mediaUris,omitempty"`


	// EstimatedTranscodeTimeMs
	EstimatedTranscodeTimeMs *int `json:"estimatedTranscodeTimeMs,omitempty"`


	// ActualTranscodeTimeMs
	ActualTranscodeTimeMs *int `json:"actualTranscodeTimeMs,omitempty"`


	// ArchiveDate - The date the recording will be archived. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ArchiveDate *time.Time `json:"archiveDate,omitempty"`


	// ArchiveMedium - The type of archive medium used. Example: CloudArchive
	ArchiveMedium *string `json:"archiveMedium,omitempty"`


	// DeleteDate - The date the recording will be deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DeleteDate *time.Time `json:"deleteDate,omitempty"`


	// ExportDate - The date the recording will be exported. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExportDate *time.Time `json:"exportDate,omitempty"`


	// ExportedDate - The date the recording was exported. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExportedDate *time.Time `json:"exportedDate,omitempty"`


	// OutputDurationMs - Duration of transcoded media in milliseconds
	OutputDurationMs *int `json:"outputDurationMs,omitempty"`


	// OutputSizeInBytes - Size of transcoded media in bytes. 0 if there is no transcoded media.
	OutputSizeInBytes *int `json:"outputSizeInBytes,omitempty"`


	// MaxAllowedRestorationsForOrg - How many archive restorations the organization is allowed to have.
	MaxAllowedRestorationsForOrg *int `json:"maxAllowedRestorationsForOrg,omitempty"`


	// RemainingRestorationsAllowedForOrg - The remaining archive restorations the organization has.
	RemainingRestorationsAllowedForOrg *int `json:"remainingRestorationsAllowedForOrg,omitempty"`


	// SessionId - The session id represents an external resource id, such as email, call, chat, etc
	SessionId *string `json:"sessionId,omitempty"`


	// Users - The users participating in the conversation
	Users *[]User `json:"users,omitempty"`


	// RecordingFileRole - Role of the file recording. It can be either customer_experience or adhoc.
	RecordingFileRole *string `json:"recordingFileRole,omitempty"`


	// RecordingErrorStatus - Status of a recording that cannot be returned because of an error
	RecordingErrorStatus *string `json:"recordingErrorStatus,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
