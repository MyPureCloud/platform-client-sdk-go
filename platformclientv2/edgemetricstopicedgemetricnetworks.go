package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopicedgemetricnetworks
type Edgemetricstopicedgemetricnetworks struct { 
	// Ifname
	Ifname *string `json:"ifname,omitempty"`


	// SentBytesPerSec
	SentBytesPerSec *int `json:"sentBytesPerSec,omitempty"`


	// ReceivedBytesPerSec
	ReceivedBytesPerSec *int `json:"receivedBytesPerSec,omitempty"`


	// BandwidthBitsPerSec
	BandwidthBitsPerSec *int `json:"bandwidthBitsPerSec,omitempty"`


	// UtilizationPct
	UtilizationPct *float32 `json:"utilizationPct,omitempty"`

}

func (o *Edgemetricstopicedgemetricnetworks) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetricnetworks
	
	return json.Marshal(&struct { 
		Ifname *string `json:"ifname,omitempty"`
		
		SentBytesPerSec *int `json:"sentBytesPerSec,omitempty"`
		
		ReceivedBytesPerSec *int `json:"receivedBytesPerSec,omitempty"`
		
		BandwidthBitsPerSec *int `json:"bandwidthBitsPerSec,omitempty"`
		
		UtilizationPct *float32 `json:"utilizationPct,omitempty"`
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

func (o *Edgemetricstopicedgemetricnetworks) UnmarshalJSON(b []byte) error {
	var EdgemetricstopicedgemetricnetworksMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricstopicedgemetricnetworksMap)
	if err != nil {
		return err
	}
	
	if Ifname, ok := EdgemetricstopicedgemetricnetworksMap["ifname"].(string); ok {
		o.Ifname = &Ifname
	}
    
	if SentBytesPerSec, ok := EdgemetricstopicedgemetricnetworksMap["sentBytesPerSec"].(float64); ok {
		SentBytesPerSecInt := int(SentBytesPerSec)
		o.SentBytesPerSec = &SentBytesPerSecInt
	}
	
	if ReceivedBytesPerSec, ok := EdgemetricstopicedgemetricnetworksMap["receivedBytesPerSec"].(float64); ok {
		ReceivedBytesPerSecInt := int(ReceivedBytesPerSec)
		o.ReceivedBytesPerSec = &ReceivedBytesPerSecInt
	}
	
	if BandwidthBitsPerSec, ok := EdgemetricstopicedgemetricnetworksMap["bandwidthBitsPerSec"].(float64); ok {
		BandwidthBitsPerSecInt := int(BandwidthBitsPerSec)
		o.BandwidthBitsPerSec = &BandwidthBitsPerSecInt
	}
	
	if UtilizationPct, ok := EdgemetricstopicedgemetricnetworksMap["utilizationPct"].(float64); ok {
		UtilizationPctFloat32 := float32(UtilizationPct)
		o.UtilizationPct = &UtilizationPctFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricnetworks) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
