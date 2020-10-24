package platformclientv2
import (
	"time"
	"encoding/json"
)

// Webchatmemberinfo
type Webchatmemberinfo struct { 
	// Id - The communicationId of this member.
	Id *string `json:"id,omitempty"`


	// DisplayName - The display name of the member.
	DisplayName *string `json:"displayName,omitempty"`


	// FirstName - The first name of the member.
	FirstName *string `json:"firstName,omitempty"`


	// LastName - The last name of the member.
	LastName *string `json:"lastName,omitempty"`


	// Email - The email address of the member.
	Email *string `json:"email,omitempty"`


	// PhoneNumber - The phone number of the member.
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// AvatarImageUrl - The url to the avatar image of the member.
	AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`


	// Role - The role of the member, one of [agent, customer, acd, workflow]
	Role *string `json:"role,omitempty"`


	// JoinDate - The time the member joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	JoinDate *time.Time `json:"joinDate,omitempty"`


	// LeaveDate - The time the member left the conversation, or null if the member is still active in the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	LeaveDate *time.Time `json:"leaveDate,omitempty"`


	// AuthenticatedGuest - If true, the guest member is an authenticated guest.
	AuthenticatedGuest *bool `json:"authenticatedGuest,omitempty"`


	// CustomFields - Any custom fields of information pertaining to this member.
	CustomFields *map[string]string `json:"customFields,omitempty"`


	// State - The connection state of this member.
	State *string `json:"state,omitempty"`

}

// String returns a JSON representation of the model
func (o *Webchatmemberinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
