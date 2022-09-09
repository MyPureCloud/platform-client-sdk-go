package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingretention
type Recordingretention struct { 
	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// RecordingId
	RecordingId *string `json:"recordingId,omitempty"`


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


	// CreationTime - The creation time of the recording. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreationTime *time.Time `json:"creationTime,omitempty"`

}

func (o *Recordingretention) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingretention
	
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
	
	CreationTime := new(string)
	if o.CreationTime != nil {
		
		*CreationTime = timeutil.Strftime(o.CreationTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreationTime = nil
	}
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		RecordingId *string `json:"recordingId,omitempty"`
		
		ArchiveDate *string `json:"archiveDate,omitempty"`
		
		ArchiveMedium *string `json:"archiveMedium,omitempty"`
		
		DeleteDate *string `json:"deleteDate,omitempty"`
		
		ExportDate *string `json:"exportDate,omitempty"`
		
		ExportedDate *string `json:"exportedDate,omitempty"`
		
		CreationTime *string `json:"creationTime,omitempty"`
		*Alias
	}{ 
		ConversationId: o.ConversationId,
		
		RecordingId: o.RecordingId,
		
		ArchiveDate: ArchiveDate,
		
		ArchiveMedium: o.ArchiveMedium,
		
		DeleteDate: DeleteDate,
		
		ExportDate: ExportDate,
		
		ExportedDate: ExportedDate,
		
		CreationTime: CreationTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingretention) UnmarshalJSON(b []byte) error {
	var RecordingretentionMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingretentionMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := RecordingretentionMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if RecordingId, ok := RecordingretentionMap["recordingId"].(string); ok {
		o.RecordingId = &RecordingId
	}
    
	if archiveDateString, ok := RecordingretentionMap["archiveDate"].(string); ok {
		ArchiveDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", archiveDateString)
		o.ArchiveDate = &ArchiveDate
	}
	
	if ArchiveMedium, ok := RecordingretentionMap["archiveMedium"].(string); ok {
		o.ArchiveMedium = &ArchiveMedium
	}
    
	if deleteDateString, ok := RecordingretentionMap["deleteDate"].(string); ok {
		DeleteDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", deleteDateString)
		o.DeleteDate = &DeleteDate
	}
	
	if exportDateString, ok := RecordingretentionMap["exportDate"].(string); ok {
		ExportDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", exportDateString)
		o.ExportDate = &ExportDate
	}
	
	if exportedDateString, ok := RecordingretentionMap["exportedDate"].(string); ok {
		ExportedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", exportedDateString)
		o.ExportedDate = &ExportedDate
	}
	
	if creationTimeString, ok := RecordingretentionMap["creationTime"].(string); ok {
		CreationTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", creationTimeString)
		o.CreationTime = &CreationTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingretention) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
