package platformclientv2
import (
	"time"
	"encoding/json"
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
	ContentLength *int32 `json:"contentLength,omitempty"`


	// Filename
	Filename *string `json:"filename,omitempty"`


	// ChangeNumber
	ChangeNumber *int32 `json:"changeNumber,omitempty"`


	// DateUploaded
	DateUploaded *time.Time `json:"dateUploaded,omitempty"`


	// UploadedBy
	UploadedBy *Adhocrecordingtopicuserdata `json:"uploadedBy,omitempty"`


	// LockInfo
	LockInfo *Adhocrecordingtopiclockdata `json:"lockInfo,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// DurationMillieconds
	DurationMillieconds *int32 `json:"durationMillieconds,omitempty"`


	// Conversation
	Conversation *Adhocrecordingtopicconversationdata `json:"conversation,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`

}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopicrecordingdatav2) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
