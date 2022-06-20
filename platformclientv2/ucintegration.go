package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ucintegration - UC Integration UI configuration data
type Ucintegration struct { 
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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Ucintegration) MarshalJSON() ([]byte, error) {
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
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		UcIntegrationKey: o.UcIntegrationKey,
		
		IntegrationPresenceSource: o.IntegrationPresenceSource,
		
		PbxPermission: o.PbxPermission,
		
		Icon: o.Icon,
		
		BadgeIcons: o.BadgeIcons,
		
		I10n: o.I10n,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
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
