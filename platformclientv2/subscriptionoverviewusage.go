package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Subscriptionoverviewusage
type Subscriptionoverviewusage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - Product charge name
	Name *string `json:"name,omitempty"`

	// PartNumber - Product part number
	PartNumber *string `json:"partNumber,omitempty"`

	// Grouping - UI grouping key
	Grouping *string `json:"grouping,omitempty"`

	// UnitOfMeasureType - UI unit of measure
	UnitOfMeasureType *string `json:"unitOfMeasureType,omitempty"`

	// UsageQuantity - Usage count for specified period
	UsageQuantity *string `json:"usageQuantity,omitempty"`

	// OveragePrice - Price for usage / overage charge
	OveragePrice *string `json:"overagePrice,omitempty"`

	// PrepayQuantity - Items prepaid for specified period
	PrepayQuantity *string `json:"prepayQuantity,omitempty"`

	// PrepayPrice - Price for prepay charge
	PrepayPrice *string `json:"prepayPrice,omitempty"`

	// UsageNotes - Notes about the usage/charge item
	UsageNotes *string `json:"usageNotes,omitempty"`

	// IsCancellable - Indicates whether the item is cancellable
	IsCancellable *bool `json:"isCancellable,omitempty"`

	// BundleQuantity - Quantity multiplier for this charge
	BundleQuantity *string `json:"bundleQuantity,omitempty"`

	// IsThirdParty - A charge from a third party entity
	IsThirdParty *bool `json:"isThirdParty,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Subscriptionoverviewusage) SetField(field string, fieldValue interface{}) {
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

func (o Subscriptionoverviewusage) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Subscriptionoverviewusage
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		PartNumber *string `json:"partNumber,omitempty"`
		
		Grouping *string `json:"grouping,omitempty"`
		
		UnitOfMeasureType *string `json:"unitOfMeasureType,omitempty"`
		
		UsageQuantity *string `json:"usageQuantity,omitempty"`
		
		OveragePrice *string `json:"overagePrice,omitempty"`
		
		PrepayQuantity *string `json:"prepayQuantity,omitempty"`
		
		PrepayPrice *string `json:"prepayPrice,omitempty"`
		
		UsageNotes *string `json:"usageNotes,omitempty"`
		
		IsCancellable *bool `json:"isCancellable,omitempty"`
		
		BundleQuantity *string `json:"bundleQuantity,omitempty"`
		
		IsThirdParty *bool `json:"isThirdParty,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		PartNumber: o.PartNumber,
		
		Grouping: o.Grouping,
		
		UnitOfMeasureType: o.UnitOfMeasureType,
		
		UsageQuantity: o.UsageQuantity,
		
		OveragePrice: o.OveragePrice,
		
		PrepayQuantity: o.PrepayQuantity,
		
		PrepayPrice: o.PrepayPrice,
		
		UsageNotes: o.UsageNotes,
		
		IsCancellable: o.IsCancellable,
		
		BundleQuantity: o.BundleQuantity,
		
		IsThirdParty: o.IsThirdParty,
		Alias:    (Alias)(o),
	})
}

func (o *Subscriptionoverviewusage) UnmarshalJSON(b []byte) error {
	var SubscriptionoverviewusageMap map[string]interface{}
	err := json.Unmarshal(b, &SubscriptionoverviewusageMap)
	if err != nil {
		return err
	}
	
	if Name, ok := SubscriptionoverviewusageMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if PartNumber, ok := SubscriptionoverviewusageMap["partNumber"].(string); ok {
		o.PartNumber = &PartNumber
	}
    
	if Grouping, ok := SubscriptionoverviewusageMap["grouping"].(string); ok {
		o.Grouping = &Grouping
	}
    
	if UnitOfMeasureType, ok := SubscriptionoverviewusageMap["unitOfMeasureType"].(string); ok {
		o.UnitOfMeasureType = &UnitOfMeasureType
	}
    
	if UsageQuantity, ok := SubscriptionoverviewusageMap["usageQuantity"].(string); ok {
		o.UsageQuantity = &UsageQuantity
	}
    
	if OveragePrice, ok := SubscriptionoverviewusageMap["overagePrice"].(string); ok {
		o.OveragePrice = &OveragePrice
	}
    
	if PrepayQuantity, ok := SubscriptionoverviewusageMap["prepayQuantity"].(string); ok {
		o.PrepayQuantity = &PrepayQuantity
	}
    
	if PrepayPrice, ok := SubscriptionoverviewusageMap["prepayPrice"].(string); ok {
		o.PrepayPrice = &PrepayPrice
	}
    
	if UsageNotes, ok := SubscriptionoverviewusageMap["usageNotes"].(string); ok {
		o.UsageNotes = &UsageNotes
	}
    
	if IsCancellable, ok := SubscriptionoverviewusageMap["isCancellable"].(bool); ok {
		o.IsCancellable = &IsCancellable
	}
    
	if BundleQuantity, ok := SubscriptionoverviewusageMap["bundleQuantity"].(string); ok {
		o.BundleQuantity = &BundleQuantity
	}
    
	if IsThirdParty, ok := SubscriptionoverviewusageMap["isThirdParty"].(bool); ok {
		o.IsThirdParty = &IsThirdParty
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Subscriptionoverviewusage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
