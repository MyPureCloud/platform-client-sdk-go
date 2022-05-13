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

func (o *Phonecapabilities) MarshalJSON() ([]byte, error) {
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
		Provisions: o.Provisions,
		
		Registers: o.Registers,
		
		DualRegisters: o.DualRegisters,
		
		HardwareIdType: o.HardwareIdType,
		
		AllowReboot: o.AllowReboot,
		
		NoRebalance: o.NoRebalance,
		
		NoCloudProvisioning: o.NoCloudProvisioning,
		
		MediaCodecs: o.MediaCodecs,
		
		Cdm: o.Cdm,
		Alias:    (*Alias)(o),
	})
}

func (o *Phonecapabilities) UnmarshalJSON(b []byte) error {
	var PhonecapabilitiesMap map[string]interface{}
	err := json.Unmarshal(b, &PhonecapabilitiesMap)
	if err != nil {
		return err
	}
	
	if Provisions, ok := PhonecapabilitiesMap["provisions"].(bool); ok {
		o.Provisions = &Provisions
	}
    
	if Registers, ok := PhonecapabilitiesMap["registers"].(bool); ok {
		o.Registers = &Registers
	}
    
	if DualRegisters, ok := PhonecapabilitiesMap["dualRegisters"].(bool); ok {
		o.DualRegisters = &DualRegisters
	}
    
	if HardwareIdType, ok := PhonecapabilitiesMap["hardwareIdType"].(string); ok {
		o.HardwareIdType = &HardwareIdType
	}
    
	if AllowReboot, ok := PhonecapabilitiesMap["allowReboot"].(bool); ok {
		o.AllowReboot = &AllowReboot
	}
    
	if NoRebalance, ok := PhonecapabilitiesMap["noRebalance"].(bool); ok {
		o.NoRebalance = &NoRebalance
	}
    
	if NoCloudProvisioning, ok := PhonecapabilitiesMap["noCloudProvisioning"].(bool); ok {
		o.NoCloudProvisioning = &NoCloudProvisioning
	}
    
	if MediaCodecs, ok := PhonecapabilitiesMap["mediaCodecs"].([]interface{}); ok {
		MediaCodecsString, _ := json.Marshal(MediaCodecs)
		json.Unmarshal(MediaCodecsString, &o.MediaCodecs)
	}
	
	if Cdm, ok := PhonecapabilitiesMap["cdm"].(bool); ok {
		o.Cdm = &Cdm
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Phonecapabilities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
