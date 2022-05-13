package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	ContentLength *int `json:"contentLength,omitempty"`


	// Filename
	Filename *string `json:"filename,omitempty"`


	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// DateUploaded
	DateUploaded *time.Time `json:"dateUploaded,omitempty"`


	// UploadedBy
	UploadedBy *Contentmanagementworkspacedocumentstopicuserdata `json:"uploadedBy,omitempty"`


	// LockInfo
	LockInfo *Contentmanagementworkspacedocumentstopiclockdata `json:"lockInfo,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Contentmanagementworkspacedocumentstopicdocumentdatav2) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentmanagementworkspacedocumentstopicdocumentdatav2
	
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
		
		Workspace *Contentmanagementworkspacedocumentstopicworkspacedata `json:"workspace,omitempty"`
		
		CreatedBy *Contentmanagementworkspacedocumentstopicuserdata `json:"createdBy,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		Filename *string `json:"filename,omitempty"`
		
		ChangeNumber *int `json:"changeNumber,omitempty"`
		
		DateUploaded *string `json:"dateUploaded,omitempty"`
		
		UploadedBy *Contentmanagementworkspacedocumentstopicuserdata `json:"uploadedBy,omitempty"`
		
		LockInfo *Contentmanagementworkspacedocumentstopiclockdata `json:"lockInfo,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
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
		Alias:    (*Alias)(o),
	})
}

func (o *Contentmanagementworkspacedocumentstopicdocumentdatav2) UnmarshalJSON(b []byte) error {
	var Contentmanagementworkspacedocumentstopicdocumentdatav2Map map[string]interface{}
	err := json.Unmarshal(b, &Contentmanagementworkspacedocumentstopicdocumentdatav2Map)
	if err != nil {
		return err
	}
	
	if Id, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Workspace, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if CreatedBy, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ContentType, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ContentLength, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if Filename, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["filename"].(string); ok {
		o.Filename = &Filename
	}
    
	if ChangeNumber, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["changeNumber"].(float64); ok {
		ChangeNumberInt := int(ChangeNumber)
		o.ChangeNumber = &ChangeNumberInt
	}
	
	if dateUploadedString, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["dateUploaded"].(string); ok {
		DateUploaded, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateUploadedString)
		o.DateUploaded = &DateUploaded
	}
	
	if UploadedBy, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["uploadedBy"].(map[string]interface{}); ok {
		UploadedByString, _ := json.Marshal(UploadedBy)
		json.Unmarshal(UploadedByString, &o.UploadedBy)
	}
	
	if LockInfo, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["lockInfo"].(map[string]interface{}); ok {
		LockInfoString, _ := json.Marshal(LockInfo)
		json.Unmarshal(LockInfoString, &o.LockInfo)
	}
	
	if SelfUri, ok := Contentmanagementworkspacedocumentstopicdocumentdatav2Map["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contentmanagementworkspacedocumentstopicdocumentdatav2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
