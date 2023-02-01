package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Networkconnectivity
type Networkconnectivity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Carrier - The name of the network carrier.
	Carrier *string `json:"carrier,omitempty"`

	// BluetoothEnabled - Whether Bluetooth is enabled.
	BluetoothEnabled *bool `json:"bluetoothEnabled,omitempty"`

	// CellularEnabled - Whether Cellular is enabled.
	CellularEnabled *bool `json:"cellularEnabled,omitempty"`

	// WifiEnabled - Whether Wi-Fi is enabled.
	WifiEnabled *bool `json:"wifiEnabled,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Networkconnectivity) SetField(field string, fieldValue interface{}) {
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

func (o Networkconnectivity) MarshalJSON() ([]byte, error) {
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
	type Alias Networkconnectivity
	
	return json.Marshal(&struct { 
		Carrier *string `json:"carrier,omitempty"`
		
		BluetoothEnabled *bool `json:"bluetoothEnabled,omitempty"`
		
		CellularEnabled *bool `json:"cellularEnabled,omitempty"`
		
		WifiEnabled *bool `json:"wifiEnabled,omitempty"`
		Alias
	}{ 
		Carrier: o.Carrier,
		
		BluetoothEnabled: o.BluetoothEnabled,
		
		CellularEnabled: o.CellularEnabled,
		
		WifiEnabled: o.WifiEnabled,
		Alias:    (Alias)(o),
	})
}

func (o *Networkconnectivity) UnmarshalJSON(b []byte) error {
	var NetworkconnectivityMap map[string]interface{}
	err := json.Unmarshal(b, &NetworkconnectivityMap)
	if err != nil {
		return err
	}
	
	if Carrier, ok := NetworkconnectivityMap["carrier"].(string); ok {
		o.Carrier = &Carrier
	}
    
	if BluetoothEnabled, ok := NetworkconnectivityMap["bluetoothEnabled"].(bool); ok {
		o.BluetoothEnabled = &BluetoothEnabled
	}
    
	if CellularEnabled, ok := NetworkconnectivityMap["cellularEnabled"].(bool); ok {
		o.CellularEnabled = &CellularEnabled
	}
    
	if WifiEnabled, ok := NetworkconnectivityMap["wifiEnabled"].(bool); ok {
		o.WifiEnabled = &WifiEnabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Networkconnectivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
