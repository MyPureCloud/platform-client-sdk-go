package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Billinginvoiceitem
type Billinginvoiceitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Product - Represents the details of a product.
	Product *Billingproduct `json:"product,omitempty"`

	// Description - Line Item Description.
	Description *string `json:"description,omitempty"`

	// DateTransacted - Invoice transaction date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateTransacted *time.Time `json:"dateTransacted,omitempty"`

	// DateStart - Invoice start date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - Invoice end date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// Organization - A Genesys Cloud Organization.
	Organization *Namedentity `json:"organization,omitempty"`

	// Quantity - Quantity of the item.
	Quantity *int `json:"quantity,omitempty"`

	// UnitOfMeasure - Unit of Measure.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`

	// Amount - Amount.
	Amount *float32 `json:"amount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Billinginvoiceitem) SetField(field string, fieldValue interface{}) {
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

func (o Billinginvoiceitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateTransacted","DateStart","DateEnd", }

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
	type Alias Billinginvoiceitem
	
	DateTransacted := new(string)
	if o.DateTransacted != nil {
		*DateTransacted = timeutil.Strftime(o.DateTransacted, "%Y-%m-%d")
	} else {
		DateTransacted = nil
	}
	
	DateStart := new(string)
	if o.DateStart != nil {
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%d")
	} else {
		DateEnd = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Product *Billingproduct `json:"product,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateTransacted *string `json:"dateTransacted,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		Organization *Namedentity `json:"organization,omitempty"`
		
		Quantity *int `json:"quantity,omitempty"`
		
		UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
		
		Amount *float32 `json:"amount,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Product: o.Product,
		
		Description: o.Description,
		
		DateTransacted: DateTransacted,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		Organization: o.Organization,
		
		Quantity: o.Quantity,
		
		UnitOfMeasure: o.UnitOfMeasure,
		
		Amount: o.Amount,
		Alias:    (Alias)(o),
	})
}

func (o *Billinginvoiceitem) UnmarshalJSON(b []byte) error {
	var BillinginvoiceitemMap map[string]interface{}
	err := json.Unmarshal(b, &BillinginvoiceitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BillinginvoiceitemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Product, ok := BillinginvoiceitemMap["product"].(map[string]interface{}); ok {
		ProductString, _ := json.Marshal(Product)
		json.Unmarshal(ProductString, &o.Product)
	}
	
	if Description, ok := BillinginvoiceitemMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateTransactedString, ok := BillinginvoiceitemMap["dateTransacted"].(string); ok {
		DateTransacted, _ := time.Parse("2006-01-02", dateTransactedString)
		o.DateTransacted = &DateTransacted
	}
	
	if dateStartString, ok := BillinginvoiceitemMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := BillinginvoiceitemMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if Organization, ok := BillinginvoiceitemMap["organization"].(map[string]interface{}); ok {
		OrganizationString, _ := json.Marshal(Organization)
		json.Unmarshal(OrganizationString, &o.Organization)
	}
	
	if Quantity, ok := BillinginvoiceitemMap["quantity"].(float64); ok {
		QuantityInt := int(Quantity)
		o.Quantity = &QuantityInt
	}
	
	if UnitOfMeasure, ok := BillinginvoiceitemMap["unitOfMeasure"].(string); ok {
		o.UnitOfMeasure = &UnitOfMeasure
	}
    
	if Amount, ok := BillinginvoiceitemMap["amount"].(float64); ok {
		AmountFloat32 := float32(Amount)
		o.Amount = &AmountFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Billinginvoiceitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
