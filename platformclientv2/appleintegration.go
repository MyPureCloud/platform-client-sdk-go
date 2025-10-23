package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Appleintegration
type Appleintegration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A unique integration Id.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Apple messaging integration.
	Name *string `json:"name,omitempty"`

	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`

	// MessagingSetting
	MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`

	// MessagesForBusinessId - The Apple Messages for Business Id for the Apple messaging integration.
	MessagesForBusinessId *string `json:"messagesForBusinessId,omitempty"`

	// BusinessName - The name of the business.
	BusinessName *string `json:"businessName,omitempty"`

	// LogoUrl - The url of the businesses logo.
	LogoUrl *string `json:"logoUrl,omitempty"`

	// Status - The status of the Apple Integration
	Status *string `json:"status,omitempty"`

	// Recipient - The recipient associated to the Apple messaging Integration. This recipient is used to associate a flow to an integration
	Recipient *Domainentityref `json:"recipient,omitempty"`

	// DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Date this Integration was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// CreatedBy - User reference that created this Integration
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`

	// ModifiedBy - User reference that last modified this Integration
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`

	// CreateStatus - Status of asynchronous create operation
	CreateStatus *string `json:"createStatus,omitempty"`

	// CreateError - Error information returned, if createStatus is set to Error
	CreateError *Errorbody `json:"createError,omitempty"`

	// AppleIMessageApp - Interactive Application (iMessage App) Settings.
	AppleIMessageApp *Appleimessageapp `json:"appleIMessageApp,omitempty"`

	// AppleAuthentication - The Apple Messages for Business authentication setting.
	AppleAuthentication *Appleauthentication `json:"appleAuthentication,omitempty"`

	// ApplePay - Apple Pay settings.
	ApplePay *Applepay `json:"applePay,omitempty"`

	// IdentityResolution - The configuration to control identity resolution.
	IdentityResolution *Appleidentityresolutionconfig `json:"identityResolution,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Appleintegration) SetField(field string, fieldValue interface{}) {
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

func (o Appleintegration) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Appleintegration
	
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
		
		MessagesForBusinessId *string `json:"messagesForBusinessId,omitempty"`
		
		BusinessName *string `json:"businessName,omitempty"`
		
		LogoUrl *string `json:"logoUrl,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Recipient *Domainentityref `json:"recipient,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		CreateStatus *string `json:"createStatus,omitempty"`
		
		CreateError *Errorbody `json:"createError,omitempty"`
		
		AppleIMessageApp *Appleimessageapp `json:"appleIMessageApp,omitempty"`
		
		AppleAuthentication *Appleauthentication `json:"appleAuthentication,omitempty"`
		
		ApplePay *Applepay `json:"applePay,omitempty"`
		
		IdentityResolution *Appleidentityresolutionconfig `json:"identityResolution,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		MessagesForBusinessId: o.MessagesForBusinessId,
		
		BusinessName: o.BusinessName,
		
		LogoUrl: o.LogoUrl,
		
		Status: o.Status,
		
		Recipient: o.Recipient,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		CreateStatus: o.CreateStatus,
		
		CreateError: o.CreateError,
		
		AppleIMessageApp: o.AppleIMessageApp,
		
		AppleAuthentication: o.AppleAuthentication,
		
		ApplePay: o.ApplePay,
		
		IdentityResolution: o.IdentityResolution,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Appleintegration) UnmarshalJSON(b []byte) error {
	var AppleintegrationMap map[string]interface{}
	err := json.Unmarshal(b, &AppleintegrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AppleintegrationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AppleintegrationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := AppleintegrationMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := AppleintegrationMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if MessagesForBusinessId, ok := AppleintegrationMap["messagesForBusinessId"].(string); ok {
		o.MessagesForBusinessId = &MessagesForBusinessId
	}
    
	if BusinessName, ok := AppleintegrationMap["businessName"].(string); ok {
		o.BusinessName = &BusinessName
	}
    
	if LogoUrl, ok := AppleintegrationMap["logoUrl"].(string); ok {
		o.LogoUrl = &LogoUrl
	}
    
	if Status, ok := AppleintegrationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Recipient, ok := AppleintegrationMap["recipient"].(map[string]interface{}); ok {
		RecipientString, _ := json.Marshal(Recipient)
		json.Unmarshal(RecipientString, &o.Recipient)
	}
	
	if dateCreatedString, ok := AppleintegrationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := AppleintegrationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := AppleintegrationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := AppleintegrationMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if CreateStatus, ok := AppleintegrationMap["createStatus"].(string); ok {
		o.CreateStatus = &CreateStatus
	}
    
	if CreateError, ok := AppleintegrationMap["createError"].(map[string]interface{}); ok {
		CreateErrorString, _ := json.Marshal(CreateError)
		json.Unmarshal(CreateErrorString, &o.CreateError)
	}
	
	if AppleIMessageApp, ok := AppleintegrationMap["appleIMessageApp"].(map[string]interface{}); ok {
		AppleIMessageAppString, _ := json.Marshal(AppleIMessageApp)
		json.Unmarshal(AppleIMessageAppString, &o.AppleIMessageApp)
	}
	
	if AppleAuthentication, ok := AppleintegrationMap["appleAuthentication"].(map[string]interface{}); ok {
		AppleAuthenticationString, _ := json.Marshal(AppleAuthentication)
		json.Unmarshal(AppleAuthenticationString, &o.AppleAuthentication)
	}
	
	if ApplePay, ok := AppleintegrationMap["applePay"].(map[string]interface{}); ok {
		ApplePayString, _ := json.Marshal(ApplePay)
		json.Unmarshal(ApplePayString, &o.ApplePay)
	}
	
	if IdentityResolution, ok := AppleintegrationMap["identityResolution"].(map[string]interface{}); ok {
		IdentityResolutionString, _ := json.Marshal(IdentityResolution)
		json.Unmarshal(IdentityResolutionString, &o.IdentityResolution)
	}
	
	if SelfUri, ok := AppleintegrationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Appleintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
