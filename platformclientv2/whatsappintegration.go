package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappintegration
type Whatsappintegration struct { 
	// Id - A unique Integration Id.
	Id *string `json:"id,omitempty"`


	// Name - The name of the WhatsApp integration.
	Name *string `json:"name,omitempty"`


	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`


	// MessagingSetting
	MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`


	// PhoneNumber - The phone number associated to the WhatsApp integration.
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// AvailablePhoneNumbers - The list of available WhatsApp phone numbers for this account. Please select one phone number from this list to use with the created integration.
	AvailablePhoneNumbers *Whatsappavailablephonenumberdetailslisting `json:"availablePhoneNumbers,omitempty"`


	// Status - The status of the WhatsApp Integration
	Status *string `json:"status,omitempty"`


	// Recipient - The recipient associated to the WhatsApp Integration. This recipient is used to associate a flow to an integration
	Recipient *Domainentityref `json:"recipient,omitempty"`


	// DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this Integration was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User reference that created this Integration
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy - User reference that last modified this Integration
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Version - Version number required for updates.
	Version *int `json:"version,omitempty"`


	// ActivationStatusCode - The status code of WhatsApp Integration activation process
	ActivationStatusCode *string `json:"activationStatusCode,omitempty"`


	// ActivationErrorInfo - The error information of WhatsApp Integration activation process
	ActivationErrorInfo *Errorbody `json:"activationErrorInfo,omitempty"`


	// CreateStatus - Status of asynchronous create operation
	CreateStatus *string `json:"createStatus,omitempty"`


	// CreateError - Error information returned, if createStatus is set to Error
	CreateError *Errorbody `json:"createError,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Whatsappintegration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappintegration
	
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
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		AvailablePhoneNumbers *Whatsappavailablephonenumberdetailslisting `json:"availablePhoneNumbers,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Recipient *Domainentityref `json:"recipient,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		ActivationStatusCode *string `json:"activationStatusCode,omitempty"`
		
		ActivationErrorInfo *Errorbody `json:"activationErrorInfo,omitempty"`
		
		CreateStatus *string `json:"createStatus,omitempty"`
		
		CreateError *Errorbody `json:"createError,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		PhoneNumber: o.PhoneNumber,
		
		AvailablePhoneNumbers: o.AvailablePhoneNumbers,
		
		Status: o.Status,
		
		Recipient: o.Recipient,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		Version: o.Version,
		
		ActivationStatusCode: o.ActivationStatusCode,
		
		ActivationErrorInfo: o.ActivationErrorInfo,
		
		CreateStatus: o.CreateStatus,
		
		CreateError: o.CreateError,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Whatsappintegration) UnmarshalJSON(b []byte) error {
	var WhatsappintegrationMap map[string]interface{}
	err := json.Unmarshal(b, &WhatsappintegrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WhatsappintegrationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WhatsappintegrationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := WhatsappintegrationMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := WhatsappintegrationMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if PhoneNumber, ok := WhatsappintegrationMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if AvailablePhoneNumbers, ok := WhatsappintegrationMap["availablePhoneNumbers"].(map[string]interface{}); ok {
		AvailablePhoneNumbersString, _ := json.Marshal(AvailablePhoneNumbers)
		json.Unmarshal(AvailablePhoneNumbersString, &o.AvailablePhoneNumbers)
	}
	
	if Status, ok := WhatsappintegrationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Recipient, ok := WhatsappintegrationMap["recipient"].(map[string]interface{}); ok {
		RecipientString, _ := json.Marshal(Recipient)
		json.Unmarshal(RecipientString, &o.Recipient)
	}
	
	if dateCreatedString, ok := WhatsappintegrationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := WhatsappintegrationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := WhatsappintegrationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := WhatsappintegrationMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Version, ok := WhatsappintegrationMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ActivationStatusCode, ok := WhatsappintegrationMap["activationStatusCode"].(string); ok {
		o.ActivationStatusCode = &ActivationStatusCode
	}
    
	if ActivationErrorInfo, ok := WhatsappintegrationMap["activationErrorInfo"].(map[string]interface{}); ok {
		ActivationErrorInfoString, _ := json.Marshal(ActivationErrorInfo)
		json.Unmarshal(ActivationErrorInfoString, &o.ActivationErrorInfo)
	}
	
	if CreateStatus, ok := WhatsappintegrationMap["createStatus"].(string); ok {
		o.CreateStatus = &CreateStatus
	}
    
	if CreateError, ok := WhatsappintegrationMap["createError"].(map[string]interface{}); ok {
		CreateErrorString, _ := json.Marshal(CreateError)
		json.Unmarshal(CreateErrorString, &o.CreateError)
	}
	
	if SelfUri, ok := WhatsappintegrationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Whatsappintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
