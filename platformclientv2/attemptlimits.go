package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Attemptlimits
type Attemptlimits struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// MaxAttemptsPerContact - The maximum number of times a contact can be called within the resetPeriod. Required if maxAttemptsPerNumber is not defined.
	MaxAttemptsPerContact *int `json:"maxAttemptsPerContact,omitempty"`


	// MaxAttemptsPerNumber - The maximum number of times a phone number can be called within the resetPeriod. Required if maxAttemptsPerContact is not defined.
	MaxAttemptsPerNumber *int `json:"maxAttemptsPerNumber,omitempty"`


	// TimeZoneId - If the resetPeriod is TODAY, this specifies the timezone in which TODAY occurs. Required if the resetPeriod is TODAY.
	TimeZoneId *string `json:"timeZoneId,omitempty"`


	// ResetPeriod - After how long the number of attempts will be set back to 0. Defaults to NEVER.
	ResetPeriod *string `json:"resetPeriod,omitempty"`


	// RecallEntries - Configuration for recall attempts.
	RecallEntries *map[string]Recallentry `json:"recallEntries,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Attemptlimits) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Attemptlimits

	
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
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		MaxAttemptsPerContact *int `json:"maxAttemptsPerContact,omitempty"`
		
		MaxAttemptsPerNumber *int `json:"maxAttemptsPerNumber,omitempty"`
		
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		
		ResetPeriod *string `json:"resetPeriod,omitempty"`
		
		RecallEntries *map[string]Recallentry `json:"recallEntries,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: u.Version,
		
		MaxAttemptsPerContact: u.MaxAttemptsPerContact,
		
		MaxAttemptsPerNumber: u.MaxAttemptsPerNumber,
		
		TimeZoneId: u.TimeZoneId,
		
		ResetPeriod: u.ResetPeriod,
		
		RecallEntries: u.RecallEntries,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Attemptlimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
