package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Limitchangerequestdetails
type Limitchangerequestdetails struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Key - Limit key to be overridden (see https://developer.mypurecloud.com/api/rest/v2/organization/limits.html#available_limits)
	Key *string `json:"key,omitempty"`


	// Namespace - Namespace the key belongs to (see https://developer.mypurecloud.com/api/rest/v2/organization/limits.html#available_limits)
	Namespace *string `json:"namespace,omitempty"`


	// RequestedValue - Requested limit value for a given key
	RequestedValue *float64 `json:"requestedValue,omitempty"`


	// Description - Description of the need for the limit change request
	Description *string `json:"description,omitempty"`


	// SupportCaseUrl - The support case url created by Care
	SupportCaseUrl *string `json:"supportCaseUrl,omitempty"`


	// CreatedBy - The user who created the change request
	CreatedBy *string `json:"createdBy,omitempty"`


	// Status - Current status of the limit change request
	Status *string `json:"status,omitempty"`


	// CurrentValue - Current limit value for a given key
	CurrentValue *float64 `json:"currentValue,omitempty"`


	// DateCreated - The date of the limit change request creation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// StatusHistory - List of statuses that a limit change request has gone through
	StatusHistory *[]Statuschange `json:"statusHistory,omitempty"`


	// DateCompleted - The date of the limit change request completion (ChangeImplemented, Rejected, or RollbackImplemented. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// LastChangedBy - The user who last updated the status of the limit change request
	LastChangedBy *string `json:"lastChangedBy,omitempty"`


	// RejectReason - The reason for rejecting the limit override request
	RejectReason *string `json:"rejectReason,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Limitchangerequestdetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Limitchangerequestdetails
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Key *string `json:"key,omitempty"`
		
		Namespace *string `json:"namespace,omitempty"`
		
		RequestedValue *float64 `json:"requestedValue,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SupportCaseUrl *string `json:"supportCaseUrl,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		CurrentValue *float64 `json:"currentValue,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		StatusHistory *[]Statuschange `json:"statusHistory,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		LastChangedBy *string `json:"lastChangedBy,omitempty"`
		
		RejectReason *string `json:"rejectReason,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Key: o.Key,
		
		Namespace: o.Namespace,
		
		RequestedValue: o.RequestedValue,
		
		Description: o.Description,
		
		SupportCaseUrl: o.SupportCaseUrl,
		
		CreatedBy: o.CreatedBy,
		
		Status: o.Status,
		
		CurrentValue: o.CurrentValue,
		
		DateCreated: DateCreated,
		
		StatusHistory: o.StatusHistory,
		
		DateCompleted: DateCompleted,
		
		LastChangedBy: o.LastChangedBy,
		
		RejectReason: o.RejectReason,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Limitchangerequestdetails) UnmarshalJSON(b []byte) error {
	var LimitchangerequestdetailsMap map[string]interface{}
	err := json.Unmarshal(b, &LimitchangerequestdetailsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LimitchangerequestdetailsMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Key, ok := LimitchangerequestdetailsMap["key"].(string); ok {
		o.Key = &Key
	}
	
	if Namespace, ok := LimitchangerequestdetailsMap["namespace"].(string); ok {
		o.Namespace = &Namespace
	}
	
	if RequestedValue, ok := LimitchangerequestdetailsMap["requestedValue"].(float64); ok {
		o.RequestedValue = &RequestedValue
	}
	
	if Description, ok := LimitchangerequestdetailsMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if SupportCaseUrl, ok := LimitchangerequestdetailsMap["supportCaseUrl"].(string); ok {
		o.SupportCaseUrl = &SupportCaseUrl
	}
	
	if CreatedBy, ok := LimitchangerequestdetailsMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if Status, ok := LimitchangerequestdetailsMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if CurrentValue, ok := LimitchangerequestdetailsMap["currentValue"].(float64); ok {
		o.CurrentValue = &CurrentValue
	}
	
	if dateCreatedString, ok := LimitchangerequestdetailsMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if StatusHistory, ok := LimitchangerequestdetailsMap["statusHistory"].([]interface{}); ok {
		StatusHistoryString, _ := json.Marshal(StatusHistory)
		json.Unmarshal(StatusHistoryString, &o.StatusHistory)
	}
	
	if dateCompletedString, ok := LimitchangerequestdetailsMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if LastChangedBy, ok := LimitchangerequestdetailsMap["lastChangedBy"].(string); ok {
		o.LastChangedBy = &LastChangedBy
	}
	
	if RejectReason, ok := LimitchangerequestdetailsMap["rejectReason"].(string); ok {
		o.RejectReason = &RejectReason
	}
	
	if SelfUri, ok := LimitchangerequestdetailsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Limitchangerequestdetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
