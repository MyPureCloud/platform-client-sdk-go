package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Publickeycredentialcreationoptions
type Publickeycredentialcreationoptions struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Challenge - Cryptographic challenge from the relying party (base64url-encoded). Must be returned to the relying party in the authenticator's response.
	Challenge *string `json:"challenge,omitempty"`

	// Rp - Information about the relying party.
	Rp *Relyingpartyentity `json:"rp,omitempty"`

	// User - Information about the user being registered.
	User *Userentity `json:"user,omitempty"`

	// PubKeyCredParams - Public key credential parameters acceptable to the relying party, in order of preference.
	PubKeyCredParams *[]Credentialparameter `json:"pubKeyCredParams,omitempty"`

	// Timeout - Time in milliseconds the relying party is willing to wait for the registration operation to complete.
	Timeout *int `json:"timeout,omitempty"`

	// ExcludeCredentials - Credentials that should be excluded from registration (e.g., to prevent re-registering an existing authenticator).
	ExcludeCredentials *[]Credentialdescriptor `json:"excludeCredentials,omitempty"`

	// AuthenticatorSelection - Constraints on the type of authenticator that can be used.
	AuthenticatorSelection *Authenticatorselection `json:"authenticatorSelection,omitempty"`

	// Hints - Hints about the type of authenticator the user should use (e.g., 'security-key', 'client-device', 'hybrid').
	Hints *[]string `json:"hints,omitempty"`

	// Attestation - The relying party's attestation conveyance preference ('none', 'indirect', 'direct', or 'enterprise').
	Attestation *string `json:"attestation,omitempty"`

	// AttestationFormats - Acceptable attestation statement formats, in order of preference.
	AttestationFormats *[]string `json:"attestationFormats,omitempty"`

	// Extensions - Inputs to client-side WebAuthn extensions.
	Extensions *map[string]interface{} `json:"extensions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Publickeycredentialcreationoptions) SetField(field string, fieldValue interface{}) {
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

func (o Publickeycredentialcreationoptions) MarshalJSON() ([]byte, error) {
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
	type Alias Publickeycredentialcreationoptions
	
	return json.Marshal(&struct { 
		Challenge *string `json:"challenge,omitempty"`
		
		Rp *Relyingpartyentity `json:"rp,omitempty"`
		
		User *Userentity `json:"user,omitempty"`
		
		PubKeyCredParams *[]Credentialparameter `json:"pubKeyCredParams,omitempty"`
		
		Timeout *int `json:"timeout,omitempty"`
		
		ExcludeCredentials *[]Credentialdescriptor `json:"excludeCredentials,omitempty"`
		
		AuthenticatorSelection *Authenticatorselection `json:"authenticatorSelection,omitempty"`
		
		Hints *[]string `json:"hints,omitempty"`
		
		Attestation *string `json:"attestation,omitempty"`
		
		AttestationFormats *[]string `json:"attestationFormats,omitempty"`
		
		Extensions *map[string]interface{} `json:"extensions,omitempty"`
		Alias
	}{ 
		Challenge: o.Challenge,
		
		Rp: o.Rp,
		
		User: o.User,
		
		PubKeyCredParams: o.PubKeyCredParams,
		
		Timeout: o.Timeout,
		
		ExcludeCredentials: o.ExcludeCredentials,
		
		AuthenticatorSelection: o.AuthenticatorSelection,
		
		Hints: o.Hints,
		
		Attestation: o.Attestation,
		
		AttestationFormats: o.AttestationFormats,
		
		Extensions: o.Extensions,
		Alias:    (Alias)(o),
	})
}

func (o *Publickeycredentialcreationoptions) UnmarshalJSON(b []byte) error {
	var PublickeycredentialcreationoptionsMap map[string]interface{}
	err := json.Unmarshal(b, &PublickeycredentialcreationoptionsMap)
	if err != nil {
		return err
	}
	
	if Challenge, ok := PublickeycredentialcreationoptionsMap["challenge"].(string); ok {
		o.Challenge = &Challenge
	}
    
	if Rp, ok := PublickeycredentialcreationoptionsMap["rp"].(map[string]interface{}); ok {
		RpString, _ := json.Marshal(Rp)
		json.Unmarshal(RpString, &o.Rp)
	}
	
	if User, ok := PublickeycredentialcreationoptionsMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if PubKeyCredParams, ok := PublickeycredentialcreationoptionsMap["pubKeyCredParams"].([]interface{}); ok {
		PubKeyCredParamsString, _ := json.Marshal(PubKeyCredParams)
		json.Unmarshal(PubKeyCredParamsString, &o.PubKeyCredParams)
	}
	
	if Timeout, ok := PublickeycredentialcreationoptionsMap["timeout"].(float64); ok {
		TimeoutInt := int(Timeout)
		o.Timeout = &TimeoutInt
	}
	
	if ExcludeCredentials, ok := PublickeycredentialcreationoptionsMap["excludeCredentials"].([]interface{}); ok {
		ExcludeCredentialsString, _ := json.Marshal(ExcludeCredentials)
		json.Unmarshal(ExcludeCredentialsString, &o.ExcludeCredentials)
	}
	
	if AuthenticatorSelection, ok := PublickeycredentialcreationoptionsMap["authenticatorSelection"].(map[string]interface{}); ok {
		AuthenticatorSelectionString, _ := json.Marshal(AuthenticatorSelection)
		json.Unmarshal(AuthenticatorSelectionString, &o.AuthenticatorSelection)
	}
	
	if Hints, ok := PublickeycredentialcreationoptionsMap["hints"].([]interface{}); ok {
		HintsString, _ := json.Marshal(Hints)
		json.Unmarshal(HintsString, &o.Hints)
	}
	
	if Attestation, ok := PublickeycredentialcreationoptionsMap["attestation"].(string); ok {
		o.Attestation = &Attestation
	}
    
	if AttestationFormats, ok := PublickeycredentialcreationoptionsMap["attestationFormats"].([]interface{}); ok {
		AttestationFormatsString, _ := json.Marshal(AttestationFormats)
		json.Unmarshal(AttestationFormatsString, &o.AttestationFormats)
	}
	
	if Extensions, ok := PublickeycredentialcreationoptionsMap["extensions"].(map[string]interface{}); ok {
		ExtensionsString, _ := json.Marshal(Extensions)
		json.Unmarshal(ExtensionsString, &o.Extensions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Publickeycredentialcreationoptions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
