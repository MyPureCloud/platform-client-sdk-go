package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignsequence
type Campaignsequence struct { 
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


	// Campaigns - The ordered list of Campaigns that this CampaignSequence will run.
	Campaigns *[]Domainentityref `json:"campaigns,omitempty"`


	// CurrentCampaign - A zero-based index indicating which Campaign this CampaignSequence is currently on.
	CurrentCampaign *int `json:"currentCampaign,omitempty"`


	// Status - The current status of the CampaignSequence. A CampaignSequence can be turned 'on' or 'off'.
	Status *string `json:"status,omitempty"`


	// StopMessage - A message indicating if and why a CampaignSequence has stopped unexpectedly.
	StopMessage *string `json:"stopMessage,omitempty"`


	// Repeat - Indicates if a sequence should repeat from the beginning after the last campaign completes. Default is false.
	Repeat *bool `json:"repeat,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Campaignsequence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignsequence
	
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
		
		Campaigns *[]Domainentityref `json:"campaigns,omitempty"`
		
		CurrentCampaign *int `json:"currentCampaign,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		StopMessage *string `json:"stopMessage,omitempty"`
		
		Repeat *bool `json:"repeat,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		Campaigns: o.Campaigns,
		
		CurrentCampaign: o.CurrentCampaign,
		
		Status: o.Status,
		
		StopMessage: o.StopMessage,
		
		Repeat: o.Repeat,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaignsequence) UnmarshalJSON(b []byte) error {
	var CampaignsequenceMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignsequenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CampaignsequenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := CampaignsequenceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := CampaignsequenceMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := CampaignsequenceMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := CampaignsequenceMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Campaigns, ok := CampaignsequenceMap["campaigns"].([]interface{}); ok {
		CampaignsString, _ := json.Marshal(Campaigns)
		json.Unmarshal(CampaignsString, &o.Campaigns)
	}
	
	if CurrentCampaign, ok := CampaignsequenceMap["currentCampaign"].(float64); ok {
		CurrentCampaignInt := int(CurrentCampaign)
		o.CurrentCampaign = &CurrentCampaignInt
	}
	
	if Status, ok := CampaignsequenceMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if StopMessage, ok := CampaignsequenceMap["stopMessage"].(string); ok {
		o.StopMessage = &StopMessage
	}
	
	if Repeat, ok := CampaignsequenceMap["repeat"].(bool); ok {
		o.Repeat = &Repeat
	}
	
	if SelfUri, ok := CampaignsequenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignsequence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
