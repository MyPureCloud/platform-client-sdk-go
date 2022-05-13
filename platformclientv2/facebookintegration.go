package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facebookintegration
type Facebookintegration struct { 
	// Id - A unique Integration Id.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Facebook Integration
	Name *string `json:"name,omitempty"`


	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`


	// MessagingSetting
	MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`


	// AppId - The App Id from Facebook messenger
	AppId *string `json:"appId,omitempty"`


	// PageId - The Page Id from Facebook messenger
	PageId *string `json:"pageId,omitempty"`


	// PageName - The name of the Facebook page
	PageName *string `json:"pageName,omitempty"`


	// PageProfileImageUrl - The url of the profile image of the Facebook page
	PageProfileImageUrl *string `json:"pageProfileImageUrl,omitempty"`


	// Status - The status of the Facebook Integration
	Status *string `json:"status,omitempty"`


	// Recipient - The recipient reference associated to the Facebook Integration. This recipient is used to associate a flow to an integration
	Recipient *Domainentityref `json:"recipient,omitempty"`


	// DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this Integration was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User reference that created this Integration
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy - User reference that last modified this Integration
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Version - Version number required for updates.
	Version *int `json:"version,omitempty"`


	// CreateStatus - Status of asynchronous create operation
	CreateStatus *string `json:"createStatus,omitempty"`


	// CreateError - Error information returned, if createStatus is set to Error
	CreateError *Errorbody `json:"createError,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Facebookintegration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facebookintegration
	
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
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`
		
		AppId *string `json:"appId,omitempty"`
		
		PageId *string `json:"pageId,omitempty"`
		
		PageName *string `json:"pageName,omitempty"`
		
		PageProfileImageUrl *string `json:"pageProfileImageUrl,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Recipient *Domainentityref `json:"recipient,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		CreateStatus *string `json:"createStatus,omitempty"`
		
		CreateError *Errorbody `json:"createError,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		AppId: o.AppId,
		
		PageId: o.PageId,
		
		PageName: o.PageName,
		
		PageProfileImageUrl: o.PageProfileImageUrl,
		
		Status: o.Status,
		
		Recipient: o.Recipient,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		Version: o.Version,
		
		CreateStatus: o.CreateStatus,
		
		CreateError: o.CreateError,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Facebookintegration) UnmarshalJSON(b []byte) error {
	var FacebookintegrationMap map[string]interface{}
	err := json.Unmarshal(b, &FacebookintegrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FacebookintegrationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FacebookintegrationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := FacebookintegrationMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := FacebookintegrationMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if AppId, ok := FacebookintegrationMap["appId"].(string); ok {
		o.AppId = &AppId
	}
    
	if PageId, ok := FacebookintegrationMap["pageId"].(string); ok {
		o.PageId = &PageId
	}
    
	if PageName, ok := FacebookintegrationMap["pageName"].(string); ok {
		o.PageName = &PageName
	}
    
	if PageProfileImageUrl, ok := FacebookintegrationMap["pageProfileImageUrl"].(string); ok {
		o.PageProfileImageUrl = &PageProfileImageUrl
	}
    
	if Status, ok := FacebookintegrationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Recipient, ok := FacebookintegrationMap["recipient"].(map[string]interface{}); ok {
		RecipientString, _ := json.Marshal(Recipient)
		json.Unmarshal(RecipientString, &o.Recipient)
	}
	
	if dateCreatedString, ok := FacebookintegrationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := FacebookintegrationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := FacebookintegrationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := FacebookintegrationMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Version, ok := FacebookintegrationMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if CreateStatus, ok := FacebookintegrationMap["createStatus"].(string); ok {
		o.CreateStatus = &CreateStatus
	}
    
	if CreateError, ok := FacebookintegrationMap["createError"].(map[string]interface{}); ok {
		CreateErrorString, _ := json.Marshal(CreateError)
		json.Unmarshal(CreateErrorString, &o.CreateError)
	}
	
	if SelfUri, ok := FacebookintegrationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Facebookintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
