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

func (o *Group) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Group
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DateModified: DateModified,
		
		MemberCount: o.MemberCount,
		
		State: o.State,
		
		Version: o.Version,
		
		VarType: o.VarType,
		
		Images: o.Images,
		
		Addresses: o.Addresses,
		
		RulesVisible: o.RulesVisible,
		
		Visibility: o.Visibility,
		
		Owners: o.Owners,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Group) UnmarshalJSON(b []byte) error {
	var GroupMap map[string]interface{}
	err := json.Unmarshal(b, &GroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GroupMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := GroupMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := GroupMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if dateModifiedString, ok := GroupMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if MemberCount, ok := GroupMap["memberCount"].(float64); ok {
		MemberCountInt := int(MemberCount)
		o.MemberCount = &MemberCountInt
	}
	
	if State, ok := GroupMap["state"].(string); ok {
		o.State = &State
	}
	
	if Version, ok := GroupMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if VarType, ok := GroupMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Images, ok := GroupMap["images"].([]interface{}); ok {
		ImagesString, _ := json.Marshal(Images)
		json.Unmarshal(ImagesString, &o.Images)
	}
	
	if Addresses, ok := GroupMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if RulesVisible, ok := GroupMap["rulesVisible"].(bool); ok {
		o.RulesVisible = &RulesVisible
	}
	
	if Visibility, ok := GroupMap["visibility"].(string); ok {
		o.Visibility = &Visibility
	}
	
	if Owners, ok := GroupMap["owners"].([]interface{}); ok {
		OwnersString, _ := json.Marshal(Owners)
		json.Unmarshal(OwnersString, &o.Owners)
	}
	
	if SelfUri, ok := GroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Group) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
