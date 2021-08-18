package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Faxtopicfaxdatav2) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Faxtopicfaxdatav2

	
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
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Workspace *Faxtopicworkspacedata `json:"workspace,omitempty"`
		
		CreatedBy *Faxtopicuserdata `json:"createdBy,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		Filename *string `json:"filename,omitempty"`
		
		ChangeNumber *int `json:"changeNumber,omitempty"`
		
		DateUploaded *string `json:"dateUploaded,omitempty"`
		
		UploadedBy *Faxtopicuserdata `json:"uploadedBy,omitempty"`
		
		LockInfo *Faxtopiclockdata `json:"lockInfo,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		ReceiverAddress *string `json:"receiverAddress,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Workspace: u.Workspace,
		
		CreatedBy: u.CreatedBy,
		
		ContentType: u.ContentType,
		
		ContentLength: u.ContentLength,
		
		Filename: u.Filename,
		
		ChangeNumber: u.ChangeNumber,
		
		DateUploaded: DateUploaded,
		
		UploadedBy: u.UploadedBy,
		
		LockInfo: u.LockInfo,
		
		SelfUri: u.SelfUri,
		
		CallerAddress: u.CallerAddress,
		
		ReceiverAddress: u.ReceiverAddress,
		
		Read: u.Read,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Faxtopicfaxdatav2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
