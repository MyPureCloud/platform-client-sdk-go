package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Posttextrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
