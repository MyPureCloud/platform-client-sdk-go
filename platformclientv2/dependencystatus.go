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

func (o *Dependencystatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dependencystatus
	
	DateStarted := new(string)
	if o.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(o.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		User: o.User,
		
		Client: o.Client,
		
		BuildId: o.BuildId,
		
		DateStarted: DateStarted,
		
		DateCompleted: DateCompleted,
		
		Status: o.Status,
		
		FailedObjects: o.FailedObjects,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Dependencystatus) UnmarshalJSON(b []byte) error {
	var DependencystatusMap map[string]interface{}
	err := json.Unmarshal(b, &DependencystatusMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DependencystatusMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DependencystatusMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if User, ok := DependencystatusMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := DependencystatusMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if BuildId, ok := DependencystatusMap["buildId"].(string); ok {
		o.BuildId = &BuildId
	}
	
	if dateStartedString, ok := DependencystatusMap["dateStarted"].(string); ok {
		DateStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartedString)
		o.DateStarted = &DateStarted
	}
	
	if dateCompletedString, ok := DependencystatusMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if Status, ok := DependencystatusMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if FailedObjects, ok := DependencystatusMap["failedObjects"].([]interface{}); ok {
		FailedObjectsString, _ := json.Marshal(FailedObjects)
		json.Unmarshal(FailedObjectsString, &o.FailedObjects)
	}
	
	if SelfUri, ok := DependencystatusMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dependencystatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
