package platformclientv2
import (
	"time"
	"encoding/json"
)

// Document
type Document struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// DateUploaded - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateUploaded *time.Time `json:"dateUploaded,omitempty"`


	// ContentUri
	ContentUri *string `json:"contentUri,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// CreatedBy
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// UploadedBy
	UploadedBy *Domainentityref `json:"uploadedBy,omitempty"`


	// ContentType
	ContentType *string `json:"contentType,omitempty"`


	// ContentLength
	ContentLength *int `json:"contentLength,omitempty"`


	// SystemType
	SystemType *string `json:"systemType,omitempty"`


	// Filename
	Filename *string `json:"filename,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// CallerAddress
	CallerAddress *string `json:"callerAddress,omitempty"`


	// ReceiverAddress
	ReceiverAddress *string `json:"receiverAddress,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// TagValues
	TagValues *[]Tagvalue `json:"tagValues,omitempty"`


	// Attributes
	Attributes *[]Documentattribute `json:"attributes,omitempty"`


	// Thumbnails
	Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`


	// UploadStatus
	UploadStatus *Domainentityref `json:"uploadStatus,omitempty"`


	// UploadDestinationUri
	UploadDestinationUri *string `json:"uploadDestinationUri,omitempty"`


	// UploadMethod
	UploadMethod *string `json:"uploadMethod,omitempty"`


	// LockInfo
	LockInfo *Lockinfo `json:"lockInfo,omitempty"`


	// Acl - A list of permitted action rights for the user making the request
	Acl *[]string `json:"acl,omitempty"`


	// SharingStatus
	SharingStatus *string `json:"sharingStatus,omitempty"`


	// SharingUri
	SharingUri *string `json:"sharingUri,omitempty"`


	// DownloadSharingUri
	DownloadSharingUri *string `json:"downloadSharingUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Document) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
