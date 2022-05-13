package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustee
type Trustee struct { 
	// Id - Organization Id for this trust.
	Id *string `json:"id,omitempty"`


	// Enabled - If disabled no trustee user will have access, even if they were previously added.
	Enabled *bool `json:"enabled,omitempty"`


	// UsesDefaultRole - Denotes if trustee uses admin role by default.
	UsesDefaultRole *bool `json:"usesDefaultRole,omitempty"`


	// DateCreated - Date Trust was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateExpired - The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpired *time.Time `json:"dateExpired,omitempty"`


	// CreatedBy - User that created trust.
	CreatedBy *Orguser `json:"createdBy,omitempty"`


	// Organization - Organization associated with this trust.
	Organization *Organization `json:"organization,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Trustee) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustee
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateExpired := new(string)
	if o.DateExpired != nil {
		
		*DateExpired = timeutil.Strftime(o.DateExpired, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpired = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		UsesDefaultRole *bool `json:"usesDefaultRole,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateExpired *string `json:"dateExpired,omitempty"`
		
		CreatedBy *Orguser `json:"createdBy,omitempty"`
		
		Organization *Organization `json:"organization,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Enabled: o.Enabled,
		
		UsesDefaultRole: o.UsesDefaultRole,
		
		DateCreated: DateCreated,
		
		DateExpired: DateExpired,
		
		CreatedBy: o.CreatedBy,
		
		Organization: o.Organization,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Trustee) UnmarshalJSON(b []byte) error {
	var TrusteeMap map[string]interface{}
	err := json.Unmarshal(b, &TrusteeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TrusteeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Enabled, ok := TrusteeMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if UsesDefaultRole, ok := TrusteeMap["usesDefaultRole"].(bool); ok {
		o.UsesDefaultRole = &UsesDefaultRole
	}
    
	if dateCreatedString, ok := TrusteeMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateExpiredString, ok := TrusteeMap["dateExpired"].(string); ok {
		DateExpired, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiredString)
		o.DateExpired = &DateExpired
	}
	
	if CreatedBy, ok := TrusteeMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if Organization, ok := TrusteeMap["organization"].(map[string]interface{}); ok {
		OrganizationString, _ := json.Marshal(Organization)
		json.Unmarshal(OrganizationString, &o.Organization)
	}
	
	if SelfUri, ok := TrusteeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Trustee) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
