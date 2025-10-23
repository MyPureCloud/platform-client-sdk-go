package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Appleintegrationupdaterequest
type Appleintegrationupdaterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Apple messaging integration.
	Name *string `json:"name,omitempty"`

	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`

	// MessagingSetting - Defines the message settings to be applied for this integration
	MessagingSetting *Messagingsettingrequestreference `json:"messagingSetting,omitempty"`

	// BusinessName - The name of the business.
	BusinessName *string `json:"businessName,omitempty"`

	// LogoUrl - The url of the businesses logo.
	LogoUrl *string `json:"logoUrl,omitempty"`

	// AppleIMessageApp - Interactive Application (iMessage App) Settings.
	AppleIMessageApp *Appleimessageapp `json:"appleIMessageApp,omitempty"`

	// AppleAuthentication - The Apple Messages for Business authentication setting.
	AppleAuthentication *Appleauthentication `json:"appleAuthentication,omitempty"`

	// ApplePay - Apple Pay settings.
	ApplePay *Applepay `json:"applePay,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Appleintegrationupdaterequest) SetField(field string, fieldValue interface{}) {
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

func (o Appleintegrationupdaterequest) MarshalJSON() ([]byte, error) {
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
	type Alias Appleintegrationupdaterequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		MessagingSetting *Messagingsettingrequestreference `json:"messagingSetting,omitempty"`
		
		BusinessName *string `json:"businessName,omitempty"`
		
		LogoUrl *string `json:"logoUrl,omitempty"`
		
		AppleIMessageApp *Appleimessageapp `json:"appleIMessageApp,omitempty"`
		
		AppleAuthentication *Appleauthentication `json:"appleAuthentication,omitempty"`
		
		ApplePay *Applepay `json:"applePay,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		BusinessName: o.BusinessName,
		
		LogoUrl: o.LogoUrl,
		
		AppleIMessageApp: o.AppleIMessageApp,
		
		AppleAuthentication: o.AppleAuthentication,
		
		ApplePay: o.ApplePay,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Appleintegrationupdaterequest) UnmarshalJSON(b []byte) error {
	var AppleintegrationupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &AppleintegrationupdaterequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AppleintegrationupdaterequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AppleintegrationupdaterequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := AppleintegrationupdaterequestMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := AppleintegrationupdaterequestMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if BusinessName, ok := AppleintegrationupdaterequestMap["businessName"].(string); ok {
		o.BusinessName = &BusinessName
	}
    
	if LogoUrl, ok := AppleintegrationupdaterequestMap["logoUrl"].(string); ok {
		o.LogoUrl = &LogoUrl
	}
    
	if AppleIMessageApp, ok := AppleintegrationupdaterequestMap["appleIMessageApp"].(map[string]interface{}); ok {
		AppleIMessageAppString, _ := json.Marshal(AppleIMessageApp)
		json.Unmarshal(AppleIMessageAppString, &o.AppleIMessageApp)
	}
	
	if AppleAuthentication, ok := AppleintegrationupdaterequestMap["appleAuthentication"].(map[string]interface{}); ok {
		AppleAuthenticationString, _ := json.Marshal(AppleAuthentication)
		json.Unmarshal(AppleAuthenticationString, &o.AppleAuthentication)
	}
	
	if ApplePay, ok := AppleintegrationupdaterequestMap["applePay"].(map[string]interface{}); ok {
		ApplePayString, _ := json.Marshal(ApplePay)
		json.Unmarshal(ApplePayString, &o.ApplePay)
	}
	
	if SelfUri, ok := AppleintegrationupdaterequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Appleintegrationupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
