package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Invoiceaddress
type Invoiceaddress struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// GetdateEffective - The date when the Address became effective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	GetdateEffective *time.Time `json:"getdateEffective,omitempty"`

	// AddressType - The type of address.
	AddressType *string `json:"addressType,omitempty"`

	// CurrencyIsoCode - Contains the ISO code for any currency allowed by the organization.
	CurrencyIsoCode *string `json:"currencyIsoCode,omitempty"`

	// Line1 - The first line of the Address.
	Line1 *string `json:"line1,omitempty"`

	// Line2 - The second line of the Address.
	Line2 *string `json:"line2,omitempty"`

	// Line3 - The third line of the Address.
	Line3 *string `json:"line3,omitempty"`

	// CityName - The city name.
	CityName *string `json:"cityName,omitempty"`

	// PostalCode - The Postal or Zip Code.
	PostalCode *string `json:"postalCode,omitempty"`

	// StateCode - The code that reflects the geographic state for the Address.
	StateCode *string `json:"stateCode,omitempty"`

	// CountryCode - The code representing the country for the Address (ISO_3166).
	CountryCode *string `json:"countryCode,omitempty"`

	// GetcitySubdivision1 - The primary subdivision within a city (e.g., district, neighborhood).
	GetcitySubdivision1 *string `json:"getcitySubdivision1,omitempty"`

	// RegionSubdivision1 - The primary administrative division within a region (e.g., state, province).
	RegionSubdivision1 *string `json:"regionSubdivision1,omitempty"`

	// RegionSubdivision2 - A secondary subdivision within the primary region subdivision (e.g., county, district).
	RegionSubdivision2 *string `json:"regionSubdivision2,omitempty"`

	// Country - Stores the name of the country in which the address is situated.
	Country *string `json:"country,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Invoiceaddress) SetField(field string, fieldValue interface{}) {
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

func (o Invoiceaddress) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "GetdateEffective", }

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
	type Alias Invoiceaddress
	
	GetdateEffective := new(string)
	if o.GetdateEffective != nil {
		*GetdateEffective = timeutil.Strftime(o.GetdateEffective, "%Y-%m-%d")
	} else {
		GetdateEffective = nil
	}
	
	return json.Marshal(&struct { 
		GetdateEffective *string `json:"getdateEffective,omitempty"`
		
		AddressType *string `json:"addressType,omitempty"`
		
		CurrencyIsoCode *string `json:"currencyIsoCode,omitempty"`
		
		Line1 *string `json:"line1,omitempty"`
		
		Line2 *string `json:"line2,omitempty"`
		
		Line3 *string `json:"line3,omitempty"`
		
		CityName *string `json:"cityName,omitempty"`
		
		PostalCode *string `json:"postalCode,omitempty"`
		
		StateCode *string `json:"stateCode,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		GetcitySubdivision1 *string `json:"getcitySubdivision1,omitempty"`
		
		RegionSubdivision1 *string `json:"regionSubdivision1,omitempty"`
		
		RegionSubdivision2 *string `json:"regionSubdivision2,omitempty"`
		
		Country *string `json:"country,omitempty"`
		Alias
	}{ 
		GetdateEffective: GetdateEffective,
		
		AddressType: o.AddressType,
		
		CurrencyIsoCode: o.CurrencyIsoCode,
		
		Line1: o.Line1,
		
		Line2: o.Line2,
		
		Line3: o.Line3,
		
		CityName: o.CityName,
		
		PostalCode: o.PostalCode,
		
		StateCode: o.StateCode,
		
		CountryCode: o.CountryCode,
		
		GetcitySubdivision1: o.GetcitySubdivision1,
		
		RegionSubdivision1: o.RegionSubdivision1,
		
		RegionSubdivision2: o.RegionSubdivision2,
		
		Country: o.Country,
		Alias:    (Alias)(o),
	})
}

func (o *Invoiceaddress) UnmarshalJSON(b []byte) error {
	var InvoiceaddressMap map[string]interface{}
	err := json.Unmarshal(b, &InvoiceaddressMap)
	if err != nil {
		return err
	}
	
	if getdateEffectiveString, ok := InvoiceaddressMap["getdateEffective"].(string); ok {
		GetdateEffective, _ := time.Parse("2006-01-02", getdateEffectiveString)
		o.GetdateEffective = &GetdateEffective
	}
	
	if AddressType, ok := InvoiceaddressMap["addressType"].(string); ok {
		o.AddressType = &AddressType
	}
    
	if CurrencyIsoCode, ok := InvoiceaddressMap["currencyIsoCode"].(string); ok {
		o.CurrencyIsoCode = &CurrencyIsoCode
	}
    
	if Line1, ok := InvoiceaddressMap["line1"].(string); ok {
		o.Line1 = &Line1
	}
    
	if Line2, ok := InvoiceaddressMap["line2"].(string); ok {
		o.Line2 = &Line2
	}
    
	if Line3, ok := InvoiceaddressMap["line3"].(string); ok {
		o.Line3 = &Line3
	}
    
	if CityName, ok := InvoiceaddressMap["cityName"].(string); ok {
		o.CityName = &CityName
	}
    
	if PostalCode, ok := InvoiceaddressMap["postalCode"].(string); ok {
		o.PostalCode = &PostalCode
	}
    
	if StateCode, ok := InvoiceaddressMap["stateCode"].(string); ok {
		o.StateCode = &StateCode
	}
    
	if CountryCode, ok := InvoiceaddressMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if GetcitySubdivision1, ok := InvoiceaddressMap["getcitySubdivision1"].(string); ok {
		o.GetcitySubdivision1 = &GetcitySubdivision1
	}
    
	if RegionSubdivision1, ok := InvoiceaddressMap["regionSubdivision1"].(string); ok {
		o.RegionSubdivision1 = &RegionSubdivision1
	}
    
	if RegionSubdivision2, ok := InvoiceaddressMap["regionSubdivision2"].(string); ok {
		o.RegionSubdivision2 = &RegionSubdivision2
	}
    
	if Country, ok := InvoiceaddressMap["country"].(string); ok {
		o.Country = &Country
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Invoiceaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
