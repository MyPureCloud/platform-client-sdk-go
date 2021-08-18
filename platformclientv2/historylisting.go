package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historylisting
type Historylisting struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Complete
	Complete *bool `json:"complete,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// Client
	Client *Domainentityref `json:"client,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorDetails
	ErrorDetails *[]Detail `json:"errorDetails,omitempty"`


	// ErrorMessageParams
	ErrorMessageParams *map[string]string `json:"errorMessageParams,omitempty"`


	// ActionName - Action name
	ActionName *string `json:"actionName,omitempty"`


	// ActionStatus - Action status
	ActionStatus *string `json:"actionStatus,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// System
	System *bool `json:"system,omitempty"`


	// Started - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Started *time.Time `json:"started,omitempty"`


	// Completed - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Completed *time.Time `json:"completed,omitempty"`


	// Entities
	Entities *[]Historyentry `json:"entities,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

func (u *Historylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historylisting

	
	Started := new(string)
	if u.Started != nil {
		
		*Started = timeutil.Strftime(u.Started, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Started = nil
	}
	
	Completed := new(string)
	if u.Completed != nil {
		
		*Completed = timeutil.Strftime(u.Completed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Completed = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Complete *bool `json:"complete,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Client *Domainentityref `json:"client,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorDetails *[]Detail `json:"errorDetails,omitempty"`
		
		ErrorMessageParams *map[string]string `json:"errorMessageParams,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionStatus *string `json:"actionStatus,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		System *bool `json:"system,omitempty"`
		
		Started *string `json:"started,omitempty"`
		
		Completed *string `json:"completed,omitempty"`
		
		Entities *[]Historyentry `json:"entities,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Complete: u.Complete,
		
		User: u.User,
		
		Client: u.Client,
		
		ErrorMessage: u.ErrorMessage,
		
		ErrorCode: u.ErrorCode,
		
		ErrorDetails: u.ErrorDetails,
		
		ErrorMessageParams: u.ErrorMessageParams,
		
		ActionName: u.ActionName,
		
		ActionStatus: u.ActionStatus,
		
		Name: u.Name,
		
		Description: u.Description,
		
		System: u.System,
		
		Started: Started,
		
		Completed: Completed,
		
		Entities: u.Entities,
		
		Total: u.Total,
		
		PageSize: u.PageSize,
		
		PageNumber: u.PageNumber,
		
		PageCount: u.PageCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Historylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
