package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Adhocrecordingtopicrecordingdatav2
type Adhocrecordingtopicrecordingdatav2 struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Workspace
	Workspace *Adhocrecordingtopicworkspacedata `json:"workspace,omitempty"`


	// CreatedBy
	CreatedBy *Adhocrecordingtopicuserdata `json:"createdBy,omitempty"`


	// ContentType
	ContentType *string `json:"contentType,omitempty"`


	// ContentLength
	ContentLength *int `json:"contentLength,omitempty"`


	// Filename
	Filename *string `json:"filename,omitempty"`


	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// DateUploaded
	DateUploaded *time.Time `json:"dateUploaded,omitempty"`


	// UploadedBy
	UploadedBy *Adhocrecordingtopicuserdata `json:"uploadedBy,omitempty"`


	// LockInfo
	LockInfo *Adhocrecordingtopiclockdata `json:"lockInfo,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// DurationMillieconds
	DurationMillieconds *int `json:"durationMillieconds,omitempty"`


	// Conversation
	Conversation *Adhocrecordingtopicconversationdata `json:"conversation,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`

}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopicrecordingdatav2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
