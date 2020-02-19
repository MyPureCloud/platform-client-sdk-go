package platformclientv2
import (
	"encoding/json"
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

}

// String returns a JSON representation of the model
func (o *Phonecapabilities) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
