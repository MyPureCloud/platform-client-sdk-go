package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Createwebchatconversationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createwebchatconversationrequest

	

	return json.Marshal(&struct { 
		OrganizationId *string `json:"organizationId,omitempty"`
		
		DeploymentId *string `json:"deploymentId,omitempty"`
		
		RoutingTarget *Webchatroutingtarget `json:"routingTarget,omitempty"`
		
		MemberInfo *Guestmemberinfo `json:"memberInfo,omitempty"`
		
		MemberAuthToken *string `json:"memberAuthToken,omitempty"`
		
		JourneyContext *Journeycontext `json:"journeyContext,omitempty"`
		*Alias
	}{ 
		OrganizationId: u.OrganizationId,
		
		DeploymentId: u.DeploymentId,
		
		RoutingTarget: u.RoutingTarget,
		
		MemberInfo: u.MemberInfo,
		
		MemberAuthToken: u.MemberAuthToken,
		
		JourneyContext: u.JourneyContext,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createwebchatconversationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
