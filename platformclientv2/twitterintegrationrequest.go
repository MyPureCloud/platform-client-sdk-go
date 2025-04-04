package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Twitterintegrationrequest
type Twitterintegrationrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Twitter Integration
	Name *string `json:"name,omitempty"`

	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`

	// MessagingSetting - Defines the message settings to be applied for this integration
	MessagingSetting *Messagingsettingrequestreference `json:"messagingSetting,omitempty"`

	// SignupCode - The authorization code returned from the signup flow to request access tokens later on
	SignupCode *string `json:"signupCode,omitempty"`

	// AppId - The appId of the Twitter app to register the integration on
	AppId *string `json:"appId,omitempty"`

	// CodeChallenge - The codeChallenge used during the signup flow
	CodeChallenge *string `json:"codeChallenge,omitempty"`

	// RedirectUri - The redirectUri used during the signup flow
	RedirectUri *string `json:"redirectUri,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Twitterintegrationrequest) SetField(field string, fieldValue interface{}) {
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

func (o Twitterintegrationrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Twitterintegrationrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		MessagingSetting *Messagingsettingrequestreference `json:"messagingSetting,omitempty"`
		
		SignupCode *string `json:"signupCode,omitempty"`
		
		AppId *string `json:"appId,omitempty"`
		
		CodeChallenge *string `json:"codeChallenge,omitempty"`
		
		RedirectUri *string `json:"redirectUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		SignupCode: o.SignupCode,
		
		AppId: o.AppId,
		
		CodeChallenge: o.CodeChallenge,
		
		RedirectUri: o.RedirectUri,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Twitterintegrationrequest) UnmarshalJSON(b []byte) error {
	var TwitterintegrationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TwitterintegrationrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TwitterintegrationrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := TwitterintegrationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := TwitterintegrationrequestMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := TwitterintegrationrequestMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if SignupCode, ok := TwitterintegrationrequestMap["signupCode"].(string); ok {
		o.SignupCode = &SignupCode
	}
    
	if AppId, ok := TwitterintegrationrequestMap["appId"].(string); ok {
		o.AppId = &AppId
	}
    
	if CodeChallenge, ok := TwitterintegrationrequestMap["codeChallenge"].(string); ok {
		o.CodeChallenge = &CodeChallenge
	}
    
	if RedirectUri, ok := TwitterintegrationrequestMap["redirectUri"].(string); ok {
		o.RedirectUri = &RedirectUri
	}
    
	if SelfUri, ok := TwitterintegrationrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Twitterintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
