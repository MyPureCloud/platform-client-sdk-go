package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeofflineconfiguration
type Edgeofflineconfiguration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PairingId - The pairingId for your hardware Edge in the format: 00000-00000-00000-00000-00000.
	PairingId *string `json:"pairingId,omitempty"`

	// Network - Network settings for your hardware Edge.
	Network *Edgeofflineconfigurationnetwork `json:"network,omitempty"`

	// UseVerificationCode - Boolean to know if the verification code will be used to provision the Edge. Only used if the Edge is being provisioned.
	UseVerificationCode *bool `json:"useVerificationCode,omitempty"`

	// CertType - The type of Certificate Authority this Edge will use. Defaults to NotRequested if the Edge is already provisioned. PureCloud signed CA is recommended. Public CA signed by a trusted third party. China CA must be used if the Site's Location is in China.
	CertType *string `json:"certType,omitempty"`

	// Site - The Site that will be associated to the Edge. Required if the Edge is being provisioned.
	Site *Domainentityref `json:"site,omitempty"`

	// Proxy - Edge HTTP proxy configuration for the WAN port. The field can be a hostname, FQDN, IPv4 or IPv6 address. If port is not included, port 80 is assumed.
	Proxy *string `json:"proxy,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Edgeofflineconfiguration) SetField(field string, fieldValue interface{}) {
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

func (o Edgeofflineconfiguration) MarshalJSON() ([]byte, error) {
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
	type Alias Edgeofflineconfiguration
	
	return json.Marshal(&struct { 
		PairingId *string `json:"pairingId,omitempty"`
		
		Network *Edgeofflineconfigurationnetwork `json:"network,omitempty"`
		
		UseVerificationCode *bool `json:"useVerificationCode,omitempty"`
		
		CertType *string `json:"certType,omitempty"`
		
		Site *Domainentityref `json:"site,omitempty"`
		
		Proxy *string `json:"proxy,omitempty"`
		Alias
	}{ 
		PairingId: o.PairingId,
		
		Network: o.Network,
		
		UseVerificationCode: o.UseVerificationCode,
		
		CertType: o.CertType,
		
		Site: o.Site,
		
		Proxy: o.Proxy,
		Alias:    (Alias)(o),
	})
}

func (o *Edgeofflineconfiguration) UnmarshalJSON(b []byte) error {
	var EdgeofflineconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeofflineconfigurationMap)
	if err != nil {
		return err
	}
	
	if PairingId, ok := EdgeofflineconfigurationMap["pairingId"].(string); ok {
		o.PairingId = &PairingId
	}
    
	if Network, ok := EdgeofflineconfigurationMap["network"].(map[string]interface{}); ok {
		NetworkString, _ := json.Marshal(Network)
		json.Unmarshal(NetworkString, &o.Network)
	}
	
	if UseVerificationCode, ok := EdgeofflineconfigurationMap["useVerificationCode"].(bool); ok {
		o.UseVerificationCode = &UseVerificationCode
	}
    
	if CertType, ok := EdgeofflineconfigurationMap["certType"].(string); ok {
		o.CertType = &CertType
	}
    
	if Site, ok := EdgeofflineconfigurationMap["site"].(map[string]interface{}); ok {
		SiteString, _ := json.Marshal(Site)
		json.Unmarshal(SiteString, &o.Site)
	}
	
	if Proxy, ok := EdgeofflineconfigurationMap["proxy"].(string); ok {
		o.Proxy = &Proxy
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeofflineconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
