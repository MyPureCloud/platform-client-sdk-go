package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Posttextrequest
type Posttextrequest struct { 
	// BotId - ID of the bot to send the text to.
	BotId *string `json:"botId,omitempty"`


	// BotAlias - Alias/Version of the bot
	BotAlias *string `json:"botAlias,omitempty"`


	// IntegrationId - the integration service id for the bot's credentials
	IntegrationId *string `json:"integrationId,omitempty"`


	// BotSessionId - GUID for this bot's session
	BotSessionId *string `json:"botSessionId,omitempty"`


	// PostTextMessage - Message to send to the bot
	PostTextMessage *Posttextmessage `json:"postTextMessage,omitempty"`


	// LanguageCode - The launguage code the bot will run under
	LanguageCode *string `json:"languageCode,omitempty"`


	// BotSessionTimeoutMinutes - Override timeout for the bot session. This should be greater than 10 minutes.
	BotSessionTimeoutMinutes *int `json:"botSessionTimeoutMinutes,omitempty"`


	// BotChannels - The channels this bot is utilizing
	BotChannels *[]string `json:"botChannels,omitempty"`


	// BotCorrelationId - Id for tracking the activity - this will be returned in the response
	BotCorrelationId *string `json:"botCorrelationId,omitempty"`


	// MessagingPlatformType - If the channels list contains a 'Messaging' item and the messaging platform is known, include it here to get accurate analytics
	MessagingPlatformType *string `json:"messagingPlatformType,omitempty"`


	// AmazonLexRequest - Provider specific settings, if any
	AmazonLexRequest *Amazonlexrequest `json:"amazonLexRequest,omitempty"`


	// GoogleDialogflow - Provider specific settings, if any
	GoogleDialogflow *Googledialogflowcustomsettings `json:"googleDialogflow,omitempty"`


	// GenesysBotConnector - Provider specific settings, if any
	GenesysBotConnector *Genesysbotconnector `json:"genesysBotConnector,omitempty"`


	// NuanceMixDlg - Provider specific settings, if any
	NuanceMixDlg *Nuancemixdlgsettings `json:"nuanceMixDlg,omitempty"`

}

func (u *Posttextrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Posttextrequest

	

	return json.Marshal(&struct { 
		BotId *string `json:"botId,omitempty"`
		
		BotAlias *string `json:"botAlias,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		BotSessionId *string `json:"botSessionId,omitempty"`
		
		PostTextMessage *Posttextmessage `json:"postTextMessage,omitempty"`
		
		LanguageCode *string `json:"languageCode,omitempty"`
		
		BotSessionTimeoutMinutes *int `json:"botSessionTimeoutMinutes,omitempty"`
		
		BotChannels *[]string `json:"botChannels,omitempty"`
		
		BotCorrelationId *string `json:"botCorrelationId,omitempty"`
		
		MessagingPlatformType *string `json:"messagingPlatformType,omitempty"`
		
		AmazonLexRequest *Amazonlexrequest `json:"amazonLexRequest,omitempty"`
		
		GoogleDialogflow *Googledialogflowcustomsettings `json:"googleDialogflow,omitempty"`
		
		GenesysBotConnector *Genesysbotconnector `json:"genesysBotConnector,omitempty"`
		
		NuanceMixDlg *Nuancemixdlgsettings `json:"nuanceMixDlg,omitempty"`
		*Alias
	}{ 
		BotId: u.BotId,
		
		BotAlias: u.BotAlias,
		
		IntegrationId: u.IntegrationId,
		
		BotSessionId: u.BotSessionId,
		
		PostTextMessage: u.PostTextMessage,
		
		LanguageCode: u.LanguageCode,
		
		BotSessionTimeoutMinutes: u.BotSessionTimeoutMinutes,
		
		BotChannels: u.BotChannels,
		
		BotCorrelationId: u.BotCorrelationId,
		
		MessagingPlatformType: u.MessagingPlatformType,
		
		AmazonLexRequest: u.AmazonLexRequest,
		
		GoogleDialogflow: u.GoogleDialogflow,
		
		GenesysBotConnector: u.GenesysBotConnector,
		
		NuanceMixDlg: u.NuanceMixDlg,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Posttextrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
