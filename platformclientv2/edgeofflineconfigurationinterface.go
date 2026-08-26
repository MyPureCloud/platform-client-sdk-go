package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeofflineconfigurationinterface
type Edgeofflineconfigurationinterface struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Routes - The list of routes assigned to this interface.
	Routes *[]Domainnetworkroute `json:"routes,omitempty"`

	// Addresses - The list of IP addresses on this interface.  Priority of dns addresses are based on order in the list.
	Addresses *[]Domainnetworkaddress `json:"addresses,omitempty"`

	// Ipv4Capabilities - IPv4 interface settings.
	Ipv4Capabilities *Domaincapabilities `json:"ipv4Capabilities,omitempty"`

	// Ipv6Capabilities - IPv6 interface settings.
	Ipv6Capabilities *Domaincapabilities `json:"ipv6Capabilities,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Edgeofflineconfigurationinterface) SetField(field string, fieldValue interface{}) {
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

func (o Edgeofflineconfigurationinterface) MarshalJSON() ([]byte, error) {
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
	type Alias Edgeofflineconfigurationinterface
	
	return json.Marshal(&struct { 
		Routes *[]Domainnetworkroute `json:"routes,omitempty"`
		
		Addresses *[]Domainnetworkaddress `json:"addresses,omitempty"`
		
		Ipv4Capabilities *Domaincapabilities `json:"ipv4Capabilities,omitempty"`
		
		Ipv6Capabilities *Domaincapabilities `json:"ipv6Capabilities,omitempty"`
		Alias
	}{ 
		Routes: o.Routes,
		
		Addresses: o.Addresses,
		
		Ipv4Capabilities: o.Ipv4Capabilities,
		
		Ipv6Capabilities: o.Ipv6Capabilities,
		Alias:    (Alias)(o),
	})
}

func (o *Edgeofflineconfigurationinterface) UnmarshalJSON(b []byte) error {
	var EdgeofflineconfigurationinterfaceMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeofflineconfigurationinterfaceMap)
	if err != nil {
		return err
	}
	
	if Routes, ok := EdgeofflineconfigurationinterfaceMap["routes"].([]interface{}); ok {
		RoutesString, _ := json.Marshal(Routes)
		json.Unmarshal(RoutesString, &o.Routes)
	}
	
	if Addresses, ok := EdgeofflineconfigurationinterfaceMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if Ipv4Capabilities, ok := EdgeofflineconfigurationinterfaceMap["ipv4Capabilities"].(map[string]interface{}); ok {
		Ipv4CapabilitiesString, _ := json.Marshal(Ipv4Capabilities)
		json.Unmarshal(Ipv4CapabilitiesString, &o.Ipv4Capabilities)
	}
	
	if Ipv6Capabilities, ok := EdgeofflineconfigurationinterfaceMap["ipv6Capabilities"].(map[string]interface{}); ok {
		Ipv6CapabilitiesString, _ := json.Marshal(Ipv6Capabilities)
		json.Unmarshal(Ipv6CapabilitiesString, &o.Ipv6Capabilities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeofflineconfigurationinterface) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
