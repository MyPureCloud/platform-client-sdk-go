package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectpromptresourcenotificationpromptresourcenotification
type Architectpromptresourcenotificationpromptresourcenotification struct { 
	// PromptId - Id of the prompt that this notification is for.
	PromptId *string `json:"promptId,omitempty"`


	// Id - Id of the prompt resource that this notification is for.
	Id *string `json:"id,omitempty"`


	// Language - Language resource that this notification is for.
	Language *string `json:"language,omitempty"`


	// MediaUri - Uri to the file for this prompt resource.
	MediaUri *string `json:"mediaUri,omitempty"`


	// UploadStatus - Current upload status of the prompt resource (created, uploaded, transcoded, transcodeFailed).
	UploadStatus *string `json:"uploadStatus,omitempty"`


	// DurationSeconds - Duration (in seconds) for the transcoded audio file.
	DurationSeconds *float32 `json:"durationSeconds,omitempty"`

}

func (o *Architectpromptresourcenotificationpromptresourcenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectpromptresourcenotificationpromptresourcenotification
	
	return json.Marshal(&struct { 
		PromptId *string `json:"promptId,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		MediaUri *string `json:"mediaUri,omitempty"`
		
		UploadStatus *string `json:"uploadStatus,omitempty"`
		
		DurationSeconds *float32 `json:"durationSeconds,omitempty"`
		*Alias
	}{ 
		PromptId: o.PromptId,
		
		Id: o.Id,
		
		Language: o.Language,
		
		MediaUri: o.MediaUri,
		
		UploadStatus: o.UploadStatus,
		
		DurationSeconds: o.DurationSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectpromptresourcenotificationpromptresourcenotification) UnmarshalJSON(b []byte) error {
	var ArchitectpromptresourcenotificationpromptresourcenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectpromptresourcenotificationpromptresourcenotificationMap)
	if err != nil {
		return err
	}
	
	if PromptId, ok := ArchitectpromptresourcenotificationpromptresourcenotificationMap["promptId"].(string); ok {
		o.PromptId = &PromptId
	}
    
	if Id, ok := ArchitectpromptresourcenotificationpromptresourcenotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Language, ok := ArchitectpromptresourcenotificationpromptresourcenotificationMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if MediaUri, ok := ArchitectpromptresourcenotificationpromptresourcenotificationMap["mediaUri"].(string); ok {
		o.MediaUri = &MediaUri
	}
    
	if UploadStatus, ok := ArchitectpromptresourcenotificationpromptresourcenotificationMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
    
	if DurationSeconds, ok := ArchitectpromptresourcenotificationpromptresourcenotificationMap["durationSeconds"].(float64); ok {
		DurationSecondsFloat32 := float32(DurationSeconds)
		o.DurationSeconds = &DurationSecondsFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Architectpromptresourcenotificationpromptresourcenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
