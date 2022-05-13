package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricsnetwork
type Edgemetricsnetwork struct { 
	// Ifname - Identifier for the network adapter.
	Ifname *string `json:"ifname,omitempty"`


	// SentBytesPerSec - Number of byes sent per second.
	SentBytesPerSec *int `json:"sentBytesPerSec,omitempty"`


	// ReceivedBytesPerSec - Number of byes received per second.
	ReceivedBytesPerSec *int `json:"receivedBytesPerSec,omitempty"`


	// BandwidthBitsPerSec - Total bandwidth of the adapter in bits per second.
	BandwidthBitsPerSec *float64 `json:"bandwidthBitsPerSec,omitempty"`


	// UtilizationPct - Percent utilization of the network adapter.
	UtilizationPct *float64 `json:"utilizationPct,omitempty"`

}

func (o *Edgemetricsnetwork) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricsnetwork
	
	return json.Marshal(&struct { 
		Ifname *string `json:"ifname,omitempty"`
		
		SentBytesPerSec *int `json:"sentBytesPerSec,omitempty"`
		
		ReceivedBytesPerSec *int `json:"receivedBytesPerSec,omitempty"`
		
		BandwidthBitsPerSec *float64 `json:"bandwidthBitsPerSec,omitempty"`
		
		UtilizationPct *float64 `json:"utilizationPct,omitempty"`
		*Alias
	}{ 
		Ifname: o.Ifname,
		
		SentBytesPerSec: o.SentBytesPerSec,
		
		ReceivedBytesPerSec: o.ReceivedBytesPerSec,
		
		BandwidthBitsPerSec: o.BandwidthBitsPerSec,
		
		UtilizationPct: o.UtilizationPct,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetricsnetwork) UnmarshalJSON(b []byte) error {
	var EdgemetricsnetworkMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricsnetworkMap)
	if err != nil {
		return err
	}
	
	if Ifname, ok := EdgemetricsnetworkMap["ifname"].(string); ok {
		o.Ifname = &Ifname
	}
    
	if SentBytesPerSec, ok := EdgemetricsnetworkMap["sentBytesPerSec"].(float64); ok {
		SentBytesPerSecInt := int(SentBytesPerSec)
		o.SentBytesPerSec = &SentBytesPerSecInt
	}
	
	if ReceivedBytesPerSec, ok := EdgemetricsnetworkMap["receivedBytesPerSec"].(float64); ok {
		ReceivedBytesPerSecInt := int(ReceivedBytesPerSec)
		o.ReceivedBytesPerSec = &ReceivedBytesPerSecInt
	}
	
	if BandwidthBitsPerSec, ok := EdgemetricsnetworkMap["bandwidthBitsPerSec"].(float64); ok {
		o.BandwidthBitsPerSec = &BandwidthBitsPerSec
	}
    
	if UtilizationPct, ok := EdgemetricsnetworkMap["utilizationPct"].(float64); ok {
		o.UtilizationPct = &UtilizationPct
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricsnetwork) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
