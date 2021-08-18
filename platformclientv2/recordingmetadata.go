package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingmetadata
type Recordingmetadata struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// Path
	Path *string `json:"path,omitempty"`


	// StartTime - The start time of the recording for screen recordings. Null for other types.
	StartTime *string `json:"startTime,omitempty"`


	// EndTime
	EndTime *string `json:"endTime,omitempty"`


	// Media - The type of media that the recording is. At the moment that could be audio, chat, email, or message.
	Media *string `json:"media,omitempty"`


	// Annotations - Annotations that belong to the recording. Populated when recording filestate is AVAILABLE.
	Annotations *[]Annotation `json:"annotations,omitempty"`


	// FileState - Represents the current file state for a recording. Examples: Uploading, Archived, etc
	FileState *string `json:"fileState,omitempty"`


	// RestoreExpirationTime - The amount of time a restored recording will remain restored before being archived again. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RestoreExpirationTime *time.Time `json:"restoreExpirationTime,omitempty"`


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


	// MaxAllowedRestorationsForOrg - How many archive restorations the organization is allowed to have.
	MaxAllowedRestorationsForOrg *int `json:"maxAllowedRestorationsForOrg,omitempty"`


	// RemainingRestorationsAllowedForOrg - The remaining archive restorations the organization has.
	RemainingRestorationsAllowedForOrg *int `json:"remainingRestorationsAllowedForOrg,omitempty"`


	// SessionId - The session id represents an external resource id, such as email, call, chat, etc
	SessionId *string `json:"sessionId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Recordingmetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingmetadata

	
	RestoreExpirationTime := new(string)
	if u.RestoreExpirationTime != nil {
		
		*RestoreExpirationTime = timeutil.Strftime(u.RestoreExpirationTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RestoreExpirationTime = nil
	}
	
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
	
	ExportDate := new(string)
	if u.ExportDate != nil {
		
		*ExportDate = timeutil.Strftime(u.ExportDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExportDate = nil
	}
	
	ExportedDate := new(string)
	if u.ExportedDate != nil {
		
		*ExportedDate = timeutil.Strftime(u.ExportedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExportedDate = nil
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
		
		FileState *string `json:"fileState,omitempty"`
		
		RestoreExpirationTime *string `json:"restoreExpirationTime,omitempty"`
		
		ArchiveDate *string `json:"archiveDate,omitempty"`
		
		ArchiveMedium *string `json:"archiveMedium,omitempty"`
		
		DeleteDate *string `json:"deleteDate,omitempty"`
		
		ExportDate *string `json:"exportDate,omitempty"`
		
		ExportedDate *string `json:"exportedDate,omitempty"`
		
		MaxAllowedRestorationsForOrg *int `json:"maxAllowedRestorationsForOrg,omitempty"`
		
		RemainingRestorationsAllowedForOrg *int `json:"remainingRestorationsAllowedForOrg,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ConversationId: u.ConversationId,
		
		Path: u.Path,
		
		StartTime: u.StartTime,
		
		EndTime: u.EndTime,
		
		Media: u.Media,
		
		Annotations: u.Annotations,
		
		FileState: u.FileState,
		
		RestoreExpirationTime: RestoreExpirationTime,
		
		ArchiveDate: ArchiveDate,
		
		ArchiveMedium: u.ArchiveMedium,
		
		DeleteDate: DeleteDate,
		
		ExportDate: ExportDate,
		
		ExportedDate: ExportedDate,
		
		MaxAllowedRestorationsForOrg: u.MaxAllowedRestorationsForOrg,
		
		RemainingRestorationsAllowedForOrg: u.RemainingRestorationsAllowedForOrg,
		
		SessionId: u.SessionId,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Recordingmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
