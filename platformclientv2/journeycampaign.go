package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeycampaign
type Journeycampaign struct { 
	// Content - Differentiate ads or links that point to the same URL (e.g. textlink).
	Content *string `json:"content,omitempty"`


	// Medium - Identify a medium such as email or cost-per-click (e.g. CPC).
	Medium *string `json:"medium,omitempty"`


	// Name - Identify a specific product promotion or strategic campaign (e.g. 320banner).
	Name *string `json:"name,omitempty"`


	// Source - Identify a search engine, newsletter name, or other source (e.g. Google).
	Source *string `json:"source,omitempty"`


	// Term - Note the keywords for this ad (e.g. running+shoes).
	Term *string `json:"term,omitempty"`


	// ClickId - The click ID (unique number that is generated when a potential customer clicks on an affiliate link).
	ClickId *string `json:"clickId,omitempty"`


	// Network - The ad network to which the click ID belongs.
	Network *string `json:"network,omitempty"`

}

func (o *Journeycampaign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeycampaign
	
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

func (o *Journeycampaign) UnmarshalJSON(b []byte) error {
	var JourneycampaignMap map[string]interface{}
	err := json.Unmarshal(b, &JourneycampaignMap)
	if err != nil {
		return err
	}
	
	if Content, ok := JourneycampaignMap["content"].(string); ok {
		o.Content = &Content
	}
	
	if Medium, ok := JourneycampaignMap["medium"].(string); ok {
		o.Medium = &Medium
	}
	
	if Name, ok := JourneycampaignMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Source, ok := JourneycampaignMap["source"].(string); ok {
		o.Source = &Source
	}
	
	if Term, ok := JourneycampaignMap["term"].(string); ok {
		o.Term = &Term
	}
	
	if ClickId, ok := JourneycampaignMap["clickId"].(string); ok {
		o.ClickId = &ClickId
	}
	
	if Network, ok := JourneycampaignMap["network"].(string); ok {
		o.Network = &Network
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeycampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
