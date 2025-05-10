package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Billingproduct
type Billingproduct struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Sku - The product associated with the fund.
	Sku *string `json:"sku,omitempty"`

	// Name - The name of the product.
	Name *string `json:"name,omitempty"`

	// UnitOfMeasure - The unit of measure for the product.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Billingproduct) SetField(field string, fieldValue interface{}) {
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

func (o Billingproduct) MarshalJSON() ([]byte, error) {
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
	type Alias Billingproduct
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Sku *string `json:"sku,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Sku: o.Sku,
		
		Name: o.Name,
		
		UnitOfMeasure: o.UnitOfMeasure,
		Alias:    (Alias)(o),
	})
}

func (o *Billingproduct) UnmarshalJSON(b []byte) error {
	var BillingproductMap map[string]interface{}
	err := json.Unmarshal(b, &BillingproductMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BillingproductMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Sku, ok := BillingproductMap["sku"].(string); ok {
		o.Sku = &Sku
	}
    
	if Name, ok := BillingproductMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if UnitOfMeasure, ok := BillingproductMap["unitOfMeasure"].(string); ok {
		o.UnitOfMeasure = &UnitOfMeasure
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Billingproduct) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
