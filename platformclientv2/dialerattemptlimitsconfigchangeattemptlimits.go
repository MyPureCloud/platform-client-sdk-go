package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerattemptlimitsconfigchangeattemptlimits
type Dialerattemptlimitsconfigchangeattemptlimits struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// MaxAttemptsPerContact
	MaxAttemptsPerContact *int `json:"maxAttemptsPerContact,omitempty"`


	// MaxAttemptsPerNumber
	MaxAttemptsPerNumber *int `json:"maxAttemptsPerNumber,omitempty"`


	// TimeZoneId
	TimeZoneId *string `json:"timeZoneId,omitempty"`


	// ResetPeriod
	ResetPeriod *string `json:"resetPeriod,omitempty"`


	// RecallEntries
	RecallEntries *map[string]Dialerattemptlimitsconfigchangerecallentry `json:"recallEntries,omitempty"`


	// BreadthFirstRecalls
	BreadthFirstRecalls *bool `json:"breadthFirstRecalls,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialerattemptlimitsconfigchangeattemptlimits) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerattemptlimitsconfigchangeattemptlimits

	
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
		
		RecallEntries *map[string]Dialerattemptlimitsconfigchangerecallentry `json:"recallEntries,omitempty"`
		
		BreadthFirstRecalls *bool `json:"breadthFirstRecalls,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
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
		
		BreadthFirstRecalls: u.BreadthFirstRecalls,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialerattemptlimitsconfigchangeattemptlimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
