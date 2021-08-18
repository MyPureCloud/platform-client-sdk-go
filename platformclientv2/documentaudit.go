package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentaudit
type Documentaudit struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User
	User *Domainentityref `json:"user,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// TransactionId
	TransactionId *string `json:"transactionId,omitempty"`


	// TransactionInitiator
	TransactionInitiator *bool `json:"transactionInitiator,omitempty"`


	// Application
	Application *string `json:"application,omitempty"`


	// ServiceName
	ServiceName *string `json:"serviceName,omitempty"`


	// Level
	Level *string `json:"level,omitempty"`


	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// ActionContext
	ActionContext *string `json:"actionContext,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`


	// Entity
	Entity *Auditentityreference `json:"entity,omitempty"`


	// Changes
	Changes *[]Auditchange `json:"changes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Documentaudit) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentaudit

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User *Domainentityref `json:"user,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		TransactionId *string `json:"transactionId,omitempty"`
		
		TransactionInitiator *bool `json:"transactionInitiator,omitempty"`
		
		Application *string `json:"application,omitempty"`
		
		ServiceName *string `json:"serviceName,omitempty"`
		
		Level *string `json:"level,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ActionContext *string `json:"actionContext,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		Entity *Auditentityreference `json:"entity,omitempty"`
		
		Changes *[]Auditchange `json:"changes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		User: u.User,
		
		Workspace: u.Workspace,
		
		TransactionId: u.TransactionId,
		
		TransactionInitiator: u.TransactionInitiator,
		
		Application: u.Application,
		
		ServiceName: u.ServiceName,
		
		Level: u.Level,
		
		Timestamp: Timestamp,
		
		Status: u.Status,
		
		ActionContext: u.ActionContext,
		
		Action: u.Action,
		
		Entity: u.Entity,
		
		Changes: u.Changes,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Documentaudit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
