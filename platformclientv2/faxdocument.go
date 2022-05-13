package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Faxdocument
type Faxdocument struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ContentUri
	ContentUri *string `json:"contentUri,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// CreatedBy
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// SharingUri
	SharingUri *string `json:"sharingUri,omitempty"`


	// ContentType
	ContentType *string `json:"contentType,omitempty"`


	// ContentLength
	ContentLength *int `json:"contentLength,omitempty"`


	// Filename
	Filename *string `json:"filename,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`


	// CallerAddress
	CallerAddress *string `json:"callerAddress,omitempty"`


	// ReceiverAddress
	ReceiverAddress *string `json:"receiverAddress,omitempty"`


	// Thumbnails
	Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`


	// DownloadSharingUri
	DownloadSharingUri *string `json:"downloadSharingUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Faxdocument) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Faxdocument
	
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
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ContentUri *string `json:"contentUri,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		SharingUri *string `json:"sharingUri,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		Filename *string `json:"filename,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		ReceiverAddress *string `json:"receiverAddress,omitempty"`
		
		Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`
		
		DownloadSharingUri *string `json:"downloadSharingUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ContentUri: o.ContentUri,
		
		Workspace: o.Workspace,
		
		CreatedBy: o.CreatedBy,
		
		SharingUri: o.SharingUri,
		
		ContentType: o.ContentType,
		
		ContentLength: o.ContentLength,
		
		Filename: o.Filename,
		
		Read: o.Read,
		
		PageCount: o.PageCount,
		
		CallerAddress: o.CallerAddress,
		
		ReceiverAddress: o.ReceiverAddress,
		
		Thumbnails: o.Thumbnails,
		
		DownloadSharingUri: o.DownloadSharingUri,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Faxdocument) UnmarshalJSON(b []byte) error {
	var FaxdocumentMap map[string]interface{}
	err := json.Unmarshal(b, &FaxdocumentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FaxdocumentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FaxdocumentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := FaxdocumentMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := FaxdocumentMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ContentUri, ok := FaxdocumentMap["contentUri"].(string); ok {
		o.ContentUri = &ContentUri
	}
    
	if Workspace, ok := FaxdocumentMap["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if CreatedBy, ok := FaxdocumentMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if SharingUri, ok := FaxdocumentMap["sharingUri"].(string); ok {
		o.SharingUri = &SharingUri
	}
    
	if ContentType, ok := FaxdocumentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ContentLength, ok := FaxdocumentMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if Filename, ok := FaxdocumentMap["filename"].(string); ok {
		o.Filename = &Filename
	}
    
	if Read, ok := FaxdocumentMap["read"].(bool); ok {
		o.Read = &Read
	}
    
	if PageCount, ok := FaxdocumentMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if CallerAddress, ok := FaxdocumentMap["callerAddress"].(string); ok {
		o.CallerAddress = &CallerAddress
	}
    
	if ReceiverAddress, ok := FaxdocumentMap["receiverAddress"].(string); ok {
		o.ReceiverAddress = &ReceiverAddress
	}
    
	if Thumbnails, ok := FaxdocumentMap["thumbnails"].([]interface{}); ok {
		ThumbnailsString, _ := json.Marshal(Thumbnails)
		json.Unmarshal(ThumbnailsString, &o.Thumbnails)
	}
	
	if DownloadSharingUri, ok := FaxdocumentMap["downloadSharingUri"].(string); ok {
		o.DownloadSharingUri = &DownloadSharingUri
	}
    
	if SelfUri, ok := FaxdocumentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Faxdocument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
