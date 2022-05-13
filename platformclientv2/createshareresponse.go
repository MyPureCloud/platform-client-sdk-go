package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createshareresponse
type Createshareresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SharedEntityType
	SharedEntityType *string `json:"sharedEntityType,omitempty"`


	// SharedEntity
	SharedEntity *Domainentityref `json:"sharedEntity,omitempty"`


	// MemberType
	MemberType *string `json:"memberType,omitempty"`


	// Member
	Member *Domainentityref `json:"member,omitempty"`


	// SharedBy
	SharedBy *Domainentityref `json:"sharedBy,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// Succeeded
	Succeeded *[]Share `json:"succeeded,omitempty"`


	// Failed
	Failed *[]Share `json:"failed,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Createshareresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createshareresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SharedEntityType *string `json:"sharedEntityType,omitempty"`
		
		SharedEntity *Domainentityref `json:"sharedEntity,omitempty"`
		
		MemberType *string `json:"memberType,omitempty"`
		
		Member *Domainentityref `json:"member,omitempty"`
		
		SharedBy *Domainentityref `json:"sharedBy,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		Succeeded *[]Share `json:"succeeded,omitempty"`
		
		Failed *[]Share `json:"failed,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SharedEntityType: o.SharedEntityType,
		
		SharedEntity: o.SharedEntity,
		
		MemberType: o.MemberType,
		
		Member: o.Member,
		
		SharedBy: o.SharedBy,
		
		Workspace: o.Workspace,
		
		Succeeded: o.Succeeded,
		
		Failed: o.Failed,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Createshareresponse) UnmarshalJSON(b []byte) error {
	var CreateshareresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CreateshareresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CreateshareresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CreateshareresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SharedEntityType, ok := CreateshareresponseMap["sharedEntityType"].(string); ok {
		o.SharedEntityType = &SharedEntityType
	}
    
	if SharedEntity, ok := CreateshareresponseMap["sharedEntity"].(map[string]interface{}); ok {
		SharedEntityString, _ := json.Marshal(SharedEntity)
		json.Unmarshal(SharedEntityString, &o.SharedEntity)
	}
	
	if MemberType, ok := CreateshareresponseMap["memberType"].(string); ok {
		o.MemberType = &MemberType
	}
    
	if Member, ok := CreateshareresponseMap["member"].(map[string]interface{}); ok {
		MemberString, _ := json.Marshal(Member)
		json.Unmarshal(MemberString, &o.Member)
	}
	
	if SharedBy, ok := CreateshareresponseMap["sharedBy"].(map[string]interface{}); ok {
		SharedByString, _ := json.Marshal(SharedBy)
		json.Unmarshal(SharedByString, &o.SharedBy)
	}
	
	if Workspace, ok := CreateshareresponseMap["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if Succeeded, ok := CreateshareresponseMap["succeeded"].([]interface{}); ok {
		SucceededString, _ := json.Marshal(Succeeded)
		json.Unmarshal(SucceededString, &o.Succeeded)
	}
	
	if Failed, ok := CreateshareresponseMap["failed"].([]interface{}); ok {
		FailedString, _ := json.Marshal(Failed)
		json.Unmarshal(FailedString, &o.Failed)
	}
	
	if SelfUri, ok := CreateshareresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createshareresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
