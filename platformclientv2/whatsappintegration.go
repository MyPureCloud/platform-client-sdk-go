package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappintegration
type Whatsappintegration struct { 
	// Id - A unique Integration Id.
	Id *string `json:"id,omitempty"`


	// Name - The name of the WhatsApp integration.
	Name *string `json:"name,omitempty"`


	// PhoneNumber - The phone number associated to the whatsApp integration.
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// Status - The status of the WhatsApp Integration
	Status *string `json:"status,omitempty"`


	// Recipient - The recipient associated to the WhatsApp Integration. This recipient is used to associate a flow to an integration
	Recipient *Domainentityref `json:"recipient,omitempty"`


	// DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this Integration was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User reference that created this Integration
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy - User reference that last modified this Integration
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Version - Version number required for updates.
	Version *int `json:"version,omitempty"`


	// ActivationStatusCode - The status code of WhatsApp Integration activation process
	ActivationStatusCode *string `json:"activationStatusCode,omitempty"`


	// ActivationErrorInfo - The error information of WhatsApp Integration activation process
	ActivationErrorInfo *Errorbody `json:"activationErrorInfo,omitempty"`


	// CreateStatus - Status of asynchronous create operation
	CreateStatus *string `json:"createStatus,omitempty"`


	// CreateError - Error information returned, if createStatus is set to Error
	CreateError *Errorbody `json:"createError,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Whatsappintegration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappintegration

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Recipient *Domainentityref `json:"recipient,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		ActivationStatusCode *string `json:"activationStatusCode,omitempty"`
		
		ActivationErrorInfo *Errorbody `json:"activationErrorInfo,omitempty"`
		
		CreateStatus *string `json:"createStatus,omitempty"`
		
		CreateError *Errorbody `json:"createError,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		PhoneNumber: u.PhoneNumber,
		
		Status: u.Status,
		
		Recipient: u.Recipient,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: u.CreatedBy,
		
		ModifiedBy: u.ModifiedBy,
		
		Version: u.Version,
		
		ActivationStatusCode: u.ActivationStatusCode,
		
		ActivationErrorInfo: u.ActivationErrorInfo,
		
		CreateStatus: u.CreateStatus,
		
		CreateError: u.CreateError,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Whatsappintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
