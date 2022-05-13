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

func (o *Faxtopicfaxdatav2) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Faxtopicfaxdatav2
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateUploaded := new(string)
	if o.DateUploaded != nil {
		
		*DateUploaded = timeutil.Strftime(o.DateUploaded, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Workspace: o.Workspace,
		
		CreatedBy: o.CreatedBy,
		
		ContentType: o.ContentType,
		
		ContentLength: o.ContentLength,
		
		Filename: o.Filename,
		
		ChangeNumber: o.ChangeNumber,
		
		DateUploaded: DateUploaded,
		
		UploadedBy: o.UploadedBy,
		
		LockInfo: o.LockInfo,
		
		SelfUri: o.SelfUri,
		
		CallerAddress: o.CallerAddress,
		
		ReceiverAddress: o.ReceiverAddress,
		
		Read: o.Read,
		Alias:    (*Alias)(o),
	})
}

func (o *Faxtopicfaxdatav2) UnmarshalJSON(b []byte) error {
	var Faxtopicfaxdatav2Map map[string]interface{}
	err := json.Unmarshal(b, &Faxtopicfaxdatav2Map)
	if err != nil {
		return err
	}
	
	if Id, ok := Faxtopicfaxdatav2Map["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := Faxtopicfaxdatav2Map["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := Faxtopicfaxdatav2Map["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := Faxtopicfaxdatav2Map["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Workspace, ok := Faxtopicfaxdatav2Map["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if CreatedBy, ok := Faxtopicfaxdatav2Map["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ContentType, ok := Faxtopicfaxdatav2Map["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ContentLength, ok := Faxtopicfaxdatav2Map["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if Filename, ok := Faxtopicfaxdatav2Map["filename"].(string); ok {
		o.Filename = &Filename
	}
    
	if ChangeNumber, ok := Faxtopicfaxdatav2Map["changeNumber"].(float64); ok {
		ChangeNumberInt := int(ChangeNumber)
		o.ChangeNumber = &ChangeNumberInt
	}
	
	if dateUploadedString, ok := Faxtopicfaxdatav2Map["dateUploaded"].(string); ok {
		DateUploaded, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateUploadedString)
		o.DateUploaded = &DateUploaded
	}
	
	if UploadedBy, ok := Faxtopicfaxdatav2Map["uploadedBy"].(map[string]interface{}); ok {
		UploadedByString, _ := json.Marshal(UploadedBy)
		json.Unmarshal(UploadedByString, &o.UploadedBy)
	}
	
	if LockInfo, ok := Faxtopicfaxdatav2Map["lockInfo"].(map[string]interface{}); ok {
		LockInfoString, _ := json.Marshal(LockInfo)
		json.Unmarshal(LockInfoString, &o.LockInfo)
	}
	
	if SelfUri, ok := Faxtopicfaxdatav2Map["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if CallerAddress, ok := Faxtopicfaxdatav2Map["callerAddress"].(string); ok {
		o.CallerAddress = &CallerAddress
	}
    
	if ReceiverAddress, ok := Faxtopicfaxdatav2Map["receiverAddress"].(string); ok {
		o.ReceiverAddress = &ReceiverAddress
	}
    
	if Read, ok := Faxtopicfaxdatav2Map["read"].(bool); ok {
		o.Read = &Read
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Faxtopicfaxdatav2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
