package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Openintegrationrequest
type Openintegrationrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Open messaging integration.
	Name *string `json:"name,omitempty"`

	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`

	// MessagingSetting - Defines the message settings to be applied for this integration
	MessagingSetting *Messagingsettingrequestreference `json:"messagingSetting,omitempty"`

	// OutboundNotificationWebhookUrl - The outbound notification webhook URL for the Open messaging integration.
	OutboundNotificationWebhookUrl *string `json:"outboundNotificationWebhookUrl,omitempty"`

	// OutboundNotificationWebhookSignatureSecretToken - The outbound notification webhook signature secret token. This token must be longer than 15 characters.
	OutboundNotificationWebhookSignatureSecretToken *string `json:"outboundNotificationWebhookSignatureSecretToken,omitempty"`

	// WebhookHeaders - The user specified headers for the Open messaging integration.
	WebhookHeaders *map[string]string `json:"webhookHeaders,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Openintegrationrequest) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Openintegrationrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Openintegrationrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		MessagingSetting *Messagingsettingrequestreference `json:"messagingSetting,omitempty"`
		
		OutboundNotificationWebhookUrl *string `json:"outboundNotificationWebhookUrl,omitempty"`
		
		OutboundNotificationWebhookSignatureSecretToken *string `json:"outboundNotificationWebhookSignatureSecretToken,omitempty"`
		
		WebhookHeaders *map[string]string `json:"webhookHeaders,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		OutboundNotificationWebhookUrl: o.OutboundNotificationWebhookUrl,
		
		OutboundNotificationWebhookSignatureSecretToken: o.OutboundNotificationWebhookSignatureSecretToken,
		
		WebhookHeaders: o.WebhookHeaders,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Openintegrationrequest) UnmarshalJSON(b []byte) error {
	var OpenintegrationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &OpenintegrationrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OpenintegrationrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OpenintegrationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := OpenintegrationrequestMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := OpenintegrationrequestMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if OutboundNotificationWebhookUrl, ok := OpenintegrationrequestMap["outboundNotificationWebhookUrl"].(string); ok {
		o.OutboundNotificationWebhookUrl = &OutboundNotificationWebhookUrl
	}
    
	if OutboundNotificationWebhookSignatureSecretToken, ok := OpenintegrationrequestMap["outboundNotificationWebhookSignatureSecretToken"].(string); ok {
		o.OutboundNotificationWebhookSignatureSecretToken = &OutboundNotificationWebhookSignatureSecretToken
	}
    
	if WebhookHeaders, ok := OpenintegrationrequestMap["webhookHeaders"].(map[string]interface{}); ok {
		WebhookHeadersString, _ := json.Marshal(WebhookHeaders)
		json.Unmarshal(WebhookHeadersString, &o.WebhookHeaders)
	}
	
	if SelfUri, ok := OpenintegrationrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Openintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
