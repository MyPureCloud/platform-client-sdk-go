package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Organization
type Organization struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// DefaultLanguage - The default language for this organization. Example: 'en'
	DefaultLanguage *string `json:"defaultLanguage,omitempty"`

	// DefaultCountryCode - The default country code for this organization. Example: 'US'
	DefaultCountryCode *string `json:"defaultCountryCode,omitempty"`

	// ThirdPartyOrgName - The short name for the organization. This field is globally unique and cannot be changed.
	ThirdPartyOrgName *string `json:"thirdPartyOrgName,omitempty"`

	// ThirdPartyURI
	ThirdPartyURI *string `json:"thirdPartyURI,omitempty"`

	// Domain
	Domain *string `json:"domain,omitempty"`

	// Version - The current version of the organization.
	Version *int `json:"version,omitempty"`

	// State - The current state. Examples are active, inactive, deleted.
	State *string `json:"state,omitempty"`

	// DefaultSiteId
	DefaultSiteId *string `json:"defaultSiteId,omitempty"`

	// SupportURI - Email address where support tickets are sent to.
	SupportURI *string `json:"supportURI,omitempty"`

	// VoicemailEnabled
	VoicemailEnabled *bool `json:"voicemailEnabled,omitempty"`

	// ProductPlatform - Organizations Originating Platform.
	ProductPlatform *string `json:"productPlatform,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// Features - The state of features available for the organization.
	Features *map[string]bool `json:"features,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Organization) SetField(field string, fieldValue interface{}) {
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

func (o Organization) MarshalJSON() ([]byte, error) {
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
	type Alias Organization
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DefaultLanguage *string `json:"defaultLanguage,omitempty"`
		
		DefaultCountryCode *string `json:"defaultCountryCode,omitempty"`
		
		ThirdPartyOrgName *string `json:"thirdPartyOrgName,omitempty"`
		
		ThirdPartyURI *string `json:"thirdPartyURI,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DefaultSiteId *string `json:"defaultSiteId,omitempty"`
		
		SupportURI *string `json:"supportURI,omitempty"`
		
		VoicemailEnabled *bool `json:"voicemailEnabled,omitempty"`
		
		ProductPlatform *string `json:"productPlatform,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Features *map[string]bool `json:"features,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DefaultLanguage: o.DefaultLanguage,
		
		DefaultCountryCode: o.DefaultCountryCode,
		
		ThirdPartyOrgName: o.ThirdPartyOrgName,
		
		ThirdPartyURI: o.ThirdPartyURI,
		
		Domain: o.Domain,
		
		Version: o.Version,
		
		State: o.State,
		
		DefaultSiteId: o.DefaultSiteId,
		
		SupportURI: o.SupportURI,
		
		VoicemailEnabled: o.VoicemailEnabled,
		
		ProductPlatform: o.ProductPlatform,
		
		SelfUri: o.SelfUri,
		
		Features: o.Features,
		Alias:    (Alias)(o),
	})
}

func (o *Organization) UnmarshalJSON(b []byte) error {
	var OrganizationMap map[string]interface{}
	err := json.Unmarshal(b, &OrganizationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OrganizationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OrganizationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if DefaultLanguage, ok := OrganizationMap["defaultLanguage"].(string); ok {
		o.DefaultLanguage = &DefaultLanguage
	}
    
	if DefaultCountryCode, ok := OrganizationMap["defaultCountryCode"].(string); ok {
		o.DefaultCountryCode = &DefaultCountryCode
	}
    
	if ThirdPartyOrgName, ok := OrganizationMap["thirdPartyOrgName"].(string); ok {
		o.ThirdPartyOrgName = &ThirdPartyOrgName
	}
    
	if ThirdPartyURI, ok := OrganizationMap["thirdPartyURI"].(string); ok {
		o.ThirdPartyURI = &ThirdPartyURI
	}
    
	if Domain, ok := OrganizationMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if Version, ok := OrganizationMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if State, ok := OrganizationMap["state"].(string); ok {
		o.State = &State
	}
    
	if DefaultSiteId, ok := OrganizationMap["defaultSiteId"].(string); ok {
		o.DefaultSiteId = &DefaultSiteId
	}
    
	if SupportURI, ok := OrganizationMap["supportURI"].(string); ok {
		o.SupportURI = &SupportURI
	}
    
	if VoicemailEnabled, ok := OrganizationMap["voicemailEnabled"].(bool); ok {
		o.VoicemailEnabled = &VoicemailEnabled
	}
    
	if ProductPlatform, ok := OrganizationMap["productPlatform"].(string); ok {
		o.ProductPlatform = &ProductPlatform
	}
    
	if SelfUri, ok := OrganizationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Features, ok := OrganizationMap["features"].(map[string]interface{}); ok {
		FeaturesString, _ := json.Marshal(Features)
		json.Unmarshal(FeaturesString, &o.Features)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Organization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
