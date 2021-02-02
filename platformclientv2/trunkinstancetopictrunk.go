package platformclientv2
import (
	"encoding/json"
)

// Trunkinstancetopictrunk
type Trunkinstancetopictrunk struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ConnectedStatus
	ConnectedStatus *Trunkinstancetopictrunkconnectedstatus `json:"connectedStatus,omitempty"`


	// OptionsStatus
	OptionsStatus *[]Trunkinstancetopictrunkmetricsoptions `json:"optionsStatus,omitempty"`


	// RegistersStatus
	RegistersStatus *[]Trunkinstancetopictrunkmetricsregisters `json:"registersStatus,omitempty"`


	// IpStatus
	IpStatus *Trunkinstancetopictrunkmetricsnetworktypeip `json:"ipStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunk) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
