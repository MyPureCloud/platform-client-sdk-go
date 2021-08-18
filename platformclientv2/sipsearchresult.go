package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sipsearchresult
type Sipsearchresult struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Status - Status of the search request
	Status *int `json:"status,omitempty"`


	// Sid - Session id associated to the search request
	Sid *string `json:"sid,omitempty"`


	// Auth - Auth token used for this search request
	Auth *string `json:"auth,omitempty"`


	// Message - Any messages returned from homer as part of the response
	Message *string `json:"message,omitempty"`


	// Data - Homer search data that is returned
	Data *[]Homerrecord `json:"data,omitempty"`


	// Count - Number of records returned
	Count *int `json:"count,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Sipsearchresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sipsearchresult

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		Sid *string `json:"sid,omitempty"`
		
		Auth *string `json:"auth,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		Data *[]Homerrecord `json:"data,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Status: u.Status,
		
		Sid: u.Sid,
		
		Auth: u.Auth,
		
		Message: u.Message,
		
		Data: u.Data,
		
		Count: u.Count,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Sipsearchresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
