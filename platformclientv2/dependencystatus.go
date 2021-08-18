package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dependencystatus
type Dependencystatus struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User - User that initiated the build.
	User *User `json:"user,omitempty"`


	// Client - OAuth client that initiated the build.
	Client *Domainentityref `json:"client,omitempty"`


	// BuildId
	BuildId *string `json:"buildId,omitempty"`


	// DateStarted - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`


	// DateCompleted - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// FailedObjects
	FailedObjects *[]Failedobject `json:"failedObjects,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Dependencystatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dependencystatus

	
	DateStarted := new(string)
	if u.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(u.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	DateCompleted := new(string)
	if u.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(u.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Client *Domainentityref `json:"client,omitempty"`
		
		BuildId *string `json:"buildId,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		FailedObjects *[]Failedobject `json:"failedObjects,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		User: u.User,
		
		Client: u.Client,
		
		BuildId: u.BuildId,
		
		DateStarted: DateStarted,
		
		DateCompleted: DateCompleted,
		
		Status: u.Status,
		
		FailedObjects: u.FailedObjects,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dependencystatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
