package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentofferstylingconfiguration
type Contentofferstylingconfiguration struct { 
	// Position - Properties for customizing the positioning of the content offer.
	Position *Contentpositionproperties `json:"position,omitempty"`


	// Offer - Properties for customizing the appearance of the content offer.
	Offer *Contentofferstyleproperties `json:"offer,omitempty"`


	// CloseButton - Properties for customizing the appearance of the close button.
	CloseButton *Closebuttonstyleproperties `json:"closeButton,omitempty"`


	// CtaButton - Properties for customizing the appearance of the CTA button.
	CtaButton *Ctabuttonstyleproperties `json:"ctaButton,omitempty"`


	// Title - Properties for customizing the appearance of the title text.
	Title *Textstyleproperties `json:"title,omitempty"`


	// Headline - Properties for customizing the appearance of the headline text.
	Headline *Textstyleproperties `json:"headline,omitempty"`


	// Body - Properties for customizing the appearance of the body text.
	Body *Textstyleproperties `json:"body,omitempty"`

}

func (o *Contentofferstylingconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentofferstylingconfiguration
	
	return json.Marshal(&struct { 
		Position *Contentpositionproperties `json:"position,omitempty"`
		
		Offer *Contentofferstyleproperties `json:"offer,omitempty"`
		
		CloseButton *Closebuttonstyleproperties `json:"closeButton,omitempty"`
		
		CtaButton *Ctabuttonstyleproperties `json:"ctaButton,omitempty"`
		
		Title *Textstyleproperties `json:"title,omitempty"`
		
		Headline *Textstyleproperties `json:"headline,omitempty"`
		
		Body *Textstyleproperties `json:"body,omitempty"`
		*Alias
	}{ 
		Position: o.Position,
		
		Offer: o.Offer,
		
		CloseButton: o.CloseButton,
		
		CtaButton: o.CtaButton,
		
		Title: o.Title,
		
		Headline: o.Headline,
		
		Body: o.Body,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentofferstylingconfiguration) UnmarshalJSON(b []byte) error {
	var ContentofferstylingconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &ContentofferstylingconfigurationMap)
	if err != nil {
		return err
	}
	
	if Position, ok := ContentofferstylingconfigurationMap["position"].(map[string]interface{}); ok {
		PositionString, _ := json.Marshal(Position)
		json.Unmarshal(PositionString, &o.Position)
	}
	
	if Offer, ok := ContentofferstylingconfigurationMap["offer"].(map[string]interface{}); ok {
		OfferString, _ := json.Marshal(Offer)
		json.Unmarshal(OfferString, &o.Offer)
	}
	
	if CloseButton, ok := ContentofferstylingconfigurationMap["closeButton"].(map[string]interface{}); ok {
		CloseButtonString, _ := json.Marshal(CloseButton)
		json.Unmarshal(CloseButtonString, &o.CloseButton)
	}
	
	if CtaButton, ok := ContentofferstylingconfigurationMap["ctaButton"].(map[string]interface{}); ok {
		CtaButtonString, _ := json.Marshal(CtaButton)
		json.Unmarshal(CtaButtonString, &o.CtaButton)
	}
	
	if Title, ok := ContentofferstylingconfigurationMap["title"].(map[string]interface{}); ok {
		TitleString, _ := json.Marshal(Title)
		json.Unmarshal(TitleString, &o.Title)
	}
	
	if Headline, ok := ContentofferstylingconfigurationMap["headline"].(map[string]interface{}); ok {
		HeadlineString, _ := json.Marshal(Headline)
		json.Unmarshal(HeadlineString, &o.Headline)
	}
	
	if Body, ok := ContentofferstylingconfigurationMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentofferstylingconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
