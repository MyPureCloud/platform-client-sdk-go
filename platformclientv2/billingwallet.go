package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Billingwallet
type Billingwallet struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the object.
	Name *string `json:"name,omitempty"`

	// Organizations - A Genesys Cloud Organization and it's related plans.
	Organizations *[]Namedentity `json:"organizations,omitempty"`

	// Product - Represents the details of a product.
	Product *Billingproduct `json:"product,omitempty"`

	// StartingBalance - The initial balance in the wallet.
	StartingBalance *float32 `json:"startingBalance,omitempty"`

	// EndingBalance - The final balance in the wallet after transactions.
	EndingBalance *float32 `json:"endingBalance,omitempty"`

	// BalanceIncrease - Total amount added to the wallet.
	BalanceIncrease *float32 `json:"balanceIncrease,omitempty"`

	// BalanceDecrease - The amount removed from the wallet due to a contract change.
	BalanceDecrease *float32 `json:"balanceDecrease,omitempty"`

	// BalanceConsumption - Total consumption deducted from the wallet.
	BalanceConsumption *float32 `json:"balanceConsumption,omitempty"`

	// BalanceOverage - The amount exceeding a predefined balance threshold.
	BalanceOverage *float32 `json:"balanceOverage,omitempty"`

	// BalanceOverageRate - The rate charged for an overage..
	BalanceOverageRate *float32 `json:"balanceOverageRate,omitempty"`

	// BalanceOverageCharge - The amount to be charged.
	BalanceOverageCharge *float32 `json:"balanceOverageCharge,omitempty"`

	// BalanceOverageCurrency - The currency in which the overage charge is invoiced.
	BalanceOverageCurrency *string `json:"balanceOverageCurrency,omitempty"`

	// UnitOfMeasure - The unit of measure for the wallet.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Billingwallet) SetField(field string, fieldValue interface{}) {
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

func (o Billingwallet) MarshalJSON() ([]byte, error) {
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
	type Alias Billingwallet
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Organizations *[]Namedentity `json:"organizations,omitempty"`
		
		Product *Billingproduct `json:"product,omitempty"`
		
		StartingBalance *float32 `json:"startingBalance,omitempty"`
		
		EndingBalance *float32 `json:"endingBalance,omitempty"`
		
		BalanceIncrease *float32 `json:"balanceIncrease,omitempty"`
		
		BalanceDecrease *float32 `json:"balanceDecrease,omitempty"`
		
		BalanceConsumption *float32 `json:"balanceConsumption,omitempty"`
		
		BalanceOverage *float32 `json:"balanceOverage,omitempty"`
		
		BalanceOverageRate *float32 `json:"balanceOverageRate,omitempty"`
		
		BalanceOverageCharge *float32 `json:"balanceOverageCharge,omitempty"`
		
		BalanceOverageCurrency *string `json:"balanceOverageCurrency,omitempty"`
		
		UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Organizations: o.Organizations,
		
		Product: o.Product,
		
		StartingBalance: o.StartingBalance,
		
		EndingBalance: o.EndingBalance,
		
		BalanceIncrease: o.BalanceIncrease,
		
		BalanceDecrease: o.BalanceDecrease,
		
		BalanceConsumption: o.BalanceConsumption,
		
		BalanceOverage: o.BalanceOverage,
		
		BalanceOverageRate: o.BalanceOverageRate,
		
		BalanceOverageCharge: o.BalanceOverageCharge,
		
		BalanceOverageCurrency: o.BalanceOverageCurrency,
		
		UnitOfMeasure: o.UnitOfMeasure,
		Alias:    (Alias)(o),
	})
}

func (o *Billingwallet) UnmarshalJSON(b []byte) error {
	var BillingwalletMap map[string]interface{}
	err := json.Unmarshal(b, &BillingwalletMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BillingwalletMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := BillingwalletMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Organizations, ok := BillingwalletMap["organizations"].([]interface{}); ok {
		OrganizationsString, _ := json.Marshal(Organizations)
		json.Unmarshal(OrganizationsString, &o.Organizations)
	}
	
	if Product, ok := BillingwalletMap["product"].(map[string]interface{}); ok {
		ProductString, _ := json.Marshal(Product)
		json.Unmarshal(ProductString, &o.Product)
	}
	
	if StartingBalance, ok := BillingwalletMap["startingBalance"].(float64); ok {
		StartingBalanceFloat32 := float32(StartingBalance)
		o.StartingBalance = &StartingBalanceFloat32
	}
    
	if EndingBalance, ok := BillingwalletMap["endingBalance"].(float64); ok {
		EndingBalanceFloat32 := float32(EndingBalance)
		o.EndingBalance = &EndingBalanceFloat32
	}
    
	if BalanceIncrease, ok := BillingwalletMap["balanceIncrease"].(float64); ok {
		BalanceIncreaseFloat32 := float32(BalanceIncrease)
		o.BalanceIncrease = &BalanceIncreaseFloat32
	}
    
	if BalanceDecrease, ok := BillingwalletMap["balanceDecrease"].(float64); ok {
		BalanceDecreaseFloat32 := float32(BalanceDecrease)
		o.BalanceDecrease = &BalanceDecreaseFloat32
	}
    
	if BalanceConsumption, ok := BillingwalletMap["balanceConsumption"].(float64); ok {
		BalanceConsumptionFloat32 := float32(BalanceConsumption)
		o.BalanceConsumption = &BalanceConsumptionFloat32
	}
    
	if BalanceOverage, ok := BillingwalletMap["balanceOverage"].(float64); ok {
		BalanceOverageFloat32 := float32(BalanceOverage)
		o.BalanceOverage = &BalanceOverageFloat32
	}
    
	if BalanceOverageRate, ok := BillingwalletMap["balanceOverageRate"].(float64); ok {
		BalanceOverageRateFloat32 := float32(BalanceOverageRate)
		o.BalanceOverageRate = &BalanceOverageRateFloat32
	}
    
	if BalanceOverageCharge, ok := BillingwalletMap["balanceOverageCharge"].(float64); ok {
		BalanceOverageChargeFloat32 := float32(BalanceOverageCharge)
		o.BalanceOverageCharge = &BalanceOverageChargeFloat32
	}
    
	if BalanceOverageCurrency, ok := BillingwalletMap["balanceOverageCurrency"].(string); ok {
		o.BalanceOverageCurrency = &BalanceOverageCurrency
	}
    
	if UnitOfMeasure, ok := BillingwalletMap["unitOfMeasure"].(string); ok {
		o.UnitOfMeasure = &UnitOfMeasure
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Billingwallet) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
