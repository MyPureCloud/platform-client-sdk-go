package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Station) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
