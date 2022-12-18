package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Networkconnectivity
type Networkconnectivity struct { 
	// Carrier - The name of the network carrier.
	Carrier *string `json:"carrier,omitempty"`


	// BluetoothEnabled - Whether Bluetooth is enabled.
	BluetoothEnabled *bool `json:"bluetoothEnabled,omitempty"`


	// CellularEnabled - Whether Cellular is enabled.
	CellularEnabled *bool `json:"cellularEnabled,omitempty"`


	// WifiEnabled - Whether Wi-Fi is enabled.
	WifiEnabled *bool `json:"wifiEnabled,omitempty"`

}

func (o *Networkconnectivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Networkconnectivity
	
	return json.Marshal(&struct { 
		Carrier *string `json:"carrier,omitempty"`
		
		BluetoothEnabled *bool `json:"bluetoothEnabled,omitempty"`
		
		CellularEnabled *bool `json:"cellularEnabled,omitempty"`
		
		WifiEnabled *bool `json:"wifiEnabled,omitempty"`
		*Alias
	}{ 
		Carrier: o.Carrier,
		
		BluetoothEnabled: o.BluetoothEnabled,
		
		CellularEnabled: o.CellularEnabled,
		
		WifiEnabled: o.WifiEnabled,
		Alias:    (*Alias)(o),
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
