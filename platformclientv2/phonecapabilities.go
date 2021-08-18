package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonecapabilities
type Phonecapabilities struct { 
	// Provisions
	Provisions *bool `json:"provisions,omitempty"`


	// Registers
	Registers *bool `json:"registers,omitempty"`


	// DualRegisters
	DualRegisters *bool `json:"dualRegisters,omitempty"`


	// HardwareIdType
	HardwareIdType *string `json:"hardwareIdType,omitempty"`


	// AllowReboot
	AllowReboot *bool `json:"allowReboot,omitempty"`


	// NoRebalance
	NoRebalance *bool `json:"noRebalance,omitempty"`


	// NoCloudProvisioning
	NoCloudProvisioning *bool `json:"noCloudProvisioning,omitempty"`


	// MediaCodecs
	MediaCodecs *[]string `json:"mediaCodecs,omitempty"`


	// Cdm
	Cdm *bool `json:"cdm,omitempty"`

}

func (u *Phonecapabilities) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonecapabilities

	

	return json.Marshal(&struct { 
		Provisions *bool `json:"provisions,omitempty"`
		
		Registers *bool `json:"registers,omitempty"`
		
		DualRegisters *bool `json:"dualRegisters,omitempty"`
		
		HardwareIdType *string `json:"hardwareIdType,omitempty"`
		
		AllowReboot *bool `json:"allowReboot,omitempty"`
		
		NoRebalance *bool `json:"noRebalance,omitempty"`
		
		NoCloudProvisioning *bool `json:"noCloudProvisioning,omitempty"`
		
		MediaCodecs *[]string `json:"mediaCodecs,omitempty"`
		
		Cdm *bool `json:"cdm,omitempty"`
		*Alias
	}{ 
		Provisions: u.Provisions,
		
		Registers: u.Registers,
		
		DualRegisters: u.DualRegisters,
		
		HardwareIdType: u.HardwareIdType,
		
		AllowReboot: u.AllowReboot,
		
		NoRebalance: u.NoRebalance,
		
		NoCloudProvisioning: u.NoCloudProvisioning,
		
		MediaCodecs: u.MediaCodecs,
		
		Cdm: u.Cdm,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Phonecapabilities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
