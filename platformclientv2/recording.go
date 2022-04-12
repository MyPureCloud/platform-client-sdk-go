package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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


	// OriginalRecordingStartTime - The start time of the full recording, before any segment access restrictions are applied. Null when there is no playable media. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	OriginalRecordingStartTime *time.Time `json:"originalRecordingStartTime,omitempty"`


	// CreationTime - The creation time of the recording. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreationTime *time.Time `json:"creationTime,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Recording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recording
	
	RestoreExpirationTime := new(string)
	if o.RestoreExpirationTime != nil {
		
		*RestoreExpirationTime = timeutil.Strftime(o.RestoreExpirationTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RestoreExpirationTime = nil
	}
	
	ArchiveDate := new(string)
	if o.ArchiveDate != nil {
		
		*ArchiveDate = timeutil.Strftime(o.ArchiveDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ArchiveDate = nil
	}
	
	DeleteDate := new(string)
	if o.DeleteDate != nil {
		
		*DeleteDate = timeutil.Strftime(o.DeleteDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DeleteDate = nil
	}
	
	ExportDate := new(string)
	if o.ExportDate != nil {
		
		*ExportDate = timeutil.Strftime(o.ExportDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExportDate = nil
	}
	
	ExportedDate := new(string)
	if o.ExportedDate != nil {
		
		*ExportedDate = timeutil.Strftime(o.ExportedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExportedDate = nil
	}
	
	OriginalRecordingStartTime := new(string)
	if o.OriginalRecordingStartTime != nil {
		
		*OriginalRecordingStartTime = timeutil.Strftime(o.OriginalRecordingStartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		OriginalRecordingStartTime = nil
	}
	
	CreationTime := new(string)
	if o.CreationTime != nil {
		
		*CreationTime = timeutil.Strftime(o.CreationTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreationTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		Path *string `json:"path,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		Media *string `json:"media,omitempty"`
		
		Annotations *[]Annotation `json:"annotations,omitempty"`
		
		Transcript *[]Chatmessage `json:"transcript,omitempty"`
		
		EmailTranscript *[]Recordingemailmessage `json:"emailTranscript,omitempty"`
		
		MessagingTranscript *[]Recordingmessagingmessage `json:"messagingTranscript,omitempty"`
		
		FileState *string `json:"fileState,omitempty"`
		
		RestoreExpirationTime *string `json:"restoreExpirationTime,omitempty"`
		
		MediaUris *map[string]Mediaresult `json:"mediaUris,omitempty"`
		
		EstimatedTranscodeTimeMs *int `json:"estimatedTranscodeTimeMs,omitempty"`
		
		ActualTranscodeTimeMs *int `json:"actualTranscodeTimeMs,omitempty"`
		
		ArchiveDate *string `json:"archiveDate,omitempty"`
		
		ArchiveMedium *string `json:"archiveMedium,omitempty"`
		
		DeleteDate *string `json:"deleteDate,omitempty"`
		
		ExportDate *string `json:"exportDate,omitempty"`
		
		ExportedDate *string `json:"exportedDate,omitempty"`
		
		OutputDurationMs *int `json:"outputDurationMs,omitempty"`
		
		OutputSizeInBytes *int `json:"outputSizeInBytes,omitempty"`
		
		MaxAllowedRestorationsForOrg *int `json:"maxAllowedRestorationsForOrg,omitempty"`
		
		RemainingRestorationsAllowedForOrg *int `json:"remainingRestorationsAllowedForOrg,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		Users *[]User `json:"users,omitempty"`
		
		RecordingFileRole *string `json:"recordingFileRole,omitempty"`
		
		RecordingErrorStatus *string `json:"recordingErrorStatus,omitempty"`
		
		OriginalRecordingStartTime *string `json:"originalRecordingStartTime,omitempty"`
		
		CreationTime *string `json:"creationTime,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ConversationId: o.ConversationId,
		
		Path: o.Path,
		
		StartTime: o.StartTime,
		
		EndTime: o.EndTime,
		
		Media: o.Media,
		
		Annotations: o.Annotations,
		
		Transcript: o.Transcript,
		
		EmailTranscript: o.EmailTranscript,
		
		MessagingTranscript: o.MessagingTranscript,
		
		FileState: o.FileState,
		
		RestoreExpirationTime: RestoreExpirationTime,
		
		MediaUris: o.MediaUris,
		
		EstimatedTranscodeTimeMs: o.EstimatedTranscodeTimeMs,
		
		ActualTranscodeTimeMs: o.ActualTranscodeTimeMs,
		
		ArchiveDate: ArchiveDate,
		
		ArchiveMedium: o.ArchiveMedium,
		
		DeleteDate: DeleteDate,
		
		ExportDate: ExportDate,
		
		ExportedDate: ExportedDate,
		
		OutputDurationMs: o.OutputDurationMs,
		
		OutputSizeInBytes: o.OutputSizeInBytes,
		
		MaxAllowedRestorationsForOrg: o.MaxAllowedRestorationsForOrg,
		
		RemainingRestorationsAllowedForOrg: o.RemainingRestorationsAllowedForOrg,
		
		SessionId: o.SessionId,
		
		Users: o.Users,
		
		RecordingFileRole: o.RecordingFileRole,
		
		RecordingErrorStatus: o.RecordingErrorStatus,
		
		OriginalRecordingStartTime: OriginalRecordingStartTime,
		
		CreationTime: CreationTime,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Recording) UnmarshalJSON(b []byte) error {
	var RecordingMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecordingMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := RecordingMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ConversationId, ok := RecordingMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if Path, ok := RecordingMap["path"].(string); ok {
		o.Path = &Path
	}
	
	if StartTime, ok := RecordingMap["startTime"].(string); ok {
		o.StartTime = &StartTime
	}
	
	if EndTime, ok := RecordingMap["endTime"].(string); ok {
		o.EndTime = &EndTime
	}
	
	if Media, ok := RecordingMap["media"].(string); ok {
		o.Media = &Media
	}
	
	if Annotations, ok := RecordingMap["annotations"].([]interface{}); ok {
		AnnotationsString, _ := json.Marshal(Annotations)
		json.Unmarshal(AnnotationsString, &o.Annotations)
	}
	
	if Transcript, ok := RecordingMap["transcript"].([]interface{}); ok {
		TranscriptString, _ := json.Marshal(Transcript)
		json.Unmarshal(TranscriptString, &o.Transcript)
	}
	
	if EmailTranscript, ok := RecordingMap["emailTranscript"].([]interface{}); ok {
		EmailTranscriptString, _ := json.Marshal(EmailTranscript)
		json.Unmarshal(EmailTranscriptString, &o.EmailTranscript)
	}
	
	if MessagingTranscript, ok := RecordingMap["messagingTranscript"].([]interface{}); ok {
		MessagingTranscriptString, _ := json.Marshal(MessagingTranscript)
		json.Unmarshal(MessagingTranscriptString, &o.MessagingTranscript)
	}
	
	if FileState, ok := RecordingMap["fileState"].(string); ok {
		o.FileState = &FileState
	}
	
	if restoreExpirationTimeString, ok := RecordingMap["restoreExpirationTime"].(string); ok {
		RestoreExpirationTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", restoreExpirationTimeString)
		o.RestoreExpirationTime = &RestoreExpirationTime
	}
	
	if MediaUris, ok := RecordingMap["mediaUris"].(map[string]interface{}); ok {
		MediaUrisString, _ := json.Marshal(MediaUris)
		json.Unmarshal(MediaUrisString, &o.MediaUris)
	}
	
	if EstimatedTranscodeTimeMs, ok := RecordingMap["estimatedTranscodeTimeMs"].(float64); ok {
		EstimatedTranscodeTimeMsInt := int(EstimatedTranscodeTimeMs)
		o.EstimatedTranscodeTimeMs = &EstimatedTranscodeTimeMsInt
	}
	
	if ActualTranscodeTimeMs, ok := RecordingMap["actualTranscodeTimeMs"].(float64); ok {
		ActualTranscodeTimeMsInt := int(ActualTranscodeTimeMs)
		o.ActualTranscodeTimeMs = &ActualTranscodeTimeMsInt
	}
	
	if archiveDateString, ok := RecordingMap["archiveDate"].(string); ok {
		ArchiveDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", archiveDateString)
		o.ArchiveDate = &ArchiveDate
	}
	
	if ArchiveMedium, ok := RecordingMap["archiveMedium"].(string); ok {
		o.ArchiveMedium = &ArchiveMedium
	}
	
	if deleteDateString, ok := RecordingMap["deleteDate"].(string); ok {
		DeleteDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", deleteDateString)
		o.DeleteDate = &DeleteDate
	}
	
	if exportDateString, ok := RecordingMap["exportDate"].(string); ok {
		ExportDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", exportDateString)
		o.ExportDate = &ExportDate
	}
	
	if exportedDateString, ok := RecordingMap["exportedDate"].(string); ok {
		ExportedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", exportedDateString)
		o.ExportedDate = &ExportedDate
	}
	
	if OutputDurationMs, ok := RecordingMap["outputDurationMs"].(float64); ok {
		OutputDurationMsInt := int(OutputDurationMs)
		o.OutputDurationMs = &OutputDurationMsInt
	}
	
	if OutputSizeInBytes, ok := RecordingMap["outputSizeInBytes"].(float64); ok {
		OutputSizeInBytesInt := int(OutputSizeInBytes)
		o.OutputSizeInBytes = &OutputSizeInBytesInt
	}
	
	if MaxAllowedRestorationsForOrg, ok := RecordingMap["maxAllowedRestorationsForOrg"].(float64); ok {
		MaxAllowedRestorationsForOrgInt := int(MaxAllowedRestorationsForOrg)
		o.MaxAllowedRestorationsForOrg = &MaxAllowedRestorationsForOrgInt
	}
	
	if RemainingRestorationsAllowedForOrg, ok := RecordingMap["remainingRestorationsAllowedForOrg"].(float64); ok {
		RemainingRestorationsAllowedForOrgInt := int(RemainingRestorationsAllowedForOrg)
		o.RemainingRestorationsAllowedForOrg = &RemainingRestorationsAllowedForOrgInt
	}
	
	if SessionId, ok := RecordingMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if Users, ok := RecordingMap["users"].([]interface{}); ok {
		UsersString, _ := json.Marshal(Users)
		json.Unmarshal(UsersString, &o.Users)
	}
	
	if RecordingFileRole, ok := RecordingMap["recordingFileRole"].(string); ok {
		o.RecordingFileRole = &RecordingFileRole
	}
	
	if RecordingErrorStatus, ok := RecordingMap["recordingErrorStatus"].(string); ok {
		o.RecordingErrorStatus = &RecordingErrorStatus
	}
	
	if originalRecordingStartTimeString, ok := RecordingMap["originalRecordingStartTime"].(string); ok {
		OriginalRecordingStartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", originalRecordingStartTimeString)
		o.OriginalRecordingStartTime = &OriginalRecordingStartTime
	}
	
	if creationTimeString, ok := RecordingMap["creationTime"].(string); ok {
		CreationTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", creationTimeString)
		o.CreationTime = &CreationTime
	}
	
	if SelfUri, ok := RecordingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
