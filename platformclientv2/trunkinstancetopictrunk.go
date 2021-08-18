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

func (u *Trunkinstancetopictrunk) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		ConnectedStatus: u.ConnectedStatus,
		
		OptionsStatus: u.OptionsStatus,
		
		RegistersStatus: u.RegistersStatus,
		
		IpStatus: u.IpStatus,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunk) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
