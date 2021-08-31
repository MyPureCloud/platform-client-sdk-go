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


	// Campaigns
	Campaigns *[]Dialersequenceconfigchangeurireference `json:"campaigns,omitempty"`


	// CurrentCampaign
	CurrentCampaign *int `json:"currentCampaign,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// StopMessage
	StopMessage *string `json:"stopMessage,omitempty"`


	// Repeat
	Repeat *bool `json:"repeat,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

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
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Campaigns *[]Dialersequenceconfigchangeurireference `json:"campaigns,omitempty"`
		
		CurrentCampaign *int `json:"currentCampaign,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		StopMessage *string `json:"stopMessage,omitempty"`
		
		Repeat *bool `json:"repeat,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
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
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialersequenceconfigchangecampaignsequence) UnmarshalJSON(b []byte) error {
	var DialersequenceconfigchangecampaignsequenceMap map[string]interface{}
	err := json.Unmarshal(b, &DialersequenceconfigchangecampaignsequenceMap)
	if err != nil {
		return err
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
	
	if AdditionalProperties, ok := DialersequenceconfigchangecampaignsequenceMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialersequenceconfigchangecampaignsequence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
