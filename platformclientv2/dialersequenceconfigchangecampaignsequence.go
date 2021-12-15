package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialersequenceconfigchangecampaignsequence
type Dialersequenceconfigchangecampaignsequence struct { 
	// Campaigns - the ordered list of campaign identifiers
	Campaigns *[]Dialersequenceconfigchangeurireference `json:"campaigns,omitempty"`


	// CurrentCampaign - the zero-based index of the current campaign in the campaigns list
	CurrentCampaign *int `json:"currentCampaign,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// StopMessage - if a sequence has unexpectedly stopped, this message provides the reason
	StopMessage *string `json:"stopMessage,omitempty"`


	// Repeat - indicates if a sequence is to repeat from the beginning after the last campaign completes; default is false
	Repeat *bool `json:"repeat,omitempty"`


	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

}

func (o *Dialersequenceconfigchangecampaignsequence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialersequenceconfigchangecampaignsequence
	
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
		Campaigns *[]Dialersequenceconfigchangeurireference `json:"campaigns,omitempty"`
		
		CurrentCampaign *int `json:"currentCampaign,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		StopMessage *string `json:"stopMessage,omitempty"`
		
		Repeat *bool `json:"repeat,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Campaigns: o.Campaigns,
		
		CurrentCampaign: o.CurrentCampaign,
		
		Status: o.Status,
		
		StopMessage: o.StopMessage,
		
		Repeat: o.Repeat,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialersequenceconfigchangecampaignsequence) UnmarshalJSON(b []byte) error {
	var DialersequenceconfigchangecampaignsequenceMap map[string]interface{}
	err := json.Unmarshal(b, &DialersequenceconfigchangecampaignsequenceMap)
	if err != nil {
		return err
	}
	
	if Campaigns, ok := DialersequenceconfigchangecampaignsequenceMap["campaigns"].([]interface{}); ok {
		CampaignsString, _ := json.Marshal(Campaigns)
		json.Unmarshal(CampaignsString, &o.Campaigns)
	}
	
	if CurrentCampaign, ok := DialersequenceconfigchangecampaignsequenceMap["currentCampaign"].(float64); ok {
		CurrentCampaignInt := int(CurrentCampaign)
		o.CurrentCampaign = &CurrentCampaignInt
	}
	
	if Status, ok := DialersequenceconfigchangecampaignsequenceMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if StopMessage, ok := DialersequenceconfigchangecampaignsequenceMap["stopMessage"].(string); ok {
		o.StopMessage = &StopMessage
	}
	
	if Repeat, ok := DialersequenceconfigchangecampaignsequenceMap["repeat"].(bool); ok {
		o.Repeat = &Repeat
	}
	
	if Id, ok := DialersequenceconfigchangecampaignsequenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialersequenceconfigchangecampaignsequenceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := DialersequenceconfigchangecampaignsequenceMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialersequenceconfigchangecampaignsequenceMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialersequenceconfigchangecampaignsequenceMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialersequenceconfigchangecampaignsequence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
