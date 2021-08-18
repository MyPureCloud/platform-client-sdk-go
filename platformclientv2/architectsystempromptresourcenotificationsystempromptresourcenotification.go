package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectsystempromptresourcenotificationsystempromptresourcenotification
type Architectsystempromptresourcenotificationsystempromptresourcenotification struct { 
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

func (u *Architectsystempromptresourcenotificationsystempromptresourcenotification) MarshalJSON() ([]byte, error) {
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
		PromptId: u.PromptId,
		
		Id: u.Id,
		
		Language: u.Language,
		
		MediaUri: u.MediaUri,
		
		UploadStatus: u.UploadStatus,
		
		DurationSeconds: u.DurationSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Architectsystempromptresourcenotificationsystempromptresourcenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
