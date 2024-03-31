package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingmetadata
type Recordingmetadata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// MediaSubtype - The recording media subtype.
	MediaSubtype *string `json:"mediaSubtype,omitempty"`

	// MediaSubject - The recording media subject.
	MediaSubject *string `json:"mediaSubject,omitempty"`

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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingmetadata) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Recordingmetadata) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "RestoreExpirationTime","ArchiveDate","DeleteDate","ExportDate","ExportedDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingmetadata
	
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
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		Path *string `json:"path,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		Media *string `json:"media,omitempty"`
		
		MediaSubtype *string `json:"mediaSubtype,omitempty"`
		
		MediaSubject *string `json:"mediaSubject,omitempty"`
		
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
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ConversationId: o.ConversationId,
		
		Path: o.Path,
		
		StartTime: o.StartTime,
		
		EndTime: o.EndTime,
		
		Media: o.Media,
		
		MediaSubtype: o.MediaSubtype,
		
		MediaSubject: o.MediaSubject,
		
		Annotations: o.Annotations,
		
		FileState: o.FileState,
		
		RestoreExpirationTime: RestoreExpirationTime,
		
		ArchiveDate: ArchiveDate,
		
		ArchiveMedium: o.ArchiveMedium,
		
		DeleteDate: DeleteDate,
		
		ExportDate: ExportDate,
		
		ExportedDate: ExportedDate,
		
		MaxAllowedRestorationsForOrg: o.MaxAllowedRestorationsForOrg,
		
		RemainingRestorationsAllowedForOrg: o.RemainingRestorationsAllowedForOrg,
		
		SessionId: o.SessionId,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Recordingmetadata) UnmarshalJSON(b []byte) error {
	var RecordingmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingmetadataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecordingmetadataMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := RecordingmetadataMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ConversationId, ok := RecordingmetadataMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if Path, ok := RecordingmetadataMap["path"].(string); ok {
		o.Path = &Path
	}
    
	if StartTime, ok := RecordingmetadataMap["startTime"].(string); ok {
		o.StartTime = &StartTime
	}
    
	if EndTime, ok := RecordingmetadataMap["endTime"].(string); ok {
		o.EndTime = &EndTime
	}
    
	if Media, ok := RecordingmetadataMap["media"].(string); ok {
		o.Media = &Media
	}
    
	if MediaSubtype, ok := RecordingmetadataMap["mediaSubtype"].(string); ok {
		o.MediaSubtype = &MediaSubtype
	}
    
	if MediaSubject, ok := RecordingmetadataMap["mediaSubject"].(string); ok {
		o.MediaSubject = &MediaSubject
	}
    
	if Annotations, ok := RecordingmetadataMap["annotations"].([]interface{}); ok {
		AnnotationsString, _ := json.Marshal(Annotations)
		json.Unmarshal(AnnotationsString, &o.Annotations)
	}
	
	if FileState, ok := RecordingmetadataMap["fileState"].(string); ok {
		o.FileState = &FileState
	}
    
	if restoreExpirationTimeString, ok := RecordingmetadataMap["restoreExpirationTime"].(string); ok {
		RestoreExpirationTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", restoreExpirationTimeString)
		o.RestoreExpirationTime = &RestoreExpirationTime
	}
	
	if archiveDateString, ok := RecordingmetadataMap["archiveDate"].(string); ok {
		ArchiveDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", archiveDateString)
		o.ArchiveDate = &ArchiveDate
	}
	
	if ArchiveMedium, ok := RecordingmetadataMap["archiveMedium"].(string); ok {
		o.ArchiveMedium = &ArchiveMedium
	}
    
	if deleteDateString, ok := RecordingmetadataMap["deleteDate"].(string); ok {
		DeleteDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", deleteDateString)
		o.DeleteDate = &DeleteDate
	}
	
	if exportDateString, ok := RecordingmetadataMap["exportDate"].(string); ok {
		ExportDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", exportDateString)
		o.ExportDate = &ExportDate
	}
	
	if exportedDateString, ok := RecordingmetadataMap["exportedDate"].(string); ok {
		ExportedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", exportedDateString)
		o.ExportedDate = &ExportedDate
	}
	
	if MaxAllowedRestorationsForOrg, ok := RecordingmetadataMap["maxAllowedRestorationsForOrg"].(float64); ok {
		MaxAllowedRestorationsForOrgInt := int(MaxAllowedRestorationsForOrg)
		o.MaxAllowedRestorationsForOrg = &MaxAllowedRestorationsForOrgInt
	}
	
	if RemainingRestorationsAllowedForOrg, ok := RecordingmetadataMap["remainingRestorationsAllowedForOrg"].(float64); ok {
		RemainingRestorationsAllowedForOrgInt := int(RemainingRestorationsAllowedForOrg)
		o.RemainingRestorationsAllowedForOrg = &RemainingRestorationsAllowedForOrgInt
	}
	
	if SessionId, ok := RecordingmetadataMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if SelfUri, ok := RecordingmetadataMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
