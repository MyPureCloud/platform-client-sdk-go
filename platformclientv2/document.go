package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// SharingUri
	SharingUri *string `json:"sharingUri,omitempty"`


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


	// DownloadSharingUri
	DownloadSharingUri *string `json:"downloadSharingUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Document) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Document

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateUploaded := new(string)
	if u.DateUploaded != nil {
		
		*DateUploaded = timeutil.Strftime(u.DateUploaded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateUploaded = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ChangeNumber *int `json:"changeNumber,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateUploaded *string `json:"dateUploaded,omitempty"`
		
		ContentUri *string `json:"contentUri,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		UploadedBy *Domainentityref `json:"uploadedBy,omitempty"`
		
		SharingUri *string `json:"sharingUri,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		SystemType *string `json:"systemType,omitempty"`
		
		Filename *string `json:"filename,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		ReceiverAddress *string `json:"receiverAddress,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		TagValues *[]Tagvalue `json:"tagValues,omitempty"`
		
		Attributes *[]Documentattribute `json:"attributes,omitempty"`
		
		Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`
		
		UploadStatus *Domainentityref `json:"uploadStatus,omitempty"`
		
		UploadDestinationUri *string `json:"uploadDestinationUri,omitempty"`
		
		UploadMethod *string `json:"uploadMethod,omitempty"`
		
		LockInfo *Lockinfo `json:"lockInfo,omitempty"`
		
		Acl *[]string `json:"acl,omitempty"`
		
		SharingStatus *string `json:"sharingStatus,omitempty"`
		
		DownloadSharingUri *string `json:"downloadSharingUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ChangeNumber: u.ChangeNumber,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DateUploaded: DateUploaded,
		
		ContentUri: u.ContentUri,
		
		Workspace: u.Workspace,
		
		CreatedBy: u.CreatedBy,
		
		UploadedBy: u.UploadedBy,
		
		SharingUri: u.SharingUri,
		
		ContentType: u.ContentType,
		
		ContentLength: u.ContentLength,
		
		SystemType: u.SystemType,
		
		Filename: u.Filename,
		
		PageCount: u.PageCount,
		
		Read: u.Read,
		
		CallerAddress: u.CallerAddress,
		
		ReceiverAddress: u.ReceiverAddress,
		
		Tags: u.Tags,
		
		TagValues: u.TagValues,
		
		Attributes: u.Attributes,
		
		Thumbnails: u.Thumbnails,
		
		UploadStatus: u.UploadStatus,
		
		UploadDestinationUri: u.UploadDestinationUri,
		
		UploadMethod: u.UploadMethod,
		
		LockInfo: u.LockInfo,
		
		Acl: u.Acl,
		
		SharingStatus: u.SharingStatus,
		
		DownloadSharingUri: u.DownloadSharingUri,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Document) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
