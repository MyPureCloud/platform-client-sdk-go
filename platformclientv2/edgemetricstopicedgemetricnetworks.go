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

func (u *Edgemetricstopicedgemetricnetworks) MarshalJSON() ([]byte, error) {
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
		Ifname: u.Ifname,
		
		SentBytesPerSec: u.SentBytesPerSec,
		
		ReceivedBytesPerSec: u.ReceivedBytesPerSec,
		
		BandwidthBitsPerSec: u.BandwidthBitsPerSec,
		
		UtilizationPct: u.UtilizationPct,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricnetworks) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
