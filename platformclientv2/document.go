package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Document
type Document struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Document) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Document) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateUploaded", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Document
	
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
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ChangeNumber: o.ChangeNumber,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DateUploaded: DateUploaded,
		
		ContentUri: o.ContentUri,
		
		Workspace: o.Workspace,
		
		CreatedBy: o.CreatedBy,
		
		UploadedBy: o.UploadedBy,
		
		SharingUri: o.SharingUri,
		
		ContentType: o.ContentType,
		
		ContentLength: o.ContentLength,
		
		SystemType: o.SystemType,
		
		Filename: o.Filename,
		
		PageCount: o.PageCount,
		
		Read: o.Read,
		
		CallerAddress: o.CallerAddress,
		
		ReceiverAddress: o.ReceiverAddress,
		
		Tags: o.Tags,
		
		TagValues: o.TagValues,
		
		Attributes: o.Attributes,
		
		Thumbnails: o.Thumbnails,
		
		UploadStatus: o.UploadStatus,
		
		UploadDestinationUri: o.UploadDestinationUri,
		
		UploadMethod: o.UploadMethod,
		
		LockInfo: o.LockInfo,
		
		Acl: o.Acl,
		
		SharingStatus: o.SharingStatus,
		
		DownloadSharingUri: o.DownloadSharingUri,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Document) UnmarshalJSON(b []byte) error {
	var DocumentMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DocumentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DocumentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ChangeNumber, ok := DocumentMap["changeNumber"].(float64); ok {
		ChangeNumberInt := int(ChangeNumber)
		o.ChangeNumber = &ChangeNumberInt
	}
	
	if dateCreatedString, ok := DocumentMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DocumentMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateUploadedString, ok := DocumentMap["dateUploaded"].(string); ok {
		DateUploaded, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateUploadedString)
		o.DateUploaded = &DateUploaded
	}
	
	if ContentUri, ok := DocumentMap["contentUri"].(string); ok {
		o.ContentUri = &ContentUri
	}
    
	if Workspace, ok := DocumentMap["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if CreatedBy, ok := DocumentMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if UploadedBy, ok := DocumentMap["uploadedBy"].(map[string]interface{}); ok {
		UploadedByString, _ := json.Marshal(UploadedBy)
		json.Unmarshal(UploadedByString, &o.UploadedBy)
	}
	
	if SharingUri, ok := DocumentMap["sharingUri"].(string); ok {
		o.SharingUri = &SharingUri
	}
    
	if ContentType, ok := DocumentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ContentLength, ok := DocumentMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if SystemType, ok := DocumentMap["systemType"].(string); ok {
		o.SystemType = &SystemType
	}
    
	if Filename, ok := DocumentMap["filename"].(string); ok {
		o.Filename = &Filename
	}
    
	if PageCount, ok := DocumentMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if Read, ok := DocumentMap["read"].(bool); ok {
		o.Read = &Read
	}
    
	if CallerAddress, ok := DocumentMap["callerAddress"].(string); ok {
		o.CallerAddress = &CallerAddress
	}
    
	if ReceiverAddress, ok := DocumentMap["receiverAddress"].(string); ok {
		o.ReceiverAddress = &ReceiverAddress
	}
    
	if Tags, ok := DocumentMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if TagValues, ok := DocumentMap["tagValues"].([]interface{}); ok {
		TagValuesString, _ := json.Marshal(TagValues)
		json.Unmarshal(TagValuesString, &o.TagValues)
	}
	
	if Attributes, ok := DocumentMap["attributes"].([]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Thumbnails, ok := DocumentMap["thumbnails"].([]interface{}); ok {
		ThumbnailsString, _ := json.Marshal(Thumbnails)
		json.Unmarshal(ThumbnailsString, &o.Thumbnails)
	}
	
	if UploadStatus, ok := DocumentMap["uploadStatus"].(map[string]interface{}); ok {
		UploadStatusString, _ := json.Marshal(UploadStatus)
		json.Unmarshal(UploadStatusString, &o.UploadStatus)
	}
	
	if UploadDestinationUri, ok := DocumentMap["uploadDestinationUri"].(string); ok {
		o.UploadDestinationUri = &UploadDestinationUri
	}
    
	if UploadMethod, ok := DocumentMap["uploadMethod"].(string); ok {
		o.UploadMethod = &UploadMethod
	}
    
	if LockInfo, ok := DocumentMap["lockInfo"].(map[string]interface{}); ok {
		LockInfoString, _ := json.Marshal(LockInfo)
		json.Unmarshal(LockInfoString, &o.LockInfo)
	}
	
	if Acl, ok := DocumentMap["acl"].([]interface{}); ok {
		AclString, _ := json.Marshal(Acl)
		json.Unmarshal(AclString, &o.Acl)
	}
	
	if SharingStatus, ok := DocumentMap["sharingStatus"].(string); ok {
		o.SharingStatus = &SharingStatus
	}
    
	if DownloadSharingUri, ok := DocumentMap["downloadSharingUri"].(string); ok {
		o.DownloadSharingUri = &DownloadSharingUri
	}
    
	if SelfUri, ok := DocumentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Document) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
