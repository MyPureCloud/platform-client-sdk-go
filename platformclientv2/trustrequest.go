package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustrequest
type Trustrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// CreatedBy - User who created this request.
	CreatedBy *Orguser `json:"createdBy,omitempty"`


	// DateCreated - Date request was created. There is a 48 hour expiration on all requests. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// Trustee - Trustee organization who generated this request.
	Trustee *Organization `json:"trustee,omitempty"`


	// Users - The list of trustee users that are requesting access.
	Users *[]Orguser `json:"users,omitempty"`


	// Groups - The list of trustee groups that are requesting access.
	Groups *[]Trustgroup `json:"groups,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Trustrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustrequest
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		CreatedBy *Orguser `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		Trustee *Organization `json:"trustee,omitempty"`
		
		Users *[]Orguser `json:"users,omitempty"`
		
		Groups *[]Trustgroup `json:"groups,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		Trustee: o.Trustee,
		
		Users: o.Users,
		
		Groups: o.Groups,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Trustrequest) UnmarshalJSON(b []byte) error {
	var TrustrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TrustrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TrustrequestMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if CreatedBy, ok := TrustrequestMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := TrustrequestMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if Trustee, ok := TrustrequestMap["trustee"].(map[string]interface{}); ok {
		TrusteeString, _ := json.Marshal(Trustee)
		json.Unmarshal(TrusteeString, &o.Trustee)
	}
	
	if Users, ok := TrustrequestMap["users"].([]interface{}); ok {
		UsersString, _ := json.Marshal(Users)
		json.Unmarshal(UsersString, &o.Users)
	}
	
	if Groups, ok := TrustrequestMap["groups"].([]interface{}); ok {
		GroupsString, _ := json.Marshal(Groups)
		json.Unmarshal(GroupsString, &o.Groups)
	}
	
	if SelfUri, ok := TrustrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trustrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
