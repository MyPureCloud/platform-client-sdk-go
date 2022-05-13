package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationmktcampaign
type Journeywebactioneventsnotificationmktcampaign struct { 
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

func (o *Journeywebactioneventsnotificationmktcampaign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebactioneventsnotificationmktcampaign
	
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

func (o *Journeywebactioneventsnotificationmktcampaign) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationmktcampaignMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationmktcampaignMap)
	if err != nil {
		return err
	}
	
	if Content, ok := JourneywebactioneventsnotificationmktcampaignMap["content"].(string); ok {
		o.Content = &Content
	}
    
	if Medium, ok := JourneywebactioneventsnotificationmktcampaignMap["medium"].(string); ok {
		o.Medium = &Medium
	}
    
	if Name, ok := JourneywebactioneventsnotificationmktcampaignMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Source, ok := JourneywebactioneventsnotificationmktcampaignMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if Term, ok := JourneywebactioneventsnotificationmktcampaignMap["term"].(string); ok {
		o.Term = &Term
	}
    
	if ClickId, ok := JourneywebactioneventsnotificationmktcampaignMap["clickId"].(string); ok {
		o.ClickId = &ClickId
	}
    
	if Network, ok := JourneywebactioneventsnotificationmktcampaignMap["network"].(string); ok {
		o.Network = &Network
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationmktcampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
