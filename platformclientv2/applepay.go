package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Applepay
type Applepay struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StoreName - The name of the store.
	StoreName *string `json:"storeName,omitempty"`

	// MerchantId - The stores merchant identifier.
	MerchantId *string `json:"merchantId,omitempty"`

	// DomainName - The domain name associated with the merchant account.
	DomainName *string `json:"domainName,omitempty"`

	// PaymentCapabilities - The payment capabilities supported by the merchant.
	PaymentCapabilities *[]string `json:"paymentCapabilities,omitempty"`

	// SupportedPaymentNetworks - The payment networks supported by the merchant.
	SupportedPaymentNetworks *[]string `json:"supportedPaymentNetworks,omitempty"`

	// PaymentCertificateCredentialId - The Genesys credentialId the payment certificates are stored under.
	PaymentCertificateCredentialId *string `json:"paymentCertificateCredentialId,omitempty"`

	// PaymentGatewayUrl - The url used to process payments.
	PaymentGatewayUrl *string `json:"paymentGatewayUrl,omitempty"`

	// FallbackUrl - The url opened in a web browser if the customers device is unable to make payments using Apple Pay.
	FallbackUrl *string `json:"fallbackUrl,omitempty"`

	// ShippingMethodUpdateUrl - The url called when the customer changes the shipping method.
	ShippingMethodUpdateUrl *string `json:"shippingMethodUpdateUrl,omitempty"`

	// ShippingContactUpdateUrl - The url called when the customer changes their shipping address information.
	ShippingContactUpdateUrl *string `json:"shippingContactUpdateUrl,omitempty"`

	// PaymentMethodUpdateUrl - The url called when the customer changes their payment method.
	PaymentMethodUpdateUrl *string `json:"paymentMethodUpdateUrl,omitempty"`

	// OrderTrackingUrl - The url called after completing the order to update the order information in your system
	OrderTrackingUrl *string `json:"orderTrackingUrl,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Applepay) SetField(field string, fieldValue interface{}) {
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

func (o Applepay) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Applepay
	
	return json.Marshal(&struct { 
		StoreName *string `json:"storeName,omitempty"`
		
		MerchantId *string `json:"merchantId,omitempty"`
		
		DomainName *string `json:"domainName,omitempty"`
		
		PaymentCapabilities *[]string `json:"paymentCapabilities,omitempty"`
		
		SupportedPaymentNetworks *[]string `json:"supportedPaymentNetworks,omitempty"`
		
		PaymentCertificateCredentialId *string `json:"paymentCertificateCredentialId,omitempty"`
		
		PaymentGatewayUrl *string `json:"paymentGatewayUrl,omitempty"`
		
		FallbackUrl *string `json:"fallbackUrl,omitempty"`
		
		ShippingMethodUpdateUrl *string `json:"shippingMethodUpdateUrl,omitempty"`
		
		ShippingContactUpdateUrl *string `json:"shippingContactUpdateUrl,omitempty"`
		
		PaymentMethodUpdateUrl *string `json:"paymentMethodUpdateUrl,omitempty"`
		
		OrderTrackingUrl *string `json:"orderTrackingUrl,omitempty"`
		Alias
	}{ 
		StoreName: o.StoreName,
		
		MerchantId: o.MerchantId,
		
		DomainName: o.DomainName,
		
		PaymentCapabilities: o.PaymentCapabilities,
		
		SupportedPaymentNetworks: o.SupportedPaymentNetworks,
		
		PaymentCertificateCredentialId: o.PaymentCertificateCredentialId,
		
		PaymentGatewayUrl: o.PaymentGatewayUrl,
		
		FallbackUrl: o.FallbackUrl,
		
		ShippingMethodUpdateUrl: o.ShippingMethodUpdateUrl,
		
		ShippingContactUpdateUrl: o.ShippingContactUpdateUrl,
		
		PaymentMethodUpdateUrl: o.PaymentMethodUpdateUrl,
		
		OrderTrackingUrl: o.OrderTrackingUrl,
		Alias:    (Alias)(o),
	})
}

func (o *Applepay) UnmarshalJSON(b []byte) error {
	var ApplepayMap map[string]interface{}
	err := json.Unmarshal(b, &ApplepayMap)
	if err != nil {
		return err
	}
	
	if StoreName, ok := ApplepayMap["storeName"].(string); ok {
		o.StoreName = &StoreName
	}
    
	if MerchantId, ok := ApplepayMap["merchantId"].(string); ok {
		o.MerchantId = &MerchantId
	}
    
	if DomainName, ok := ApplepayMap["domainName"].(string); ok {
		o.DomainName = &DomainName
	}
    
	if PaymentCapabilities, ok := ApplepayMap["paymentCapabilities"].([]interface{}); ok {
		PaymentCapabilitiesString, _ := json.Marshal(PaymentCapabilities)
		json.Unmarshal(PaymentCapabilitiesString, &o.PaymentCapabilities)
	}
	
	if SupportedPaymentNetworks, ok := ApplepayMap["supportedPaymentNetworks"].([]interface{}); ok {
		SupportedPaymentNetworksString, _ := json.Marshal(SupportedPaymentNetworks)
		json.Unmarshal(SupportedPaymentNetworksString, &o.SupportedPaymentNetworks)
	}
	
	if PaymentCertificateCredentialId, ok := ApplepayMap["paymentCertificateCredentialId"].(string); ok {
		o.PaymentCertificateCredentialId = &PaymentCertificateCredentialId
	}
    
	if PaymentGatewayUrl, ok := ApplepayMap["paymentGatewayUrl"].(string); ok {
		o.PaymentGatewayUrl = &PaymentGatewayUrl
	}
    
	if FallbackUrl, ok := ApplepayMap["fallbackUrl"].(string); ok {
		o.FallbackUrl = &FallbackUrl
	}
    
	if ShippingMethodUpdateUrl, ok := ApplepayMap["shippingMethodUpdateUrl"].(string); ok {
		o.ShippingMethodUpdateUrl = &ShippingMethodUpdateUrl
	}
    
	if ShippingContactUpdateUrl, ok := ApplepayMap["shippingContactUpdateUrl"].(string); ok {
		o.ShippingContactUpdateUrl = &ShippingContactUpdateUrl
	}
    
	if PaymentMethodUpdateUrl, ok := ApplepayMap["paymentMethodUpdateUrl"].(string); ok {
		o.PaymentMethodUpdateUrl = &PaymentMethodUpdateUrl
	}
    
	if OrderTrackingUrl, ok := ApplepayMap["orderTrackingUrl"].(string); ok {
		o.OrderTrackingUrl = &OrderTrackingUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Applepay) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
