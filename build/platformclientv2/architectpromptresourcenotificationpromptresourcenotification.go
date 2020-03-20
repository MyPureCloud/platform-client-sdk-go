package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Architectpromptresourcenotificationpromptresourcenotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
