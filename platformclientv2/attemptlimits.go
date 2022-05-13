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

func (o *Attemptlimits) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Attemptlimits
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		MaxAttemptsPerContact: o.MaxAttemptsPerContact,
		
		MaxAttemptsPerNumber: o.MaxAttemptsPerNumber,
		
		TimeZoneId: o.TimeZoneId,
		
		ResetPeriod: o.ResetPeriod,
		
		RecallEntries: o.RecallEntries,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Attemptlimits) UnmarshalJSON(b []byte) error {
	var AttemptlimitsMap map[string]interface{}
	err := json.Unmarshal(b, &AttemptlimitsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AttemptlimitsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AttemptlimitsMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := AttemptlimitsMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := AttemptlimitsMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := AttemptlimitsMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if MaxAttemptsPerContact, ok := AttemptlimitsMap["maxAttemptsPerContact"].(float64); ok {
		MaxAttemptsPerContactInt := int(MaxAttemptsPerContact)
		o.MaxAttemptsPerContact = &MaxAttemptsPerContactInt
	}
	
	if MaxAttemptsPerNumber, ok := AttemptlimitsMap["maxAttemptsPerNumber"].(float64); ok {
		MaxAttemptsPerNumberInt := int(MaxAttemptsPerNumber)
		o.MaxAttemptsPerNumber = &MaxAttemptsPerNumberInt
	}
	
	if TimeZoneId, ok := AttemptlimitsMap["timeZoneId"].(string); ok {
		o.TimeZoneId = &TimeZoneId
	}
    
	if ResetPeriod, ok := AttemptlimitsMap["resetPeriod"].(string); ok {
		o.ResetPeriod = &ResetPeriod
	}
    
	if RecallEntries, ok := AttemptlimitsMap["recallEntries"].(map[string]interface{}); ok {
		RecallEntriesString, _ := json.Marshal(RecallEntries)
		json.Unmarshal(RecallEntriesString, &o.RecallEntries)
	}
	
	if SelfUri, ok := AttemptlimitsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Attemptlimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
