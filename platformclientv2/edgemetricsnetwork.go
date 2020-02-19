package platformclientv2
import (
	"encoding/json"
)

// Edgemetricsnetwork
type Edgemetricsnetwork struct { 
	// Ifname - Identifier for the network adapter.
	Ifname *string `json:"ifname,omitempty"`


	// SentBytesPerSec - Number of byes sent per second.
	SentBytesPerSec *int32 `json:"sentBytesPerSec,omitempty"`


	// ReceivedBytesPerSec - Number of byes received per second.
	ReceivedBytesPerSec *int32 `json:"receivedBytesPerSec,omitempty"`


	// BandwidthBitsPerSec - Total bandwidth of the adapter in bits per second.
	BandwidthBitsPerSec *float64 `json:"bandwidthBitsPerSec,omitempty"`


	// UtilizationPct - Percent utilization of the network adapter.
	UtilizationPct *float64 `json:"utilizationPct,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricsnetwork) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
