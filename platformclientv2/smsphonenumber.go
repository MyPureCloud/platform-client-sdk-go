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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Smsphonenumber) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Smsphonenumber

	
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
	
	PurchaseDate := new(string)
	if u.PurchaseDate != nil {
		
		*PurchaseDate = timeutil.Strftime(u.PurchaseDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PurchaseDate = nil
	}
	
	CancellationDate := new(string)
	if u.CancellationDate != nil {
		
		*CancellationDate = timeutil.Strftime(u.CancellationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CancellationDate = nil
	}
	
	RenewalDate := new(string)
	if u.RenewalDate != nil {
		
		*RenewalDate = timeutil.Strftime(u.RenewalDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		PhoneNumber: u.PhoneNumber,
		
		PhoneNumberType: u.PhoneNumberType,
		
		ProvisionedThroughPureCloud: u.ProvisionedThroughPureCloud,
		
		PhoneNumberStatus: u.PhoneNumberStatus,
		
		Capabilities: u.Capabilities,
		
		CountryCode: u.CountryCode,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: u.CreatedBy,
		
		ModifiedBy: u.ModifiedBy,
		
		Version: u.Version,
		
		PurchaseDate: PurchaseDate,
		
		CancellationDate: CancellationDate,
		
		RenewalDate: RenewalDate,
		
		AutoRenewable: u.AutoRenewable,
		
		AddressId: u.AddressId,
		
		ShortCodeBillingType: u.ShortCodeBillingType,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Smsphonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
