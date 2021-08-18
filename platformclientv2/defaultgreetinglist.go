package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Defaultgreetinglist
type Defaultgreetinglist struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Owner
	Owner *Greetingowner `json:"owner,omitempty"`


	// OwnerType
	OwnerType *string `json:"ownerType,omitempty"`


	// Greetings
	Greetings *map[string]Greeting `json:"greetings,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// CreatedBy
	CreatedBy *string `json:"createdBy,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// ModifiedBy
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Defaultgreetinglist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Defaultgreetinglist

	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Owner *Greetingowner `json:"owner,omitempty"`
		
		OwnerType *string `json:"ownerType,omitempty"`
		
		Greetings *map[string]Greeting `json:"greetings,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Owner: u.Owner,
		
		OwnerType: u.OwnerType,
		
		Greetings: u.Greetings,
		
		CreatedDate: CreatedDate,
		
		CreatedBy: u.CreatedBy,
		
		ModifiedDate: ModifiedDate,
		
		ModifiedBy: u.ModifiedBy,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Defaultgreetinglist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
