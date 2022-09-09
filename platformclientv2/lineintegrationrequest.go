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


	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`


	// MessagingSetting - Defines the message settings to be applied for this integration
	MessagingSetting *Messagingsettingrequestreference `json:"messagingSetting,omitempty"`


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

func (o *Lineintegrationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lineintegrationrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		MessagingSetting *Messagingsettingrequestreference `json:"messagingSetting,omitempty"`
		
		ChannelId *string `json:"channelId,omitempty"`
		
		ChannelSecret *string `json:"channelSecret,omitempty"`
		
		SwitcherSecret *string `json:"switcherSecret,omitempty"`
		
		ServiceCode *string `json:"serviceCode,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		ChannelId: o.ChannelId,
		
		ChannelSecret: o.ChannelSecret,
		
		SwitcherSecret: o.SwitcherSecret,
		
		ServiceCode: o.ServiceCode,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Lineintegrationrequest) UnmarshalJSON(b []byte) error {
	var LineintegrationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &LineintegrationrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LineintegrationrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := LineintegrationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := LineintegrationrequestMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := LineintegrationrequestMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if ChannelId, ok := LineintegrationrequestMap["channelId"].(string); ok {
		o.ChannelId = &ChannelId
	}
    
	if ChannelSecret, ok := LineintegrationrequestMap["channelSecret"].(string); ok {
		o.ChannelSecret = &ChannelSecret
	}
    
	if SwitcherSecret, ok := LineintegrationrequestMap["switcherSecret"].(string); ok {
		o.SwitcherSecret = &SwitcherSecret
	}
    
	if ServiceCode, ok := LineintegrationrequestMap["serviceCode"].(string); ok {
		o.ServiceCode = &ServiceCode
	}
    
	if SelfUri, ok := LineintegrationrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Lineintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
