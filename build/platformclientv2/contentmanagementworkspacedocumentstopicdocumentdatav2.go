package platformclientv2
import (
	"time"
	"encoding/json"
)

// Contentmanagementworkspacedocumentstopicdocumentdatav2
type Contentmanagementworkspacedocumentstopicdocumentdatav2 struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Workspace
	Workspace *Contentmanagementworkspacedocumentstopicworkspacedata `json:"workspace,omitempty"`


	// CreatedBy
	CreatedBy *Contentmanagementworkspacedocumentstopicuserdata `json:"createdBy,omitempty"`


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
	UploadedBy *Contentmanagementworkspacedocumentstopicuserdata `json:"uploadedBy,omitempty"`


	// LockInfo
	LockInfo *Contentmanagementworkspacedocumentstopiclockdata `json:"lockInfo,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentmanagementworkspacedocumentstopicdocumentdatav2) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
