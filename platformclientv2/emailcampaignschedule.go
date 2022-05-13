package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailcampaignschedule
type Emailcampaignschedule struct { 
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


	// Intervals - A list of intervals during which to run the associated Campaign.
	Intervals *[]Scheduleinterval `json:"intervals,omitempty"`


	// TimeZone - The time zone for this email campaign schedule.
	TimeZone *string `json:"timeZone,omitempty"`


	// EmailCampaign - The Campaign that this email campaign schedule is for.
	EmailCampaign *Domainentityref `json:"emailCampaign,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Emailcampaignschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailcampaignschedule
	
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
		
		Intervals *[]Scheduleinterval `json:"intervals,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		EmailCampaign *Domainentityref `json:"emailCampaign,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		Intervals: o.Intervals,
		
		TimeZone: o.TimeZone,
		
		EmailCampaign: o.EmailCampaign,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Emailcampaignschedule) UnmarshalJSON(b []byte) error {
	var EmailcampaignscheduleMap map[string]interface{}
	err := json.Unmarshal(b, &EmailcampaignscheduleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EmailcampaignscheduleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EmailcampaignscheduleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := EmailcampaignscheduleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := EmailcampaignscheduleMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := EmailcampaignscheduleMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Intervals, ok := EmailcampaignscheduleMap["intervals"].([]interface{}); ok {
		IntervalsString, _ := json.Marshal(Intervals)
		json.Unmarshal(IntervalsString, &o.Intervals)
	}
	
	if TimeZone, ok := EmailcampaignscheduleMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if EmailCampaign, ok := EmailcampaignscheduleMap["emailCampaign"].(map[string]interface{}); ok {
		EmailCampaignString, _ := json.Marshal(EmailCampaign)
		json.Unmarshal(EmailCampaignString, &o.EmailCampaign)
	}
	
	if SelfUri, ok := EmailcampaignscheduleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Emailcampaignschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
