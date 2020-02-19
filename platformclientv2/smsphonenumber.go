package platformclientv2
import (
	"time"
	"encoding/json"
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


	// DateCreated - Date this phone number was provisioned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this phone number was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User that provisioned this phone number
	CreatedBy *User `json:"createdBy,omitempty"`


	// ModifiedBy - User that last modified this phone number
	ModifiedBy *User `json:"modifiedBy,omitempty"`


	// Version - Version number required for updates.
	Version *int32 `json:"version,omitempty"`


	// PurchaseDate - Date this phone number was purchased, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	PurchaseDate *time.Time `json:"purchaseDate,omitempty"`


	// CancellationDate - Contract end date of this phone number, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CancellationDate *time.Time `json:"cancellationDate,omitempty"`


	// RenewalDate - Contract renewal date of this phone number, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
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

// String returns a JSON representation of the model
func (o *Smsphonenumber) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
