package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Billingaisubitem
type Billingaisubitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ChargeName - Charge-level name of the member charge. Omitted when the source charge has no name.
	ChargeName *string `json:"chargeName,omitempty"`

	// ActualQuantity - Usage in the child product's original unit.
	ActualQuantity *float32 `json:"actualQuantity,omitempty"`

	// UnitOfMeasure - Original unit of measure for the child product usage.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`

	// ConvertedTokenCount - Child usage converted to parent-bundle tokens.
	ConvertedTokenCount *float32 `json:"convertedTokenCount,omitempty"`

	// TokenConversionRate - Bundle tokens per one original unit.
	TokenConversionRate *float32 `json:"tokenConversionRate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Billingaisubitem) SetField(field string, fieldValue interface{}) {
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

func (o Billingaisubitem) MarshalJSON() ([]byte, error) {
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
	type Alias Billingaisubitem
	
	return json.Marshal(&struct { 
		ChargeName *string `json:"chargeName,omitempty"`
		
		ActualQuantity *float32 `json:"actualQuantity,omitempty"`
		
		UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
		
		ConvertedTokenCount *float32 `json:"convertedTokenCount,omitempty"`
		
		TokenConversionRate *float32 `json:"tokenConversionRate,omitempty"`
		Alias
	}{ 
		ChargeName: o.ChargeName,
		
		ActualQuantity: o.ActualQuantity,
		
		UnitOfMeasure: o.UnitOfMeasure,
		
		ConvertedTokenCount: o.ConvertedTokenCount,
		
		TokenConversionRate: o.TokenConversionRate,
		Alias:    (Alias)(o),
	})
}

func (o *Billingaisubitem) UnmarshalJSON(b []byte) error {
	var BillingaisubitemMap map[string]interface{}
	err := json.Unmarshal(b, &BillingaisubitemMap)
	if err != nil {
		return err
	}
	
	if ChargeName, ok := BillingaisubitemMap["chargeName"].(string); ok {
		o.ChargeName = &ChargeName
	}
    
	if ActualQuantity, ok := BillingaisubitemMap["actualQuantity"].(float64); ok {
		ActualQuantityFloat32 := float32(ActualQuantity)
		o.ActualQuantity = &ActualQuantityFloat32
	}
    
	if UnitOfMeasure, ok := BillingaisubitemMap["unitOfMeasure"].(string); ok {
		o.UnitOfMeasure = &UnitOfMeasure
	}
    
	if ConvertedTokenCount, ok := BillingaisubitemMap["convertedTokenCount"].(float64); ok {
		ConvertedTokenCountFloat32 := float32(ConvertedTokenCount)
		o.ConvertedTokenCount = &ConvertedTokenCountFloat32
	}
    
	if TokenConversionRate, ok := BillingaisubitemMap["tokenConversionRate"].(float64); ok {
		TokenConversionRateFloat32 := float32(TokenConversionRate)
		o.TokenConversionRate = &TokenConversionRateFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Billingaisubitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
