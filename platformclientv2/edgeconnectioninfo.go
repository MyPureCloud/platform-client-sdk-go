package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeconnectioninfo
type Edgeconnectioninfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// InterfaceName - Interface used for the connection on the edge
	InterfaceName *string `json:"interfaceName,omitempty"`

	// InterfaceIpAddress - IP address of the interface
	InterfaceIpAddress *string `json:"interfaceIpAddress,omitempty"`

	// ConnectionErrors - Connection errors
	ConnectionErrors *[]string `json:"connectionErrors,omitempty"`

	// Site
	Site *Addressableentityref `json:"site,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Edgeconnectioninfo) SetField(field string, fieldValue interface{}) {
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

func (o Edgeconnectioninfo) MarshalJSON() ([]byte, error) {
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
	type Alias Edgeconnectioninfo
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		InterfaceName *string `json:"interfaceName,omitempty"`
		
		InterfaceIpAddress *string `json:"interfaceIpAddress,omitempty"`
		
		ConnectionErrors *[]string `json:"connectionErrors,omitempty"`
		
		Site *Addressableentityref `json:"site,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		InterfaceName: o.InterfaceName,
		
		InterfaceIpAddress: o.InterfaceIpAddress,
		
		ConnectionErrors: o.ConnectionErrors,
		
		Site: o.Site,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Edgeconnectioninfo) UnmarshalJSON(b []byte) error {
	var EdgeconnectioninfoMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeconnectioninfoMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgeconnectioninfoMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EdgeconnectioninfoMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if InterfaceName, ok := EdgeconnectioninfoMap["interfaceName"].(string); ok {
		o.InterfaceName = &InterfaceName
	}
    
	if InterfaceIpAddress, ok := EdgeconnectioninfoMap["interfaceIpAddress"].(string); ok {
		o.InterfaceIpAddress = &InterfaceIpAddress
	}
    
	if ConnectionErrors, ok := EdgeconnectioninfoMap["connectionErrors"].([]interface{}); ok {
		ConnectionErrorsString, _ := json.Marshal(ConnectionErrors)
		json.Unmarshal(ConnectionErrorsString, &o.ConnectionErrors)
	}
	
	if Site, ok := EdgeconnectioninfoMap["site"].(map[string]interface{}); ok {
		SiteString, _ := json.Marshal(Site)
		json.Unmarshal(SiteString, &o.Site)
	}
	
	if SelfUri, ok := EdgeconnectioninfoMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeconnectioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
