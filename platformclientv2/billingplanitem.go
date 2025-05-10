package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Billingplanitem
type Billingplanitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ItemNumber - Item number.
	ItemNumber *string `json:"itemNumber,omitempty"`

	// Name - Name of the item.
	Name *string `json:"name,omitempty"`

	// VarType - Type of the item.
	VarType *string `json:"type,omitempty"`

	// Function - Function of the item.
	Function *string `json:"function,omitempty"`

	// Description - Detailed description of the item.
	Description *string `json:"description,omitempty"`

	// DateChargedThrough - The date through which a customer has been billed for the charge. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateChargedThrough *time.Time `json:"dateChargedThrough,omitempty"`

	// CurrencyIsoCode - Contains the ISO code for any currency allowed by the organization.
	CurrencyIsoCode *string `json:"currencyIsoCode,omitempty"`

	// DiscountAmount - The amount of the discount.
	DiscountAmount *float32 `json:"discountAmount,omitempty"`

	// DateEffectiveStart - The date when the Address became effective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEffectiveStart *time.Time `json:"dateEffectiveStart,omitempty"`

	// DateEffectiveEnd - The date when the Address became effective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEffectiveEnd *time.Time `json:"dateEffectiveEnd,omitempty"`

	// OveragePrice - The price for units over the allowed amount.
	OveragePrice *float32 `json:"overagePrice,omitempty"`

	// Price - The price associated with the item expressed as a decimal.
	Price *float32 `json:"price,omitempty"`

	// Quantity - The quantity of units.
	Quantity *int `json:"quantity,omitempty"`

	// UnitOfMeasure - The unit of measure for the wallet.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Billingplanitem) SetField(field string, fieldValue interface{}) {
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

func (o Billingplanitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateChargedThrough","DateEffectiveStart","DateEffectiveEnd", }

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
	type Alias Billingplanitem
	
	DateChargedThrough := new(string)
	if o.DateChargedThrough != nil {
		*DateChargedThrough = timeutil.Strftime(o.DateChargedThrough, "%Y-%m-%d")
	} else {
		DateChargedThrough = nil
	}
	
	DateEffectiveStart := new(string)
	if o.DateEffectiveStart != nil {
		*DateEffectiveStart = timeutil.Strftime(o.DateEffectiveStart, "%Y-%m-%d")
	} else {
		DateEffectiveStart = nil
	}
	
	DateEffectiveEnd := new(string)
	if o.DateEffectiveEnd != nil {
		*DateEffectiveEnd = timeutil.Strftime(o.DateEffectiveEnd, "%Y-%m-%d")
	} else {
		DateEffectiveEnd = nil
	}
	
	return json.Marshal(&struct { 
		ItemNumber *string `json:"itemNumber,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Function *string `json:"function,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateChargedThrough *string `json:"dateChargedThrough,omitempty"`
		
		CurrencyIsoCode *string `json:"currencyIsoCode,omitempty"`
		
		DiscountAmount *float32 `json:"discountAmount,omitempty"`
		
		DateEffectiveStart *string `json:"dateEffectiveStart,omitempty"`
		
		DateEffectiveEnd *string `json:"dateEffectiveEnd,omitempty"`
		
		OveragePrice *float32 `json:"overagePrice,omitempty"`
		
		Price *float32 `json:"price,omitempty"`
		
		Quantity *int `json:"quantity,omitempty"`
		
		UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
		Alias
	}{ 
		ItemNumber: o.ItemNumber,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Function: o.Function,
		
		Description: o.Description,
		
		DateChargedThrough: DateChargedThrough,
		
		CurrencyIsoCode: o.CurrencyIsoCode,
		
		DiscountAmount: o.DiscountAmount,
		
		DateEffectiveStart: DateEffectiveStart,
		
		DateEffectiveEnd: DateEffectiveEnd,
		
		OveragePrice: o.OveragePrice,
		
		Price: o.Price,
		
		Quantity: o.Quantity,
		
		UnitOfMeasure: o.UnitOfMeasure,
		Alias:    (Alias)(o),
	})
}

func (o *Billingplanitem) UnmarshalJSON(b []byte) error {
	var BillingplanitemMap map[string]interface{}
	err := json.Unmarshal(b, &BillingplanitemMap)
	if err != nil {
		return err
	}
	
	if ItemNumber, ok := BillingplanitemMap["itemNumber"].(string); ok {
		o.ItemNumber = &ItemNumber
	}
    
	if Name, ok := BillingplanitemMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := BillingplanitemMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Function, ok := BillingplanitemMap["function"].(string); ok {
		o.Function = &Function
	}
    
	if Description, ok := BillingplanitemMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateChargedThroughString, ok := BillingplanitemMap["dateChargedThrough"].(string); ok {
		DateChargedThrough, _ := time.Parse("2006-01-02", dateChargedThroughString)
		o.DateChargedThrough = &DateChargedThrough
	}
	
	if CurrencyIsoCode, ok := BillingplanitemMap["currencyIsoCode"].(string); ok {
		o.CurrencyIsoCode = &CurrencyIsoCode
	}
    
	if DiscountAmount, ok := BillingplanitemMap["discountAmount"].(float64); ok {
		DiscountAmountFloat32 := float32(DiscountAmount)
		o.DiscountAmount = &DiscountAmountFloat32
	}
    
	if dateEffectiveStartString, ok := BillingplanitemMap["dateEffectiveStart"].(string); ok {
		DateEffectiveStart, _ := time.Parse("2006-01-02", dateEffectiveStartString)
		o.DateEffectiveStart = &DateEffectiveStart
	}
	
	if dateEffectiveEndString, ok := BillingplanitemMap["dateEffectiveEnd"].(string); ok {
		DateEffectiveEnd, _ := time.Parse("2006-01-02", dateEffectiveEndString)
		o.DateEffectiveEnd = &DateEffectiveEnd
	}
	
	if OveragePrice, ok := BillingplanitemMap["overagePrice"].(float64); ok {
		OveragePriceFloat32 := float32(OveragePrice)
		o.OveragePrice = &OveragePriceFloat32
	}
    
	if Price, ok := BillingplanitemMap["price"].(float64); ok {
		PriceFloat32 := float32(Price)
		o.Price = &PriceFloat32
	}
    
	if Quantity, ok := BillingplanitemMap["quantity"].(float64); ok {
		QuantityInt := int(Quantity)
		o.Quantity = &QuantityInt
	}
	
	if UnitOfMeasure, ok := BillingplanitemMap["unitOfMeasure"].(string); ok {
		o.UnitOfMeasure = &UnitOfMeasure
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Billingplanitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
