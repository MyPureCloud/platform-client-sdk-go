package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Billingcontract
type Billingcontract struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// ExternalNumber - Unique external number.
	ExternalNumber *string `json:"externalNumber,omitempty"`

	// Status - The status of the contract.
	Status *string `json:"status,omitempty"`

	// CommercialModel - The type of commercial model.
	CommercialModel *string `json:"commercialModel,omitempty"`

	// PurchaseOrderNumbers - List of po numbers periods for this contract.
	PurchaseOrderNumbers *[]string `json:"purchaseOrderNumbers,omitempty"`

	// BillToCustomer - The bill to customer.
	BillToCustomer *Customer `json:"billToCustomer,omitempty"`

	// SoldToCustomer - The sold to customer.
	SoldToCustomer *Customer `json:"soldToCustomer,omitempty"`

	// EndCustomer - The end customer.
	EndCustomer *Customer `json:"endCustomer,omitempty"`

	// DateStart - The start date of the contract. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - The end date of the contract. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// DateRampStart - the date when the first revenue or quantity in a ramped deal begins. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateRampStart *time.Time `json:"dateRampStart,omitempty"`

	// DateRampEnd - the date when the first revenue or quantity in a ramped deal ends. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateRampEnd *time.Time `json:"dateRampEnd,omitempty"`

	// BillingPeriods - List of billing periods for the contract.
	BillingPeriods *[]Billingcontractperiod `json:"billingPeriods,omitempty"`

	// Plans - The collection of prices for the related organizations.
	Plans *[]Billingplan `json:"plans,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Billingcontract) SetField(field string, fieldValue interface{}) {
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

func (o Billingcontract) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateStart","DateEnd","DateRampStart","DateRampEnd", }

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
	type Alias Billingcontract
	
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
	
	DateRampStart := new(string)
	if o.DateRampStart != nil {
		*DateRampStart = timeutil.Strftime(o.DateRampStart, "%Y-%m-%d")
	} else {
		DateRampStart = nil
	}
	
	DateRampEnd := new(string)
	if o.DateRampEnd != nil {
		*DateRampEnd = timeutil.Strftime(o.DateRampEnd, "%Y-%m-%d")
	} else {
		DateRampEnd = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ExternalNumber *string `json:"externalNumber,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		CommercialModel *string `json:"commercialModel,omitempty"`
		
		PurchaseOrderNumbers *[]string `json:"purchaseOrderNumbers,omitempty"`
		
		BillToCustomer *Customer `json:"billToCustomer,omitempty"`
		
		SoldToCustomer *Customer `json:"soldToCustomer,omitempty"`
		
		EndCustomer *Customer `json:"endCustomer,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		DateRampStart *string `json:"dateRampStart,omitempty"`
		
		DateRampEnd *string `json:"dateRampEnd,omitempty"`
		
		BillingPeriods *[]Billingcontractperiod `json:"billingPeriods,omitempty"`
		
		Plans *[]Billingplan `json:"plans,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ExternalNumber: o.ExternalNumber,
		
		Status: o.Status,
		
		CommercialModel: o.CommercialModel,
		
		PurchaseOrderNumbers: o.PurchaseOrderNumbers,
		
		BillToCustomer: o.BillToCustomer,
		
		SoldToCustomer: o.SoldToCustomer,
		
		EndCustomer: o.EndCustomer,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		DateRampStart: DateRampStart,
		
		DateRampEnd: DateRampEnd,
		
		BillingPeriods: o.BillingPeriods,
		
		Plans: o.Plans,
		Alias:    (Alias)(o),
	})
}

func (o *Billingcontract) UnmarshalJSON(b []byte) error {
	var BillingcontractMap map[string]interface{}
	err := json.Unmarshal(b, &BillingcontractMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BillingcontractMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ExternalNumber, ok := BillingcontractMap["externalNumber"].(string); ok {
		o.ExternalNumber = &ExternalNumber
	}
    
	if Status, ok := BillingcontractMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if CommercialModel, ok := BillingcontractMap["commercialModel"].(string); ok {
		o.CommercialModel = &CommercialModel
	}
    
	if PurchaseOrderNumbers, ok := BillingcontractMap["purchaseOrderNumbers"].([]interface{}); ok {
		PurchaseOrderNumbersString, _ := json.Marshal(PurchaseOrderNumbers)
		json.Unmarshal(PurchaseOrderNumbersString, &o.PurchaseOrderNumbers)
	}
	
	if BillToCustomer, ok := BillingcontractMap["billToCustomer"].(map[string]interface{}); ok {
		BillToCustomerString, _ := json.Marshal(BillToCustomer)
		json.Unmarshal(BillToCustomerString, &o.BillToCustomer)
	}
	
	if SoldToCustomer, ok := BillingcontractMap["soldToCustomer"].(map[string]interface{}); ok {
		SoldToCustomerString, _ := json.Marshal(SoldToCustomer)
		json.Unmarshal(SoldToCustomerString, &o.SoldToCustomer)
	}
	
	if EndCustomer, ok := BillingcontractMap["endCustomer"].(map[string]interface{}); ok {
		EndCustomerString, _ := json.Marshal(EndCustomer)
		json.Unmarshal(EndCustomerString, &o.EndCustomer)
	}
	
	if dateStartString, ok := BillingcontractMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := BillingcontractMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if dateRampStartString, ok := BillingcontractMap["dateRampStart"].(string); ok {
		DateRampStart, _ := time.Parse("2006-01-02", dateRampStartString)
		o.DateRampStart = &DateRampStart
	}
	
	if dateRampEndString, ok := BillingcontractMap["dateRampEnd"].(string); ok {
		DateRampEnd, _ := time.Parse("2006-01-02", dateRampEndString)
		o.DateRampEnd = &DateRampEnd
	}
	
	if BillingPeriods, ok := BillingcontractMap["billingPeriods"].([]interface{}); ok {
		BillingPeriodsString, _ := json.Marshal(BillingPeriods)
		json.Unmarshal(BillingPeriodsString, &o.BillingPeriods)
	}
	
	if Plans, ok := BillingcontractMap["plans"].([]interface{}); ok {
		PlansString, _ := json.Marshal(Plans)
		json.Unmarshal(PlansString, &o.Plans)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Billingcontract) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
