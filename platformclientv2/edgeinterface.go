package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeinterface
type Edgeinterface struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType
	VarType *string `json:"type,omitempty"`

	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// MacAddress
	MacAddress *string `json:"macAddress,omitempty"`

	// IfName
	IfName *string `json:"ifName,omitempty"`

	// Endpoints
	Endpoints *[]Domainentityref `json:"endpoints,omitempty"`

	// LineTypes
	LineTypes *[]string `json:"lineTypes,omitempty"`

	// AddressFamilyId
	AddressFamilyId *string `json:"addressFamilyId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Edgeinterface) SetField(field string, fieldValue interface{}) {
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

func (o Edgeinterface) MarshalJSON() ([]byte, error) {
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
	type Alias Edgeinterface
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		MacAddress *string `json:"macAddress,omitempty"`
		
		IfName *string `json:"ifName,omitempty"`
		
		Endpoints *[]Domainentityref `json:"endpoints,omitempty"`
		
		LineTypes *[]string `json:"lineTypes,omitempty"`
		
		AddressFamilyId *string `json:"addressFamilyId,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		IpAddress: o.IpAddress,
		
		Name: o.Name,
		
		MacAddress: o.MacAddress,
		
		IfName: o.IfName,
		
		Endpoints: o.Endpoints,
		
		LineTypes: o.LineTypes,
		
		AddressFamilyId: o.AddressFamilyId,
		Alias:    (Alias)(o),
	})
}

func (o *Edgeinterface) UnmarshalJSON(b []byte) error {
	var EdgeinterfaceMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeinterfaceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := EdgeinterfaceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if IpAddress, ok := EdgeinterfaceMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if Name, ok := EdgeinterfaceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if MacAddress, ok := EdgeinterfaceMap["macAddress"].(string); ok {
		o.MacAddress = &MacAddress
	}
    
	if IfName, ok := EdgeinterfaceMap["ifName"].(string); ok {
		o.IfName = &IfName
	}
    
	if Endpoints, ok := EdgeinterfaceMap["endpoints"].([]interface{}); ok {
		EndpointsString, _ := json.Marshal(Endpoints)
		json.Unmarshal(EndpointsString, &o.Endpoints)
	}
	
	if LineTypes, ok := EdgeinterfaceMap["lineTypes"].([]interface{}); ok {
		LineTypesString, _ := json.Marshal(LineTypes)
		json.Unmarshal(LineTypesString, &o.LineTypes)
	}
	
	if AddressFamilyId, ok := EdgeinterfaceMap["addressFamilyId"].(string); ok {
		o.AddressFamilyId = &AddressFamilyId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeinterface) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
