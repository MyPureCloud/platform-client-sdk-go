package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Billingcharge
type Billingcharge struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Product - Represents the details of a product.
	Product *Billingproduct `json:"product,omitempty"`

	// Organizations - List of plans within the organization.
	Organizations *[]Namedentity `json:"organizations,omitempty"`

	// PrepaidQuantity - The quantity of usage that is prepaid.
	PrepaidQuantity *int `json:"prepaidQuantity,omitempty"`

	// FairuseQuantity - The quantity of usage allowed under fair use policies.
	FairuseQuantity *int `json:"fairuseQuantity,omitempty"`

	// ActualQuantity - The actual quantity of usage.
	ActualQuantity *int `json:"actualQuantity,omitempty"`

	// OverageQuantity - The quantity of usage that exceeds prepaid or fair use limits.
	OverageQuantity *int `json:"overageQuantity,omitempty"`

	// OverageRate - The rate charged per unit of overage.
	OverageRate *float32 `json:"overageRate,omitempty"`

	// OverageCharge - The total charge for overage usage.
	OverageCharge *float32 `json:"overageCharge,omitempty"`

	// OverageCurrency - The currency in which the overage charge is billed.
	OverageCurrency *string `json:"overageCurrency,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Billingcharge) SetField(field string, fieldValue interface{}) {
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

func (o Billingcharge) MarshalJSON() ([]byte, error) {
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
	type Alias Billingcharge
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Product *Billingproduct `json:"product,omitempty"`
		
		Organizations *[]Namedentity `json:"organizations,omitempty"`
		
		PrepaidQuantity *int `json:"prepaidQuantity,omitempty"`
		
		FairuseQuantity *int `json:"fairuseQuantity,omitempty"`
		
		ActualQuantity *int `json:"actualQuantity,omitempty"`
		
		OverageQuantity *int `json:"overageQuantity,omitempty"`
		
		OverageRate *float32 `json:"overageRate,omitempty"`
		
		OverageCharge *float32 `json:"overageCharge,omitempty"`
		
		OverageCurrency *string `json:"overageCurrency,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Product: o.Product,
		
		Organizations: o.Organizations,
		
		PrepaidQuantity: o.PrepaidQuantity,
		
		FairuseQuantity: o.FairuseQuantity,
		
		ActualQuantity: o.ActualQuantity,
		
		OverageQuantity: o.OverageQuantity,
		
		OverageRate: o.OverageRate,
		
		OverageCharge: o.OverageCharge,
		
		OverageCurrency: o.OverageCurrency,
		Alias:    (Alias)(o),
	})
}

func (o *Billingcharge) UnmarshalJSON(b []byte) error {
	var BillingchargeMap map[string]interface{}
	err := json.Unmarshal(b, &BillingchargeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BillingchargeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Product, ok := BillingchargeMap["product"].(map[string]interface{}); ok {
		ProductString, _ := json.Marshal(Product)
		json.Unmarshal(ProductString, &o.Product)
	}
	
	if Organizations, ok := BillingchargeMap["organizations"].([]interface{}); ok {
		OrganizationsString, _ := json.Marshal(Organizations)
		json.Unmarshal(OrganizationsString, &o.Organizations)
	}
	
	if PrepaidQuantity, ok := BillingchargeMap["prepaidQuantity"].(float64); ok {
		PrepaidQuantityInt := int(PrepaidQuantity)
		o.PrepaidQuantity = &PrepaidQuantityInt
	}
	
	if FairuseQuantity, ok := BillingchargeMap["fairuseQuantity"].(float64); ok {
		FairuseQuantityInt := int(FairuseQuantity)
		o.FairuseQuantity = &FairuseQuantityInt
	}
	
	if ActualQuantity, ok := BillingchargeMap["actualQuantity"].(float64); ok {
		ActualQuantityInt := int(ActualQuantity)
		o.ActualQuantity = &ActualQuantityInt
	}
	
	if OverageQuantity, ok := BillingchargeMap["overageQuantity"].(float64); ok {
		OverageQuantityInt := int(OverageQuantity)
		o.OverageQuantity = &OverageQuantityInt
	}
	
	if OverageRate, ok := BillingchargeMap["overageRate"].(float64); ok {
		OverageRateFloat32 := float32(OverageRate)
		o.OverageRate = &OverageRateFloat32
	}
    
	if OverageCharge, ok := BillingchargeMap["overageCharge"].(float64); ok {
		OverageChargeFloat32 := float32(OverageCharge)
		o.OverageCharge = &OverageChargeFloat32
	}
    
	if OverageCurrency, ok := BillingchargeMap["overageCurrency"].(string); ok {
		o.OverageCurrency = &OverageCurrency
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Billingcharge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
