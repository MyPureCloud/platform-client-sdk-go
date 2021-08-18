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

func (u *Organizationfeatures) MarshalJSON() ([]byte, error) {
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
		RealtimeCIC: u.RealtimeCIC,
		
		Purecloud: u.Purecloud,
		
		Hipaa: u.Hipaa,
		
		UcEnabled: u.UcEnabled,
		
		Pci: u.Pci,
		
		PurecloudVoice: u.PurecloudVoice,
		
		XmppFederation: u.XmppFederation,
		
		Chat: u.Chat,
		
		InformalPhotos: u.InformalPhotos,
		
		Directory: u.Directory,
		
		ContactCenter: u.ContactCenter,
		
		UnifiedCommunications: u.UnifiedCommunications,
		
		Custserv: u.Custserv,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Organizationfeatures) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
