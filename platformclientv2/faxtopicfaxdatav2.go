package platformclientv2
import (
	"time"
	"encoding/json"
)

// Faxtopicfaxdatav2
type Faxtopicfaxdatav2 struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Workspace
	Workspace *Faxtopicworkspacedata `json:"workspace,omitempty"`


	// CreatedBy
	CreatedBy *Faxtopicuserdata `json:"createdBy,omitempty"`


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
	UploadedBy *Faxtopicuserdata `json:"uploadedBy,omitempty"`


	// LockInfo
	LockInfo *Faxtopiclockdata `json:"lockInfo,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// CallerAddress
	CallerAddress *string `json:"callerAddress,omitempty"`


	// ReceiverAddress
	ReceiverAddress *string `json:"receiverAddress,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`

}

// String returns a JSON representation of the model
func (o *Faxtopicfaxdatav2) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
