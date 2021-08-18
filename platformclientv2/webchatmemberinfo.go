package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Webchatmemberinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatmemberinfo

	
	JoinDate := new(string)
	if u.JoinDate != nil {
		
		*JoinDate = timeutil.Strftime(u.JoinDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		JoinDate = nil
	}
	
	LeaveDate := new(string)
	if u.LeaveDate != nil {
		
		*LeaveDate = timeutil.Strftime(u.LeaveDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LeaveDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`
		
		Role *string `json:"role,omitempty"`
		
		JoinDate *string `json:"joinDate,omitempty"`
		
		LeaveDate *string `json:"leaveDate,omitempty"`
		
		AuthenticatedGuest *bool `json:"authenticatedGuest,omitempty"`
		
		CustomFields *map[string]string `json:"customFields,omitempty"`
		
		State *string `json:"state,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		DisplayName: u.DisplayName,
		
		FirstName: u.FirstName,
		
		LastName: u.LastName,
		
		Email: u.Email,
		
		PhoneNumber: u.PhoneNumber,
		
		AvatarImageUrl: u.AvatarImageUrl,
		
		Role: u.Role,
		
		JoinDate: JoinDate,
		
		LeaveDate: LeaveDate,
		
		AuthenticatedGuest: u.AuthenticatedGuest,
		
		CustomFields: u.CustomFields,
		
		State: u.State,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webchatmemberinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
