package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricnetworks) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
