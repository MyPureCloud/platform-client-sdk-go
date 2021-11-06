package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Smsphonenumber
type Smsphonenumber struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// PhoneNumber - A phone number provisioned for SMS communications in E.164 format. E.g. +13175555555 or +34234234234
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// PhoneNumberType - Type of the phone number provisioned.
	PhoneNumberType *string `json:"phoneNumberType,omitempty"`


	// ProvisionedThroughPureCloud - Is set to false, if the phone number is provisioned through a SMS provider, outside of PureCloud
	ProvisionedThroughPureCloud *bool `json:"provisionedThroughPureCloud,omitempty"`


	// PhoneNumberStatus - Status of the provisioned phone number.
	PhoneNumberStatus *string `json:"phoneNumberStatus,omitempty"`


	// Capabilities - The capabilities of the phone number available for provisioning.
	Capabilities *[]string `json:"capabilities,omitempty"`


	// CountryCode - The ISO 3166-1 alpha-2 country code of the country this phone number is associated with.
	CountryCode *string `json:"countryCode,omitempty"`


	// DateCreated - Date this phone number was provisioned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this phone number was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User that provisioned this phone number
	CreatedBy *User `json:"createdBy,omitempty"`


	// ModifiedBy - User that last modified this phone number
	ModifiedBy *User `json:"modifiedBy,omitempty"`


	// Version - Version number required for updates.
	Version *int `json:"version,omitempty"`


	// PurchaseDate - Date this phone number was purchased, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PurchaseDate *time.Time `json:"purchaseDate,omitempty"`


	// CancellationDate - Contract end date of this phone number, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CancellationDate *time.Time `json:"cancellationDate,omitempty"`


	// RenewalDate - Contract renewal date of this phone number, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RenewalDate *time.Time `json:"renewalDate,omitempty"`


	// AutoRenewable - Renewal time period of this phone number, if the phoneNumberType is shortcode.
	AutoRenewable *string `json:"autoRenewable,omitempty"`


	// AddressId - The id of an address attached to this phone number.
	AddressId *Smsaddress `json:"addressId,omitempty"`


	// ShortCodeBillingType - BillingType of this phone number, if the phoneNumberType is shortcode.
	ShortCodeBillingType *string `json:"shortCodeBillingType,omitempty"`


	// ProvisioningStatus - Status of latest asynchronous provisioning action
	ProvisioningStatus *Smsprovisioningstatus `json:"provisioningStatus,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Smsphonenumber) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Smsphonenumber
	
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
	
	PurchaseDate := new(string)
	if o.PurchaseDate != nil {
		
		*PurchaseDate = timeutil.Strftime(o.PurchaseDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PurchaseDate = nil
	}
	
	CancellationDate := new(string)
	if o.CancellationDate != nil {
		
		*CancellationDate = timeutil.Strftime(o.CancellationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CancellationDate = nil
	}
	
	RenewalDate := new(string)
	if o.RenewalDate != nil {
		
		*RenewalDate = timeutil.Strftime(o.RenewalDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RenewalDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		PhoneNumberType *string `json:"phoneNumberType,omitempty"`
		
		ProvisionedThroughPureCloud *bool `json:"provisionedThroughPureCloud,omitempty"`
		
		PhoneNumberStatus *string `json:"phoneNumberStatus,omitempty"`
		
		Capabilities *[]string `json:"capabilities,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *User `json:"createdBy,omitempty"`
		
		ModifiedBy *User `json:"modifiedBy,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		PurchaseDate *string `json:"purchaseDate,omitempty"`
		
		CancellationDate *string `json:"cancellationDate,omitempty"`
		
		RenewalDate *string `json:"renewalDate,omitempty"`
		
		AutoRenewable *string `json:"autoRenewable,omitempty"`
		
		AddressId *Smsaddress `json:"addressId,omitempty"`
		
		ShortCodeBillingType *string `json:"shortCodeBillingType,omitempty"`
		
		ProvisioningStatus *Smsprovisioningstatus `json:"provisioningStatus,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		PhoneNumber: o.PhoneNumber,
		
		PhoneNumberType: o.PhoneNumberType,
		
		ProvisionedThroughPureCloud: o.ProvisionedThroughPureCloud,
		
		PhoneNumberStatus: o.PhoneNumberStatus,
		
		Capabilities: o.Capabilities,
		
		CountryCode: o.CountryCode,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		Version: o.Version,
		
		PurchaseDate: PurchaseDate,
		
		CancellationDate: CancellationDate,
		
		RenewalDate: RenewalDate,
		
		AutoRenewable: o.AutoRenewable,
		
		AddressId: o.AddressId,
		
		ShortCodeBillingType: o.ShortCodeBillingType,
		
		ProvisioningStatus: o.ProvisioningStatus,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Smsphonenumber) UnmarshalJSON(b []byte) error {
	var SmsphonenumberMap map[string]interface{}
	err := json.Unmarshal(b, &SmsphonenumberMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SmsphonenumberMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := SmsphonenumberMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if PhoneNumber, ok := SmsphonenumberMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
	
	if PhoneNumberType, ok := SmsphonenumberMap["phoneNumberType"].(string); ok {
		o.PhoneNumberType = &PhoneNumberType
	}
	
	if ProvisionedThroughPureCloud, ok := SmsphonenumberMap["provisionedThroughPureCloud"].(bool); ok {
		o.ProvisionedThroughPureCloud = &ProvisionedThroughPureCloud
	}
	
	if PhoneNumberStatus, ok := SmsphonenumberMap["phoneNumberStatus"].(string); ok {
		o.PhoneNumberStatus = &PhoneNumberStatus
	}
	
	if Capabilities, ok := SmsphonenumberMap["capabilities"].([]interface{}); ok {
		CapabilitiesString, _ := json.Marshal(Capabilities)
		json.Unmarshal(CapabilitiesString, &o.Capabilities)
	}
	
	if CountryCode, ok := SmsphonenumberMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
	
	if dateCreatedString, ok := SmsphonenumberMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := SmsphonenumberMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := SmsphonenumberMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := SmsphonenumberMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Version, ok := SmsphonenumberMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if purchaseDateString, ok := SmsphonenumberMap["purchaseDate"].(string); ok {
		PurchaseDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", purchaseDateString)
		o.PurchaseDate = &PurchaseDate
	}
	
	if cancellationDateString, ok := SmsphonenumberMap["cancellationDate"].(string); ok {
		CancellationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", cancellationDateString)
		o.CancellationDate = &CancellationDate
	}
	
	if renewalDateString, ok := SmsphonenumberMap["renewalDate"].(string); ok {
		RenewalDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", renewalDateString)
		o.RenewalDate = &RenewalDate
	}
	
	if AutoRenewable, ok := SmsphonenumberMap["autoRenewable"].(string); ok {
		o.AutoRenewable = &AutoRenewable
	}
	
	if AddressId, ok := SmsphonenumberMap["addressId"].(map[string]interface{}); ok {
		AddressIdString, _ := json.Marshal(AddressId)
		json.Unmarshal(AddressIdString, &o.AddressId)
	}
	
	if ShortCodeBillingType, ok := SmsphonenumberMap["shortCodeBillingType"].(string); ok {
		o.ShortCodeBillingType = &ShortCodeBillingType
	}
	
	if ProvisioningStatus, ok := SmsphonenumberMap["provisioningStatus"].(map[string]interface{}); ok {
		ProvisioningStatusString, _ := json.Marshal(ProvisioningStatus)
		json.Unmarshal(ProvisioningStatusString, &o.ProvisioningStatus)
	}
	
	if SelfUri, ok := SmsphonenumberMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Smsphonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
