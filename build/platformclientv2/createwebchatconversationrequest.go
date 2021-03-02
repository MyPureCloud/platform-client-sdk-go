package platformclientv2
import (
	"encoding/json"
)

// Createwebchatconversationrequest
type Createwebchatconversationrequest struct { 
	// OrganizationId - The organization identifier.
	OrganizationId *string `json:"organizationId,omitempty"`


	// DeploymentId - The web chat Deployment ID which contains the appropriate settings for this chat conversation.
	DeploymentId *string `json:"deploymentId,omitempty"`


	// RoutingTarget - The routing information to use for the new chat conversation.
	RoutingTarget *Webchatroutingtarget `json:"routingTarget,omitempty"`


	// MemberInfo - The guest member info to use for the new chat conversation.
	MemberInfo *Guestmemberinfo `json:"memberInfo,omitempty"`


	// MemberAuthToken - If the guest member is an authenticated member (ie, not anonymous) his JWT is provided here. The token will have been previously generated with the \"POST /api/v2/signeddata\" resource.
	MemberAuthToken *string `json:"memberAuthToken,omitempty"`


	// JourneyContext - A subset of the Journey System's data relevant to this conversation/session request (for external linkage and internal usage/context).
	JourneyContext *Journeycontext `json:"journeyContext,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createwebchatconversationrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
