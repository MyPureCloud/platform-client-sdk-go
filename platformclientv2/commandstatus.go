package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Commandstatus
type Commandstatus struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Expiration - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Expiration *time.Time `json:"expiration,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// StatusCode
	StatusCode *string `json:"statusCode,omitempty"`


	// CommandType
	CommandType *string `json:"commandType,omitempty"`


	// Document
	Document *Document `json:"document,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Commandstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Commandstatus

	
	Expiration := new(string)
	if u.Expiration != nil {
		
		*Expiration = timeutil.Strftime(u.Expiration, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Expiration = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Expiration *string `json:"expiration,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		StatusCode *string `json:"statusCode,omitempty"`
		
		CommandType *string `json:"commandType,omitempty"`
		
		Document *Document `json:"document,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Expiration: Expiration,
		
		UserId: u.UserId,
		
		StatusCode: u.StatusCode,
		
		CommandType: u.CommandType,
		
		Document: u.Document,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Commandstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
