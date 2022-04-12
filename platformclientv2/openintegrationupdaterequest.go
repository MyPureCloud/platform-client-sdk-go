package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Openintegrationupdaterequest
type Openintegrationupdaterequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Open messaging integration.
	Name *string `json:"name,omitempty"`


	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`


	// MessagingSetting
	MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`


	// OutboundNotificationWebhookUrl - The outbound notification webhook URL for the Open messaging integration.
	OutboundNotificationWebhookUrl *string `json:"outboundNotificationWebhookUrl,omitempty"`


	// OutboundNotificationWebhookSignatureSecretToken - The outbound notification webhook signature secret token.
	OutboundNotificationWebhookSignatureSecretToken *string `json:"outboundNotificationWebhookSignatureSecretToken,omitempty"`


	// WebhookHeaders - The user specified headers for the Open messaging integration.
	WebhookHeaders *map[string]string `json:"webhookHeaders,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Openintegrationupdaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Openintegrationupdaterequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`
		
		OutboundNotificationWebhookUrl *string `json:"outboundNotificationWebhookUrl,omitempty"`
		
		OutboundNotificationWebhookSignatureSecretToken *string `json:"outboundNotificationWebhookSignatureSecretToken,omitempty"`
		
		WebhookHeaders *map[string]string `json:"webhookHeaders,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		OutboundNotificationWebhookUrl: o.OutboundNotificationWebhookUrl,
		
		OutboundNotificationWebhookSignatureSecretToken: o.OutboundNotificationWebhookSignatureSecretToken,
		
		WebhookHeaders: o.WebhookHeaders,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Openintegrationupdaterequest) UnmarshalJSON(b []byte) error {
	var OpenintegrationupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &OpenintegrationupdaterequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OpenintegrationupdaterequestMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := OpenintegrationupdaterequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SupportedContent, ok := OpenintegrationupdaterequestMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := OpenintegrationupdaterequestMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if OutboundNotificationWebhookUrl, ok := OpenintegrationupdaterequestMap["outboundNotificationWebhookUrl"].(string); ok {
		o.OutboundNotificationWebhookUrl = &OutboundNotificationWebhookUrl
	}
	
	if OutboundNotificationWebhookSignatureSecretToken, ok := OpenintegrationupdaterequestMap["outboundNotificationWebhookSignatureSecretToken"].(string); ok {
		o.OutboundNotificationWebhookSignatureSecretToken = &OutboundNotificationWebhookSignatureSecretToken
	}
	
	if WebhookHeaders, ok := OpenintegrationupdaterequestMap["webhookHeaders"].(map[string]interface{}); ok {
		WebhookHeadersString, _ := json.Marshal(WebhookHeaders)
		json.Unmarshal(WebhookHeadersString, &o.WebhookHeaders)
	}
	
	if SelfUri, ok := OpenintegrationupdaterequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Openintegrationupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
