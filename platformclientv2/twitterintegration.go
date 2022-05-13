package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Twitterintegration
type Twitterintegration struct { 
	// Id - A unique Integration Id
	Id *string `json:"id,omitempty"`


	// Name - The name of the Twitter Integration
	Name *string `json:"name,omitempty"`


	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`


	// MessagingSetting
	MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`


	// AccessTokenKey - The Access Token Key from Twitter messenger
	AccessTokenKey *string `json:"accessTokenKey,omitempty"`


	// ConsumerKey - The Consumer Key from Twitter messenger
	ConsumerKey *string `json:"consumerKey,omitempty"`


	// Username - The Username from Twitter
	Username *string `json:"username,omitempty"`


	// UserId - The UserId from Twitter
	UserId *string `json:"userId,omitempty"`


	// Status - The status of the Twitter Integration
	Status *string `json:"status,omitempty"`


	// Tier - The type of twitter account to be used for the integration
	Tier *string `json:"tier,omitempty"`


	// EnvName - The Twitter environment name, e.g.: env-beta (required for premium tier)
	EnvName *string `json:"envName,omitempty"`


	// Recipient - The recipient associated to the Twitter Integration. This recipient is used to associate a flow to an integration
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

func (o *Twitterintegration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Twitterintegration
	
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
		
		AccessTokenKey *string `json:"accessTokenKey,omitempty"`
		
		ConsumerKey *string `json:"consumerKey,omitempty"`
		
		Username *string `json:"username,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Tier *string `json:"tier,omitempty"`
		
		EnvName *string `json:"envName,omitempty"`
		
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
		
		AccessTokenKey: o.AccessTokenKey,
		
		ConsumerKey: o.ConsumerKey,
		
		Username: o.Username,
		
		UserId: o.UserId,
		
		Status: o.Status,
		
		Tier: o.Tier,
		
		EnvName: o.EnvName,
		
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

func (o *Twitterintegration) UnmarshalJSON(b []byte) error {
	var TwitterintegrationMap map[string]interface{}
	err := json.Unmarshal(b, &TwitterintegrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TwitterintegrationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TwitterintegrationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := TwitterintegrationMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := TwitterintegrationMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if AccessTokenKey, ok := TwitterintegrationMap["accessTokenKey"].(string); ok {
		o.AccessTokenKey = &AccessTokenKey
	}
    
	if ConsumerKey, ok := TwitterintegrationMap["consumerKey"].(string); ok {
		o.ConsumerKey = &ConsumerKey
	}
    
	if Username, ok := TwitterintegrationMap["username"].(string); ok {
		o.Username = &Username
	}
    
	if UserId, ok := TwitterintegrationMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Status, ok := TwitterintegrationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Tier, ok := TwitterintegrationMap["tier"].(string); ok {
		o.Tier = &Tier
	}
    
	if EnvName, ok := TwitterintegrationMap["envName"].(string); ok {
		o.EnvName = &EnvName
	}
    
	if Recipient, ok := TwitterintegrationMap["recipient"].(map[string]interface{}); ok {
		RecipientString, _ := json.Marshal(Recipient)
		json.Unmarshal(RecipientString, &o.Recipient)
	}
	
	if dateCreatedString, ok := TwitterintegrationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := TwitterintegrationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := TwitterintegrationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := TwitterintegrationMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Version, ok := TwitterintegrationMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if CreateStatus, ok := TwitterintegrationMap["createStatus"].(string); ok {
		o.CreateStatus = &CreateStatus
	}
    
	if CreateError, ok := TwitterintegrationMap["createError"].(map[string]interface{}); ok {
		CreateErrorString, _ := json.Marshal(CreateError)
		json.Unmarshal(CreateErrorString, &o.CreateError)
	}
	
	if SelfUri, ok := TwitterintegrationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Twitterintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
