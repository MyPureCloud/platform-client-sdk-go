package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Organizationfeatures) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
