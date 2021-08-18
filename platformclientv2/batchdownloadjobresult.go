package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Batchdownloadjobresult
type Batchdownloadjobresult struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ConversationId - Conversation id of the result
	ConversationId *string `json:"conversationId,omitempty"`


	// RecordingId - Recording id of the result
	RecordingId *string `json:"recordingId,omitempty"`


	// ResultUrl - URL of results... HTTP GET from this location to download results for this item
	ResultUrl *string `json:"resultUrl,omitempty"`


	// ContentType - Content type of this result
	ContentType *string `json:"contentType,omitempty"`


	// ErrorMsg - An error message, in case of failed processing will indicate the cause of the failure
	ErrorMsg *string `json:"errorMsg,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Batchdownloadjobresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Batchdownloadjobresult

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		RecordingId *string `json:"recordingId,omitempty"`
		
		ResultUrl *string `json:"resultUrl,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ErrorMsg *string `json:"errorMsg,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ConversationId: u.ConversationId,
		
		RecordingId: u.RecordingId,
		
		ResultUrl: u.ResultUrl,
		
		ContentType: u.ContentType,
		
		ErrorMsg: u.ErrorMsg,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Batchdownloadjobresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
