package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Unifiedcommunicationsintegration - UC Integration UI configuration data
type Unifiedcommunicationsintegration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// UcIntegrationKey - ucIntegrationKey
	UcIntegrationKey *Addressableentityref `json:"ucIntegrationKey,omitempty"`

	// IntegrationPresenceSource - integrationPresenceType
	IntegrationPresenceSource *string `json:"integrationPresenceSource,omitempty"`

	// PbxPermission - pbxPermission
	PbxPermission *string `json:"pbxPermission,omitempty"`

	// Icon - icon
	Icon *Ucicon `json:"icon,omitempty"`

	// BadgeIcons - badgeIcon
	BadgeIcons *map[string]Ucicon `json:"badgeIcons,omitempty"`

	// I10n - i10n
	I10n *map[string]Uci10n `json:"i10n,omitempty"`

	// PolledPresence - polledPresence
	PolledPresence *bool `json:"polledPresence,omitempty"`

	// PollIntervalSec - pollIntervalSec
	PollIntervalSec *int `json:"pollIntervalSec,omitempty"`

	// UserPermissions - userPermissions
	UserPermissions *[]string `json:"userPermissions,omitempty"`

	// OauthScopes
	OauthScopes *[]string `json:"oauthScopes,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Unifiedcommunicationsintegration) SetField(field string, fieldValue interface{}) {
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

func (o Unifiedcommunicationsintegration) MarshalJSON() ([]byte, error) {
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
	type Alias Unifiedcommunicationsintegration
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UcIntegrationKey *Addressableentityref `json:"ucIntegrationKey,omitempty"`
		
		IntegrationPresenceSource *string `json:"integrationPresenceSource,omitempty"`
		
		PbxPermission *string `json:"pbxPermission,omitempty"`
		
		Icon *Ucicon `json:"icon,omitempty"`
		
		BadgeIcons *map[string]Ucicon `json:"badgeIcons,omitempty"`
		
		I10n *map[string]Uci10n `json:"i10n,omitempty"`
		
		PolledPresence *bool `json:"polledPresence,omitempty"`
		
		PollIntervalSec *int `json:"pollIntervalSec,omitempty"`
		
		UserPermissions *[]string `json:"userPermissions,omitempty"`
		
		OauthScopes *[]string `json:"oauthScopes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		UcIntegrationKey: o.UcIntegrationKey,
		
		IntegrationPresenceSource: o.IntegrationPresenceSource,
		
		PbxPermission: o.PbxPermission,
		
		Icon: o.Icon,
		
		BadgeIcons: o.BadgeIcons,
		
		I10n: o.I10n,
		
		PolledPresence: o.PolledPresence,
		
		PollIntervalSec: o.PollIntervalSec,
		
		UserPermissions: o.UserPermissions,
		
		OauthScopes: o.OauthScopes,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Unifiedcommunicationsintegration) UnmarshalJSON(b []byte) error {
	var UnifiedcommunicationsintegrationMap map[string]interface{}
	err := json.Unmarshal(b, &UnifiedcommunicationsintegrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UnifiedcommunicationsintegrationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UnifiedcommunicationsintegrationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if UcIntegrationKey, ok := UnifiedcommunicationsintegrationMap["ucIntegrationKey"].(map[string]interface{}); ok {
		UcIntegrationKeyString, _ := json.Marshal(UcIntegrationKey)
		json.Unmarshal(UcIntegrationKeyString, &o.UcIntegrationKey)
	}
	
	if IntegrationPresenceSource, ok := UnifiedcommunicationsintegrationMap["integrationPresenceSource"].(string); ok {
		o.IntegrationPresenceSource = &IntegrationPresenceSource
	}
    
	if PbxPermission, ok := UnifiedcommunicationsintegrationMap["pbxPermission"].(string); ok {
		o.PbxPermission = &PbxPermission
	}
    
	if Icon, ok := UnifiedcommunicationsintegrationMap["icon"].(map[string]interface{}); ok {
		IconString, _ := json.Marshal(Icon)
		json.Unmarshal(IconString, &o.Icon)
	}
	
	if BadgeIcons, ok := UnifiedcommunicationsintegrationMap["badgeIcons"].(map[string]interface{}); ok {
		BadgeIconsString, _ := json.Marshal(BadgeIcons)
		json.Unmarshal(BadgeIconsString, &o.BadgeIcons)
	}
	
	if I10n, ok := UnifiedcommunicationsintegrationMap["i10n"].(map[string]interface{}); ok {
		I10nString, _ := json.Marshal(I10n)
		json.Unmarshal(I10nString, &o.I10n)
	}
	
	if PolledPresence, ok := UnifiedcommunicationsintegrationMap["polledPresence"].(bool); ok {
		o.PolledPresence = &PolledPresence
	}
    
	if PollIntervalSec, ok := UnifiedcommunicationsintegrationMap["pollIntervalSec"].(float64); ok {
		PollIntervalSecInt := int(PollIntervalSec)
		o.PollIntervalSec = &PollIntervalSecInt
	}
	
	if UserPermissions, ok := UnifiedcommunicationsintegrationMap["userPermissions"].([]interface{}); ok {
		UserPermissionsString, _ := json.Marshal(UserPermissions)
		json.Unmarshal(UserPermissionsString, &o.UserPermissions)
	}
	
	if OauthScopes, ok := UnifiedcommunicationsintegrationMap["oauthScopes"].([]interface{}); ok {
		OauthScopesString, _ := json.Marshal(OauthScopes)
		json.Unmarshal(OauthScopesString, &o.OauthScopes)
	}
	
	if SelfUri, ok := UnifiedcommunicationsintegrationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Unifiedcommunicationsintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
