package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Station
type Station struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// UserId - The Id of the user currently logged in and associated with the station.
	UserId *string `json:"userId,omitempty"`


	// WebRtcUserId - The Id of the user configured for the station if it is of type inin_webrtc_softphone. Empty if station type is not inin_webrtc_softphone.
	WebRtcUserId *string `json:"webRtcUserId,omitempty"`


	// PrimaryEdge
	PrimaryEdge *Domainentityref `json:"primaryEdge,omitempty"`


	// SecondaryEdge
	SecondaryEdge *Domainentityref `json:"secondaryEdge,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// LineAppearanceId
	LineAppearanceId *string `json:"lineAppearanceId,omitempty"`


	// WebRtcMediaDscp - The default or configured value of media dscp for the station. Empty if station type is not inin_webrtc_softphone.
	WebRtcMediaDscp *int `json:"webRtcMediaDscp,omitempty"`


	// WebRtcPersistentEnabled - The default or configured value of persistent connection setting for the station. Empty if station type is not inin_webrtc_softphone.
	WebRtcPersistentEnabled *bool `json:"webRtcPersistentEnabled,omitempty"`


	// WebRtcForceTurn - Whether the station is configured to require TURN for routing WebRTC calls. Empty if station type is not inin_webrtc_softphone.
	WebRtcForceTurn *bool `json:"webRtcForceTurn,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Station) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Station

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		WebRtcUserId *string `json:"webRtcUserId,omitempty"`
		
		PrimaryEdge *Domainentityref `json:"primaryEdge,omitempty"`
		
		SecondaryEdge *Domainentityref `json:"secondaryEdge,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		LineAppearanceId *string `json:"lineAppearanceId,omitempty"`
		
		WebRtcMediaDscp *int `json:"webRtcMediaDscp,omitempty"`
		
		WebRtcPersistentEnabled *bool `json:"webRtcPersistentEnabled,omitempty"`
		
		WebRtcForceTurn *bool `json:"webRtcForceTurn,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Status: u.Status,
		
		UserId: u.UserId,
		
		WebRtcUserId: u.WebRtcUserId,
		
		PrimaryEdge: u.PrimaryEdge,
		
		SecondaryEdge: u.SecondaryEdge,
		
		VarType: u.VarType,
		
		LineAppearanceId: u.LineAppearanceId,
		
		WebRtcMediaDscp: u.WebRtcMediaDscp,
		
		WebRtcPersistentEnabled: u.WebRtcPersistentEnabled,
		
		WebRtcForceTurn: u.WebRtcForceTurn,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Station) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
