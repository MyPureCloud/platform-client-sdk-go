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

func (o *Batchdownloadjobresult) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		ConversationId: o.ConversationId,
		
		RecordingId: o.RecordingId,
		
		ResultUrl: o.ResultUrl,
		
		ContentType: o.ContentType,
		
		ErrorMsg: o.ErrorMsg,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Batchdownloadjobresult) UnmarshalJSON(b []byte) error {
	var BatchdownloadjobresultMap map[string]interface{}
	err := json.Unmarshal(b, &BatchdownloadjobresultMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BatchdownloadjobresultMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := BatchdownloadjobresultMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ConversationId, ok := BatchdownloadjobresultMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if RecordingId, ok := BatchdownloadjobresultMap["recordingId"].(string); ok {
		o.RecordingId = &RecordingId
	}
	
	if ResultUrl, ok := BatchdownloadjobresultMap["resultUrl"].(string); ok {
		o.ResultUrl = &ResultUrl
	}
	
	if ContentType, ok := BatchdownloadjobresultMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
	
	if ErrorMsg, ok := BatchdownloadjobresultMap["errorMsg"].(string); ok {
		o.ErrorMsg = &ErrorMsg
	}
	
	if SelfUri, ok := BatchdownloadjobresultMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Batchdownloadjobresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
