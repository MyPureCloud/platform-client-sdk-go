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

func (o *Posttextrequest) MarshalJSON() ([]byte, error) {
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
		BotId: o.BotId,
		
		BotAlias: o.BotAlias,
		
		IntegrationId: o.IntegrationId,
		
		BotSessionId: o.BotSessionId,
		
		PostTextMessage: o.PostTextMessage,
		
		LanguageCode: o.LanguageCode,
		
		BotSessionTimeoutMinutes: o.BotSessionTimeoutMinutes,
		
		BotChannels: o.BotChannels,
		
		BotCorrelationId: o.BotCorrelationId,
		
		MessagingPlatformType: o.MessagingPlatformType,
		
		AmazonLexRequest: o.AmazonLexRequest,
		
		GoogleDialogflow: o.GoogleDialogflow,
		
		GenesysBotConnector: o.GenesysBotConnector,
		
		NuanceMixDlg: o.NuanceMixDlg,
		Alias:    (*Alias)(o),
	})
}

func (o *Posttextrequest) UnmarshalJSON(b []byte) error {
	var PosttextrequestMap map[string]interface{}
	err := json.Unmarshal(b, &PosttextrequestMap)
	if err != nil {
		return err
	}
	
	if BotId, ok := PosttextrequestMap["botId"].(string); ok {
		o.BotId = &BotId
	}
    
	if BotAlias, ok := PosttextrequestMap["botAlias"].(string); ok {
		o.BotAlias = &BotAlias
	}
    
	if IntegrationId, ok := PosttextrequestMap["integrationId"].(string); ok {
		o.IntegrationId = &IntegrationId
	}
    
	if BotSessionId, ok := PosttextrequestMap["botSessionId"].(string); ok {
		o.BotSessionId = &BotSessionId
	}
    
	if PostTextMessage, ok := PosttextrequestMap["postTextMessage"].(map[string]interface{}); ok {
		PostTextMessageString, _ := json.Marshal(PostTextMessage)
		json.Unmarshal(PostTextMessageString, &o.PostTextMessage)
	}
	
	if LanguageCode, ok := PosttextrequestMap["languageCode"].(string); ok {
		o.LanguageCode = &LanguageCode
	}
    
	if BotSessionTimeoutMinutes, ok := PosttextrequestMap["botSessionTimeoutMinutes"].(float64); ok {
		BotSessionTimeoutMinutesInt := int(BotSessionTimeoutMinutes)
		o.BotSessionTimeoutMinutes = &BotSessionTimeoutMinutesInt
	}
	
	if BotChannels, ok := PosttextrequestMap["botChannels"].([]interface{}); ok {
		BotChannelsString, _ := json.Marshal(BotChannels)
		json.Unmarshal(BotChannelsString, &o.BotChannels)
	}
	
	if BotCorrelationId, ok := PosttextrequestMap["botCorrelationId"].(string); ok {
		o.BotCorrelationId = &BotCorrelationId
	}
    
	if MessagingPlatformType, ok := PosttextrequestMap["messagingPlatformType"].(string); ok {
		o.MessagingPlatformType = &MessagingPlatformType
	}
    
	if AmazonLexRequest, ok := PosttextrequestMap["amazonLexRequest"].(map[string]interface{}); ok {
		AmazonLexRequestString, _ := json.Marshal(AmazonLexRequest)
		json.Unmarshal(AmazonLexRequestString, &o.AmazonLexRequest)
	}
	
	if GoogleDialogflow, ok := PosttextrequestMap["googleDialogflow"].(map[string]interface{}); ok {
		GoogleDialogflowString, _ := json.Marshal(GoogleDialogflow)
		json.Unmarshal(GoogleDialogflowString, &o.GoogleDialogflow)
	}
	
	if GenesysBotConnector, ok := PosttextrequestMap["genesysBotConnector"].(map[string]interface{}); ok {
		GenesysBotConnectorString, _ := json.Marshal(GenesysBotConnector)
		json.Unmarshal(GenesysBotConnectorString, &o.GenesysBotConnector)
	}
	
	if NuanceMixDlg, ok := PosttextrequestMap["nuanceMixDlg"].(map[string]interface{}); ok {
		NuanceMixDlgString, _ := json.Marshal(NuanceMixDlg)
		json.Unmarshal(NuanceMixDlgString, &o.NuanceMixDlg)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Posttextrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
