package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Group
type Group struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The group name.
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// DateModified - Last modified date/time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// MemberCount - Number of members.
	MemberCount *int `json:"memberCount,omitempty"`


	// State - Active, inactive, or deleted state.
	State *string `json:"state,omitempty"`


	// Version - Current version for this resource.
	Version *int `json:"version,omitempty"`


	// VarType - Type of group.
	VarType *string `json:"type,omitempty"`


	// Images
	Images *[]Userimage `json:"images,omitempty"`


	// Addresses
	Addresses *[]Groupcontact `json:"addresses,omitempty"`


	// RulesVisible - Are membership rules visible to the person requesting to view the group
	RulesVisible *bool `json:"rulesVisible,omitempty"`


	// Visibility - Who can view this group
	Visibility *string `json:"visibility,omitempty"`


	// Owners - Owners of the group
	Owners *[]User `json:"owners,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Group) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Group

	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		MemberCount *int `json:"memberCount,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Images *[]Userimage `json:"images,omitempty"`
		
		Addresses *[]Groupcontact `json:"addresses,omitempty"`
		
		RulesVisible *bool `json:"rulesVisible,omitempty"`
		
		Visibility *string `json:"visibility,omitempty"`
		
		Owners *[]User `json:"owners,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		DateModified: DateModified,
		
		MemberCount: u.MemberCount,
		
		State: u.State,
		
		Version: u.Version,
		
		VarType: u.VarType,
		
		Images: u.Images,
		
		Addresses: u.Addresses,
		
		RulesVisible: u.RulesVisible,
		
		Visibility: u.Visibility,
		
		Owners: u.Owners,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Group) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
