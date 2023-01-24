package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainnetworkaddress
type Domainnetworkaddress struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of address.
	VarType *string `json:"type,omitempty"`

	// Address - An IPv4 or IPv6 IP address. When specifying an address of type \"ip\", use CIDR format for the subnet mask.
	Address *string `json:"address,omitempty"`

	// Persistent - True if this address will persist on Edge restart.  Addresses assigned by DHCP will be returned as false.
	Persistent *bool `json:"persistent,omitempty"`

	// Family - The address family for this address.
	Family *int `json:"family,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Domainnetworkaddress) SetField(field string, fieldValue interface{}) {
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

func (o Domainnetworkaddress) MarshalJSON() ([]byte, error) {
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
	type Alias Domainnetworkaddress
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		Persistent *bool `json:"persistent,omitempty"`
		
		Family *int `json:"family,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Address: o.Address,
		
		Persistent: o.Persistent,
		
		Family: o.Family,
		Alias:    (Alias)(o),
	})
}

func (o *Domainnetworkaddress) UnmarshalJSON(b []byte) error {
	var DomainnetworkaddressMap map[string]interface{}
	err := json.Unmarshal(b, &DomainnetworkaddressMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DomainnetworkaddressMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Address, ok := DomainnetworkaddressMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if Persistent, ok := DomainnetworkaddressMap["persistent"].(bool); ok {
		o.Persistent = &Persistent
	}
    
	if Family, ok := DomainnetworkaddressMap["family"].(float64); ok {
		FamilyInt := int(Family)
		o.Family = &FamilyInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainnetworkaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
