package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Billinginvoice
type Billinginvoice struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// BillToCustomer - The bill to customer.
	BillToCustomer *Customer `json:"billToCustomer,omitempty"`

	// ShipToCustomer - The ship to customer.
	ShipToCustomer *Customer `json:"shipToCustomer,omitempty"`

	// SoldToCustomer - The sold to customer.
	SoldToCustomer *Customer `json:"soldToCustomer,omitempty"`

	// DateInvoiced - Date when the invoice was issued. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateInvoiced *time.Time `json:"dateInvoiced,omitempty"`

	// BillToAddress - Details of the bill to address.
	BillToAddress *Invoiceaddress `json:"billToAddress,omitempty"`

	// ShipToAddress - Details of the ship to address.
	ShipToAddress *Invoiceaddress `json:"shipToAddress,omitempty"`

	// CurrencyIsoCode - Contains the ISO code for any currency allowed by the organization.
	CurrencyIsoCode *string `json:"currencyIsoCode,omitempty"`

	// PaymentStatus - Status of the payment.
	PaymentStatus *string `json:"paymentStatus,omitempty"`

	// PaymentTerms - Payment terms.
	PaymentTerms *string `json:"paymentTerms,omitempty"`

	// PaymentLink - Payment link.
	PaymentLink *string `json:"paymentLink,omitempty"`

	// CustomerPoNumber - Purchase Order Number.
	CustomerPoNumber *string `json:"customerPoNumber,omitempty"`

	// CustomerInvoiceType - Type of the invoice.
	CustomerInvoiceType *string `json:"customerInvoiceType,omitempty"`

	// Amount - Amount.
	Amount *float32 `json:"amount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Billinginvoice) SetField(field string, fieldValue interface{}) {
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

func (o Billinginvoice) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateInvoiced", }

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
	type Alias Billinginvoice
	
	DateInvoiced := new(string)
	if o.DateInvoiced != nil {
		*DateInvoiced = timeutil.Strftime(o.DateInvoiced, "%Y-%m-%d")
	} else {
		DateInvoiced = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		BillToCustomer *Customer `json:"billToCustomer,omitempty"`
		
		ShipToCustomer *Customer `json:"shipToCustomer,omitempty"`
		
		SoldToCustomer *Customer `json:"soldToCustomer,omitempty"`
		
		DateInvoiced *string `json:"dateInvoiced,omitempty"`
		
		BillToAddress *Invoiceaddress `json:"billToAddress,omitempty"`
		
		ShipToAddress *Invoiceaddress `json:"shipToAddress,omitempty"`
		
		CurrencyIsoCode *string `json:"currencyIsoCode,omitempty"`
		
		PaymentStatus *string `json:"paymentStatus,omitempty"`
		
		PaymentTerms *string `json:"paymentTerms,omitempty"`
		
		PaymentLink *string `json:"paymentLink,omitempty"`
		
		CustomerPoNumber *string `json:"customerPoNumber,omitempty"`
		
		CustomerInvoiceType *string `json:"customerInvoiceType,omitempty"`
		
		Amount *float32 `json:"amount,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		BillToCustomer: o.BillToCustomer,
		
		ShipToCustomer: o.ShipToCustomer,
		
		SoldToCustomer: o.SoldToCustomer,
		
		DateInvoiced: DateInvoiced,
		
		BillToAddress: o.BillToAddress,
		
		ShipToAddress: o.ShipToAddress,
		
		CurrencyIsoCode: o.CurrencyIsoCode,
		
		PaymentStatus: o.PaymentStatus,
		
		PaymentTerms: o.PaymentTerms,
		
		PaymentLink: o.PaymentLink,
		
		CustomerPoNumber: o.CustomerPoNumber,
		
		CustomerInvoiceType: o.CustomerInvoiceType,
		
		Amount: o.Amount,
		Alias:    (Alias)(o),
	})
}

func (o *Billinginvoice) UnmarshalJSON(b []byte) error {
	var BillinginvoiceMap map[string]interface{}
	err := json.Unmarshal(b, &BillinginvoiceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BillinginvoiceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if BillToCustomer, ok := BillinginvoiceMap["billToCustomer"].(map[string]interface{}); ok {
		BillToCustomerString, _ := json.Marshal(BillToCustomer)
		json.Unmarshal(BillToCustomerString, &o.BillToCustomer)
	}
	
	if ShipToCustomer, ok := BillinginvoiceMap["shipToCustomer"].(map[string]interface{}); ok {
		ShipToCustomerString, _ := json.Marshal(ShipToCustomer)
		json.Unmarshal(ShipToCustomerString, &o.ShipToCustomer)
	}
	
	if SoldToCustomer, ok := BillinginvoiceMap["soldToCustomer"].(map[string]interface{}); ok {
		SoldToCustomerString, _ := json.Marshal(SoldToCustomer)
		json.Unmarshal(SoldToCustomerString, &o.SoldToCustomer)
	}
	
	if dateInvoicedString, ok := BillinginvoiceMap["dateInvoiced"].(string); ok {
		DateInvoiced, _ := time.Parse("2006-01-02", dateInvoicedString)
		o.DateInvoiced = &DateInvoiced
	}
	
	if BillToAddress, ok := BillinginvoiceMap["billToAddress"].(map[string]interface{}); ok {
		BillToAddressString, _ := json.Marshal(BillToAddress)
		json.Unmarshal(BillToAddressString, &o.BillToAddress)
	}
	
	if ShipToAddress, ok := BillinginvoiceMap["shipToAddress"].(map[string]interface{}); ok {
		ShipToAddressString, _ := json.Marshal(ShipToAddress)
		json.Unmarshal(ShipToAddressString, &o.ShipToAddress)
	}
	
	if CurrencyIsoCode, ok := BillinginvoiceMap["currencyIsoCode"].(string); ok {
		o.CurrencyIsoCode = &CurrencyIsoCode
	}
    
	if PaymentStatus, ok := BillinginvoiceMap["paymentStatus"].(string); ok {
		o.PaymentStatus = &PaymentStatus
	}
    
	if PaymentTerms, ok := BillinginvoiceMap["paymentTerms"].(string); ok {
		o.PaymentTerms = &PaymentTerms
	}
    
	if PaymentLink, ok := BillinginvoiceMap["paymentLink"].(string); ok {
		o.PaymentLink = &PaymentLink
	}
    
	if CustomerPoNumber, ok := BillinginvoiceMap["customerPoNumber"].(string); ok {
		o.CustomerPoNumber = &CustomerPoNumber
	}
    
	if CustomerInvoiceType, ok := BillinginvoiceMap["customerInvoiceType"].(string); ok {
		o.CustomerInvoiceType = &CustomerInvoiceType
	}
    
	if Amount, ok := BillinginvoiceMap["amount"].(float64); ok {
		AmountFloat32 := float32(Amount)
		o.Amount = &AmountFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Billinginvoice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
