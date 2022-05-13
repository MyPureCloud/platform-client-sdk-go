package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Trunkinstancetopictrunk) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkinstancetopictrunk
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConnectedStatus *Trunkinstancetopictrunkconnectedstatus `json:"connectedStatus,omitempty"`
		
		OptionsStatus *[]Trunkinstancetopictrunkmetricsoptions `json:"optionsStatus,omitempty"`
		
		RegistersStatus *[]Trunkinstancetopictrunkmetricsregisters `json:"registersStatus,omitempty"`
		
		IpStatus *Trunkinstancetopictrunkmetricsnetworktypeip `json:"ipStatus,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ConnectedStatus: o.ConnectedStatus,
		
		OptionsStatus: o.OptionsStatus,
		
		RegistersStatus: o.RegistersStatus,
		
		IpStatus: o.IpStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkinstancetopictrunk) UnmarshalJSON(b []byte) error {
	var TrunkinstancetopictrunkMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkinstancetopictrunkMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TrunkinstancetopictrunkMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ConnectedStatus, ok := TrunkinstancetopictrunkMap["connectedStatus"].(map[string]interface{}); ok {
		ConnectedStatusString, _ := json.Marshal(ConnectedStatus)
		json.Unmarshal(ConnectedStatusString, &o.ConnectedStatus)
	}
	
	if OptionsStatus, ok := TrunkinstancetopictrunkMap["optionsStatus"].([]interface{}); ok {
		OptionsStatusString, _ := json.Marshal(OptionsStatus)
		json.Unmarshal(OptionsStatusString, &o.OptionsStatus)
	}
	
	if RegistersStatus, ok := TrunkinstancetopictrunkMap["registersStatus"].([]interface{}); ok {
		RegistersStatusString, _ := json.Marshal(RegistersStatus)
		json.Unmarshal(RegistersStatusString, &o.RegistersStatus)
	}
	
	if IpStatus, ok := TrunkinstancetopictrunkMap["ipStatus"].(map[string]interface{}); ok {
		IpStatusString, _ := json.Marshal(IpStatus)
		json.Unmarshal(IpStatusString, &o.IpStatus)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunk) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
