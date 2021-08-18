package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lineintegrationrequest
type Lineintegrationrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the LINE Integration
	Name *string `json:"name,omitempty"`


	// ChannelId - The Channel Id from LINE messenger. New Official LINE account: To create a new official account, LINE requires a Webhook URL. It can be created without specifying Channel Id & Channel Secret. Once the Official account is created by LINE, use the update LINE Integration API to update Channel Id and Channel Secret.  All other accounts: Channel Id is mandatory. (NOTE: ChannelId can only be updated if the integration is set to inactive)
	ChannelId *string `json:"channelId,omitempty"`


	// ChannelSecret - The Channel Secret from LINE messenger. New Official LINE account: To create a new official account, LINE requires a Webhook URL. It can be created without specifying Channel Id & Channel Secret. Once the Official account is created by LINE, use the update LINE Integration API to update Channel Id and Channel Secret.  All other accounts: Channel Secret is mandatory. (NOTE: ChannelSecret can only be updated if the integration is set to inactive)
	ChannelSecret *string `json:"channelSecret,omitempty"`


	// SwitcherSecret - The Switcher Secret from LINE messenger. Some line official accounts are switcher functionality enabled. If the LINE account used for this integration is switcher enabled, then switcher secret is a required field. This secret can be found in your create documentation provided by LINE
	SwitcherSecret *string `json:"switcherSecret,omitempty"`


	// ServiceCode - The Service Code from LINE messenger. Only applicable to LINE Enterprise accounts. This service code can be found in your create documentation provided by LINE
	ServiceCode *string `json:"serviceCode,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Lineintegrationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lineintegrationrequest

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ChannelId *string `json:"channelId,omitempty"`
		
		ChannelSecret *string `json:"channelSecret,omitempty"`
		
		SwitcherSecret *string `json:"switcherSecret,omitempty"`
		
		ServiceCode *string `json:"serviceCode,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ChannelId: u.ChannelId,
		
		ChannelSecret: u.ChannelSecret,
		
		SwitcherSecret: u.SwitcherSecret,
		
		ServiceCode: u.ServiceCode,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Lineintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
