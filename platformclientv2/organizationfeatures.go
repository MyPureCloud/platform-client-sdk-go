package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Organizationfeatures
type Organizationfeatures struct { 
	// RealtimeCIC
	RealtimeCIC *bool `json:"realtimeCIC,omitempty"`


	// Purecloud
	Purecloud *bool `json:"purecloud,omitempty"`


	// Hipaa
	Hipaa *bool `json:"hipaa,omitempty"`


	// UcEnabled
	UcEnabled *bool `json:"ucEnabled,omitempty"`


	// Pci
	Pci *bool `json:"pci,omitempty"`


	// PurecloudVoice
	PurecloudVoice *bool `json:"purecloudVoice,omitempty"`


	// XmppFederation
	XmppFederation *bool `json:"xmppFederation,omitempty"`


	// Chat
	Chat *bool `json:"chat,omitempty"`


	// InformalPhotos
	InformalPhotos *bool `json:"informalPhotos,omitempty"`


	// Directory
	Directory *bool `json:"directory,omitempty"`


	// ContactCenter
	ContactCenter *bool `json:"contactCenter,omitempty"`


	// UnifiedCommunications
	UnifiedCommunications *bool `json:"unifiedCommunications,omitempty"`


	// Custserv
	Custserv *bool `json:"custserv,omitempty"`

}

func (o *Organizationfeatures) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Organizationfeatures
	
	return json.Marshal(&struct { 
		RealtimeCIC *bool `json:"realtimeCIC,omitempty"`
		
		Purecloud *bool `json:"purecloud,omitempty"`
		
		Hipaa *bool `json:"hipaa,omitempty"`
		
		UcEnabled *bool `json:"ucEnabled,omitempty"`
		
		Pci *bool `json:"pci,omitempty"`
		
		PurecloudVoice *bool `json:"purecloudVoice,omitempty"`
		
		XmppFederation *bool `json:"xmppFederation,omitempty"`
		
		Chat *bool `json:"chat,omitempty"`
		
		InformalPhotos *bool `json:"informalPhotos,omitempty"`
		
		Directory *bool `json:"directory,omitempty"`
		
		ContactCenter *bool `json:"contactCenter,omitempty"`
		
		UnifiedCommunications *bool `json:"unifiedCommunications,omitempty"`
		
		Custserv *bool `json:"custserv,omitempty"`
		*Alias
	}{ 
		RealtimeCIC: o.RealtimeCIC,
		
		Purecloud: o.Purecloud,
		
		Hipaa: o.Hipaa,
		
		UcEnabled: o.UcEnabled,
		
		Pci: o.Pci,
		
		PurecloudVoice: o.PurecloudVoice,
		
		XmppFederation: o.XmppFederation,
		
		Chat: o.Chat,
		
		InformalPhotos: o.InformalPhotos,
		
		Directory: o.Directory,
		
		ContactCenter: o.ContactCenter,
		
		UnifiedCommunications: o.UnifiedCommunications,
		
		Custserv: o.Custserv,
		Alias:    (*Alias)(o),
	})
}

func (o *Organizationfeatures) UnmarshalJSON(b []byte) error {
	var OrganizationfeaturesMap map[string]interface{}
	err := json.Unmarshal(b, &OrganizationfeaturesMap)
	if err != nil {
		return err
	}
	
	if RealtimeCIC, ok := OrganizationfeaturesMap["realtimeCIC"].(bool); ok {
		o.RealtimeCIC = &RealtimeCIC
	}
	
	if Purecloud, ok := OrganizationfeaturesMap["purecloud"].(bool); ok {
		o.Purecloud = &Purecloud
	}
	
	if Hipaa, ok := OrganizationfeaturesMap["hipaa"].(bool); ok {
		o.Hipaa = &Hipaa
	}
	
	if UcEnabled, ok := OrganizationfeaturesMap["ucEnabled"].(bool); ok {
		o.UcEnabled = &UcEnabled
	}
	
	if Pci, ok := OrganizationfeaturesMap["pci"].(bool); ok {
		o.Pci = &Pci
	}
	
	if PurecloudVoice, ok := OrganizationfeaturesMap["purecloudVoice"].(bool); ok {
		o.PurecloudVoice = &PurecloudVoice
	}
	
	if XmppFederation, ok := OrganizationfeaturesMap["xmppFederation"].(bool); ok {
		o.XmppFederation = &XmppFederation
	}
	
	if Chat, ok := OrganizationfeaturesMap["chat"].(bool); ok {
		o.Chat = &Chat
	}
	
	if InformalPhotos, ok := OrganizationfeaturesMap["informalPhotos"].(bool); ok {
		o.InformalPhotos = &InformalPhotos
	}
	
	if Directory, ok := OrganizationfeaturesMap["directory"].(bool); ok {
		o.Directory = &Directory
	}
	
	if ContactCenter, ok := OrganizationfeaturesMap["contactCenter"].(bool); ok {
		o.ContactCenter = &ContactCenter
	}
	
	if UnifiedCommunications, ok := OrganizationfeaturesMap["unifiedCommunications"].(bool); ok {
		o.UnifiedCommunications = &UnifiedCommunications
	}
	
	if Custserv, ok := OrganizationfeaturesMap["custserv"].(bool); ok {
		o.Custserv = &Custserv
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Organizationfeatures) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
