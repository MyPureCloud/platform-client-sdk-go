package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Paymentrequest
type Paymentrequest struct { 
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
	LineItems *[]Paymentlineitem `json:"lineItems,omitempty"`

	// ShippingOptions - The available shipping options.
	ShippingOptions *[]Paymentlineitem `json:"shippingOptions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Paymentrequest) SetField(field string, fieldValue interface{}) {
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

func (o Paymentrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Paymentrequest
	
	return json.Marshal(&struct { 
		PaymentPlatform *string `json:"paymentPlatform,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		CurrencyCode *string `json:"currencyCode,omitempty"`
		
		OrderTotal *float64 `json:"orderTotal,omitempty"`
		
		LineItems *[]Paymentlineitem `json:"lineItems,omitempty"`
		
		ShippingOptions *[]Paymentlineitem `json:"shippingOptions,omitempty"`
		Alias
	}{ 
		PaymentPlatform: o.PaymentPlatform,
		
		CountryCode: o.CountryCode,
		
		CurrencyCode: o.CurrencyCode,
		
		OrderTotal: o.OrderTotal,
		
		LineItems: o.LineItems,
		
		ShippingOptions: o.ShippingOptions,
		Alias:    (Alias)(o),
	})
}

func (o *Paymentrequest) UnmarshalJSON(b []byte) error {
	var PaymentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &PaymentrequestMap)
	if err != nil {
		return err
	}
	
	if PaymentPlatform, ok := PaymentrequestMap["paymentPlatform"].(string); ok {
		o.PaymentPlatform = &PaymentPlatform
	}
    
	if CountryCode, ok := PaymentrequestMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if CurrencyCode, ok := PaymentrequestMap["currencyCode"].(string); ok {
		o.CurrencyCode = &CurrencyCode
	}
    
	if OrderTotal, ok := PaymentrequestMap["orderTotal"].(float64); ok {
		o.OrderTotal = &OrderTotal
	}
    
	if LineItems, ok := PaymentrequestMap["lineItems"].([]interface{}); ok {
		LineItemsString, _ := json.Marshal(LineItems)
		json.Unmarshal(LineItemsString, &o.LineItems)
	}
	
	if ShippingOptions, ok := PaymentrequestMap["shippingOptions"].([]interface{}); ok {
		ShippingOptionsString, _ := json.Marshal(ShippingOptions)
		json.Unmarshal(ShippingOptionsString, &o.ShippingOptions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Paymentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
