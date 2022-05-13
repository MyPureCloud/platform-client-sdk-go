package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workspace
type Workspace struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The current name of the workspace.
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// IsCurrentUserWorkspace
	IsCurrentUserWorkspace *bool `json:"isCurrentUserWorkspace,omitempty"`


	// User
	User *Domainentityref `json:"user,omitempty"`


	// Bucket
	Bucket *string `json:"bucket,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Summary
	Summary *Workspacesummary `json:"summary,omitempty"`


	// Acl
	Acl *[]string `json:"acl,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Workspace) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workspace
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		IsCurrentUserWorkspace *bool `json:"isCurrentUserWorkspace,omitempty"`
		
		User *Domainentityref `json:"user,omitempty"`
		
		Bucket *string `json:"bucket,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Summary *Workspacesummary `json:"summary,omitempty"`
		
		Acl *[]string `json:"acl,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		IsCurrentUserWorkspace: o.IsCurrentUserWorkspace,
		
		User: o.User,
		
		Bucket: o.Bucket,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Summary: o.Summary,
		
		Acl: o.Acl,
		
		Description: o.Description,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Workspace) UnmarshalJSON(b []byte) error {
	var WorkspaceMap map[string]interface{}
	err := json.Unmarshal(b, &WorkspaceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkspaceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkspaceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := WorkspaceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if IsCurrentUserWorkspace, ok := WorkspaceMap["isCurrentUserWorkspace"].(bool); ok {
		o.IsCurrentUserWorkspace = &IsCurrentUserWorkspace
	}
    
	if User, ok := WorkspaceMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Bucket, ok := WorkspaceMap["bucket"].(string); ok {
		o.Bucket = &Bucket
	}
    
	if dateCreatedString, ok := WorkspaceMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := WorkspaceMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Summary, ok := WorkspaceMap["summary"].(map[string]interface{}); ok {
		SummaryString, _ := json.Marshal(Summary)
		json.Unmarshal(SummaryString, &o.Summary)
	}
	
	if Acl, ok := WorkspaceMap["acl"].([]interface{}); ok {
		AclString, _ := json.Marshal(Acl)
		json.Unmarshal(AclString, &o.Acl)
	}
	
	if Description, ok := WorkspaceMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if SelfUri, ok := WorkspaceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workspace) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
