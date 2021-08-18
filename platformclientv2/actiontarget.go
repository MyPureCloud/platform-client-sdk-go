package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actiontarget
type Actiontarget struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// UserData - Additional user data associated with the target in key/value format.
	UserData *[]Keyvalue `json:"userData,omitempty"`


	// SupportedMediaTypes - Supported media types of the target.
	SupportedMediaTypes *[]string `json:"supportedMediaTypes,omitempty"`


	// State - Indicates the state of the target.
	State *string `json:"state,omitempty"`


	// Description - Description of the target.
	Description *string `json:"description,omitempty"`


	// ServiceLevel - Service Level of the action target. Chat offers for the target will be throttled with the aim of achieving this service level.
	ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`


	// ShortAbandonThreshold - Indicates the non-default short abandon threshold
	ShortAbandonThreshold *int `json:"shortAbandonThreshold,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate - The date the target was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - The date the target was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (u *Actiontarget) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actiontarget

	
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
		
		UserData *[]Keyvalue `json:"userData,omitempty"`
		
		SupportedMediaTypes *[]string `json:"supportedMediaTypes,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`
		
		ShortAbandonThreshold *int `json:"shortAbandonThreshold,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		UserData: u.UserData,
		
		SupportedMediaTypes: u.SupportedMediaTypes,
		
		State: u.State,
		
		Description: u.Description,
		
		ServiceLevel: u.ServiceLevel,
		
		ShortAbandonThreshold: u.ShortAbandonThreshold,
		
		SelfUri: u.SelfUri,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actiontarget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
