package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectsystempromptresourcenotificationsystempromptresourcenotification
type Architectsystempromptresourcenotificationsystempromptresourcenotification struct { 
	// PromptId - Id of the prompt that this notification is for.
	PromptId *string `json:"promptId,omitempty"`


	// Id - Id of the prompt resource that this notification is for.
	Id *string `json:"id,omitempty"`


	// Language - Language resource that this notification is for.
	Language *string `json:"language,omitempty"`


	// MediaUri - Uri to the file for this system prompt resource.
	MediaUri *string `json:"mediaUri,omitempty"`


	// UploadStatus - Current upload status of the prompt resource (created, uploaded, transcoded, transcodeFailed).
	UploadStatus *string `json:"uploadStatus,omitempty"`


	// DurationSeconds - Duration (in seconds) for the transcoded audio file.
	DurationSeconds *float32 `json:"durationSeconds,omitempty"`

}

func (o *Architectsystempromptresourcenotificationsystempromptresourcenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectsystempromptresourcenotificationsystempromptresourcenotification
	
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

func (o *Architectsystempromptresourcenotificationsystempromptresourcenotification) UnmarshalJSON(b []byte) error {
	var ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap)
	if err != nil {
		return err
	}
	
	if PromptId, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["promptId"].(string); ok {
		o.PromptId = &PromptId
	}
	
	if Id, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Language, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["language"].(string); ok {
		o.Language = &Language
	}
	
	if MediaUri, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["mediaUri"].(string); ok {
		o.MediaUri = &MediaUri
	}
	
	if UploadStatus, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
	
	if DurationSeconds, ok := ArchitectsystempromptresourcenotificationsystempromptresourcenotificationMap["durationSeconds"].(float64); ok {
		DurationSecondsFloat32 := float32(DurationSeconds)
		o.DurationSeconds = &DurationSecondsFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectsystempromptresourcenotificationsystempromptresourcenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
