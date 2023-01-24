package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Ucintegration - UC Integration UI configuration data
type Ucintegration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// UcIntegrationKey - ucIntegrationKey
	UcIntegrationKey *string `json:"ucIntegrationKey,omitempty"`

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

	// UserPermissions - userPermissions
	UserPermissions *[]string `json:"userPermissions,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Ucintegration) SetField(field string, fieldValue interface{}) {
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

func (o Ucintegration) MarshalJSON() ([]byte, error) {
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
	type Alias Ucintegration
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UcIntegrationKey *string `json:"ucIntegrationKey,omitempty"`
		
		IntegrationPresenceSource *string `json:"integrationPresenceSource,omitempty"`
		
		PbxPermission *string `json:"pbxPermission,omitempty"`
		
		Icon *Ucicon `json:"icon,omitempty"`
		
		BadgeIcons *map[string]Ucicon `json:"badgeIcons,omitempty"`
		
		I10n *map[string]Uci10n `json:"i10n,omitempty"`
		
		PolledPresence *bool `json:"polledPresence,omitempty"`
		
		UserPermissions *[]string `json:"userPermissions,omitempty"`
		
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
		
		UserPermissions: o.UserPermissions,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Ucintegration) UnmarshalJSON(b []byte) error {
	var UcintegrationMap map[string]interface{}
	err := json.Unmarshal(b, &UcintegrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UcintegrationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UcintegrationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if UcIntegrationKey, ok := UcintegrationMap["ucIntegrationKey"].(string); ok {
		o.UcIntegrationKey = &UcIntegrationKey
	}
    
	if IntegrationPresenceSource, ok := UcintegrationMap["integrationPresenceSource"].(string); ok {
		o.IntegrationPresenceSource = &IntegrationPresenceSource
	}
    
	if PbxPermission, ok := UcintegrationMap["pbxPermission"].(string); ok {
		o.PbxPermission = &PbxPermission
	}
    
	if Icon, ok := UcintegrationMap["icon"].(map[string]interface{}); ok {
		IconString, _ := json.Marshal(Icon)
		json.Unmarshal(IconString, &o.Icon)
	}
	
	if BadgeIcons, ok := UcintegrationMap["badgeIcons"].(map[string]interface{}); ok {
		BadgeIconsString, _ := json.Marshal(BadgeIcons)
		json.Unmarshal(BadgeIconsString, &o.BadgeIcons)
	}
	
	if I10n, ok := UcintegrationMap["i10n"].(map[string]interface{}); ok {
		I10nString, _ := json.Marshal(I10n)
		json.Unmarshal(I10nString, &o.I10n)
	}
	
	if PolledPresence, ok := UcintegrationMap["polledPresence"].(bool); ok {
		o.PolledPresence = &PolledPresence
	}
    
	if UserPermissions, ok := UcintegrationMap["userPermissions"].([]interface{}); ok {
		UserPermissionsString, _ := json.Marshal(UserPermissions)
		json.Unmarshal(UserPermissionsString, &o.UserPermissions)
	}
	
	if SelfUri, ok := UcintegrationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Ucintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
