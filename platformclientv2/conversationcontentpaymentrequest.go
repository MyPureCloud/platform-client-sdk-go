package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentpaymentrequest - Payment Request object used to request payment from a customer.
type Conversationcontentpaymentrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PaymentPlatform - The payment platform being used (e.g. Apple Pay)
	PaymentPlatform *string `json:"paymentPlatform,omitempty"`

	// CountryCode - The merchant's two-letter ISO 3166 country code.
	CountryCode *string `json:"countryCode,omitempty"`

	// CurrencyCode - The three-letter ISO 4217 currency code for the payment.
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// OrderTotal - The total price of the order.
	OrderTotal *float64 `json:"orderTotal,omitempty"`

	// LineItems - The items that make up the order.
	LineItems *[]Conversationcontentlineitem `json:"lineItems,omitempty"`

	// ShippingOptions - The available shipping options.
	ShippingOptions *[]Conversationcontentlineitem `json:"shippingOptions,omitempty"`

	// RequiredContactFields - Contact fields required to complete the order.
	RequiredContactFields *[]Conversationcontentrequiredcontactfield `json:"requiredContactFields,omitempty"`

	// ReceivedMessage - The message prompt to complete a payment transaction.
	ReceivedMessage *Conversationcontentreceivedreplymessage `json:"receivedMessage,omitempty"`

	// ReplyMessage - The reply message after the user has completed the payment transaction.
	ReplyMessage *Conversationcontentreceivedreplymessage `json:"replyMessage,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentpaymentrequest) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentpaymentrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationcontentpaymentrequest
	
	return json.Marshal(&struct { 
		PaymentPlatform *string `json:"paymentPlatform,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		CurrencyCode *string `json:"currencyCode,omitempty"`
		
		OrderTotal *float64 `json:"orderTotal,omitempty"`
		
		LineItems *[]Conversationcontentlineitem `json:"lineItems,omitempty"`
		
		ShippingOptions *[]Conversationcontentlineitem `json:"shippingOptions,omitempty"`
		
		RequiredContactFields *[]Conversationcontentrequiredcontactfield `json:"requiredContactFields,omitempty"`
		
		ReceivedMessage *Conversationcontentreceivedreplymessage `json:"receivedMessage,omitempty"`
		
		ReplyMessage *Conversationcontentreceivedreplymessage `json:"replyMessage,omitempty"`
		Alias
	}{ 
		PaymentPlatform: o.PaymentPlatform,
		
		CountryCode: o.CountryCode,
		
		CurrencyCode: o.CurrencyCode,
		
		OrderTotal: o.OrderTotal,
		
		LineItems: o.LineItems,
		
		ShippingOptions: o.ShippingOptions,
		
		RequiredContactFields: o.RequiredContactFields,
		
		ReceivedMessage: o.ReceivedMessage,
		
		ReplyMessage: o.ReplyMessage,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentpaymentrequest) UnmarshalJSON(b []byte) error {
	var ConversationcontentpaymentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentpaymentrequestMap)
	if err != nil {
		return err
	}
	
	if PaymentPlatform, ok := ConversationcontentpaymentrequestMap["paymentPlatform"].(string); ok {
		o.PaymentPlatform = &PaymentPlatform
	}
    
	if CountryCode, ok := ConversationcontentpaymentrequestMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if CurrencyCode, ok := ConversationcontentpaymentrequestMap["currencyCode"].(string); ok {
		o.CurrencyCode = &CurrencyCode
	}
    
	if OrderTotal, ok := ConversationcontentpaymentrequestMap["orderTotal"].(float64); ok {
		o.OrderTotal = &OrderTotal
	}
    
	if LineItems, ok := ConversationcontentpaymentrequestMap["lineItems"].([]interface{}); ok {
		LineItemsString, _ := json.Marshal(LineItems)
		json.Unmarshal(LineItemsString, &o.LineItems)
	}
	
	if ShippingOptions, ok := ConversationcontentpaymentrequestMap["shippingOptions"].([]interface{}); ok {
		ShippingOptionsString, _ := json.Marshal(ShippingOptions)
		json.Unmarshal(ShippingOptionsString, &o.ShippingOptions)
	}
	
	if RequiredContactFields, ok := ConversationcontentpaymentrequestMap["requiredContactFields"].([]interface{}); ok {
		RequiredContactFieldsString, _ := json.Marshal(RequiredContactFields)
		json.Unmarshal(RequiredContactFieldsString, &o.RequiredContactFields)
	}
	
	if ReceivedMessage, ok := ConversationcontentpaymentrequestMap["receivedMessage"].(map[string]interface{}); ok {
		ReceivedMessageString, _ := json.Marshal(ReceivedMessage)
		json.Unmarshal(ReceivedMessageString, &o.ReceivedMessage)
	}
	
	if ReplyMessage, ok := ConversationcontentpaymentrequestMap["replyMessage"].(map[string]interface{}); ok {
		ReplyMessageString, _ := json.Marshal(ReplyMessage)
		json.Unmarshal(ReplyMessageString, &o.ReplyMessage)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentpaymentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
