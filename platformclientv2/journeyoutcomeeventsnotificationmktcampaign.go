package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationmktcampaign
type Journeyoutcomeeventsnotificationmktcampaign struct { 
	// Content
	Content *string `json:"content,omitempty"`


	// Medium
	Medium *string `json:"medium,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Source
	Source *string `json:"source,omitempty"`


	// Term
	Term *string `json:"term,omitempty"`


	// ClickId
	ClickId *string `json:"clickId,omitempty"`


	// Network
	Network *string `json:"network,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationmktcampaign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationmktcampaign
	
	return json.Marshal(&struct { 
		Content *string `json:"content,omitempty"`
		
		Medium *string `json:"medium,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		Term *string `json:"term,omitempty"`
		
		ClickId *string `json:"clickId,omitempty"`
		
		Network *string `json:"network,omitempty"`
		*Alias
	}{ 
		Content: o.Content,
		
		Medium: o.Medium,
		
		Name: o.Name,
		
		Source: o.Source,
		
		Term: o.Term,
		
		ClickId: o.ClickId,
		
		Network: o.Network,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationmktcampaign) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationmktcampaignMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationmktcampaignMap)
	if err != nil {
		return err
	}
	
	if Content, ok := JourneyoutcomeeventsnotificationmktcampaignMap["content"].(string); ok {
		o.Content = &Content
	}
    
	if Medium, ok := JourneyoutcomeeventsnotificationmktcampaignMap["medium"].(string); ok {
		o.Medium = &Medium
	}
    
	if Name, ok := JourneyoutcomeeventsnotificationmktcampaignMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Source, ok := JourneyoutcomeeventsnotificationmktcampaignMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if Term, ok := JourneyoutcomeeventsnotificationmktcampaignMap["term"].(string); ok {
		o.Term = &Term
	}
    
	if ClickId, ok := JourneyoutcomeeventsnotificationmktcampaignMap["clickId"].(string); ok {
		o.ClickId = &ClickId
	}
    
	if Network, ok := JourneyoutcomeeventsnotificationmktcampaignMap["network"].(string); ok {
		o.Network = &Network
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationmktcampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
