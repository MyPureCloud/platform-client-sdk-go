package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectpromptresourcenotificationpromptresourcenotification
type Architectpromptresourcenotificationpromptresourcenotification struct { 
	// PromptId
	PromptId *string `json:"promptId,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Language
	Language *string `json:"language,omitempty"`


	// MediaUri
	MediaUri *string `json:"mediaUri,omitempty"`


	// UploadStatus
	UploadStatus *string `json:"uploadStatus,omitempty"`


	// DurationSeconds
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
