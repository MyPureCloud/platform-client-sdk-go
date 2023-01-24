package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentsoauthrequestparameters
type Webdeploymentsoauthrequestparameters struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Code - The authorization code to be sent to the authentication server during the token request.  Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	Code *string `json:"code,omitempty"`

	// RedirectUri - Redirect URI sent in the \"Authentication Request\"Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	RedirectUri *string `json:"redirectUri,omitempty"`

	// Nonce - Required if provided in the \"Authentication Request\". Otherwise should be empty.String value used to associate a Client session with an ID Token, and to mitigate replay attacks. The value is passed through unmodified from the Authentication Request to the ID Token. Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	Nonce *string `json:"nonce,omitempty"`

	// MaxAge - Required if provided in the  \"Authentication Request\". Otherwise should be empty.Specifies the allowable elapsed time in seconds since the last time the End-User was actively authenticated.Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	MaxAge *int `json:"maxAge,omitempty"`

	// CodeVerifier - Required if authorizing using Proof Key for Code Exchange (PKCE). Otherwise should be empty.Random URL-safe string with a minimum length of 43 characters generated at start of authorization flow to mitigate the threat of having the authorization code intercepted. Refer to https://datatracker.ietf.org/doc/html/rfc7636
	CodeVerifier *string `json:"codeVerifier,omitempty"`

	// Iss - Optional parameter. Set it if authorization server discovery metadata authorization_response_iss_parameter_supported is enabled. Refer to https://datatracker.ietf.org/doc/html/rfc9207
	Iss *string `json:"iss,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webdeploymentsoauthrequestparameters) SetField(field string, fieldValue interface{}) {
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

func (o Webdeploymentsoauthrequestparameters) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Webdeploymentsoauthrequestparameters
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		RedirectUri *string `json:"redirectUri,omitempty"`
		
		Nonce *string `json:"nonce,omitempty"`
		
		MaxAge *int `json:"maxAge,omitempty"`
		
		CodeVerifier *string `json:"codeVerifier,omitempty"`
		
		Iss *string `json:"iss,omitempty"`
		Alias
	}{ 
		Code: o.Code,
		
		RedirectUri: o.RedirectUri,
		
		Nonce: o.Nonce,
		
		MaxAge: o.MaxAge,
		
		CodeVerifier: o.CodeVerifier,
		
		Iss: o.Iss,
		Alias:    (Alias)(o),
	})
}

func (o *Webdeploymentsoauthrequestparameters) UnmarshalJSON(b []byte) error {
	var WebdeploymentsoauthrequestparametersMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentsoauthrequestparametersMap)
	if err != nil {
		return err
	}
	
	if Code, ok := WebdeploymentsoauthrequestparametersMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if RedirectUri, ok := WebdeploymentsoauthrequestparametersMap["redirectUri"].(string); ok {
		o.RedirectUri = &RedirectUri
	}
    
	if Nonce, ok := WebdeploymentsoauthrequestparametersMap["nonce"].(string); ok {
		o.Nonce = &Nonce
	}
    
	if MaxAge, ok := WebdeploymentsoauthrequestparametersMap["maxAge"].(float64); ok {
		MaxAgeInt := int(MaxAge)
		o.MaxAge = &MaxAgeInt
	}
	
	if CodeVerifier, ok := WebdeploymentsoauthrequestparametersMap["codeVerifier"].(string); ok {
		o.CodeVerifier = &CodeVerifier
	}
    
	if Iss, ok := WebdeploymentsoauthrequestparametersMap["iss"].(string); ok {
		o.Iss = &Iss
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentsoauthrequestparameters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
