package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchcontentofferstylingconfiguration
type Patchcontentofferstylingconfiguration struct { 
	// Position - Properties for customizing the positioning of the content offer.
	Position *Patchcontentpositionproperties `json:"position,omitempty"`


	// Offer - Properties for customizing the appearance of the content offer.
	Offer *Patchcontentofferstyleproperties `json:"offer,omitempty"`


	// CloseButton - Properties for customizing the appearance of the close button.
	CloseButton *Patchclosebuttonstyleproperties `json:"closeButton,omitempty"`


	// CtaButton - Properties for customizing the appearance of the CTA button.
	CtaButton *Patchctabuttonstyleproperties `json:"ctaButton,omitempty"`


	// Title - Properties for customizing the appearance of the title text.
	Title *Patchtextstyleproperties `json:"title,omitempty"`


	// Headline - Properties for customizing the appearance of the headline text.
	Headline *Patchtextstyleproperties `json:"headline,omitempty"`


	// Body - Properties for customizing the appearance of the body text.
	Body *Patchtextstyleproperties `json:"body,omitempty"`

}

func (o *Patchcontentofferstylingconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchcontentofferstylingconfiguration
	
	return json.Marshal(&struct { 
		Position *Patchcontentpositionproperties `json:"position,omitempty"`
		
		Offer *Patchcontentofferstyleproperties `json:"offer,omitempty"`
		
		CloseButton *Patchclosebuttonstyleproperties `json:"closeButton,omitempty"`
		
		CtaButton *Patchctabuttonstyleproperties `json:"ctaButton,omitempty"`
		
		Title *Patchtextstyleproperties `json:"title,omitempty"`
		
		Headline *Patchtextstyleproperties `json:"headline,omitempty"`
		
		Body *Patchtextstyleproperties `json:"body,omitempty"`
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

func (o *Patchcontentofferstylingconfiguration) UnmarshalJSON(b []byte) error {
	var PatchcontentofferstylingconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &PatchcontentofferstylingconfigurationMap)
	if err != nil {
		return err
	}
	
	if Position, ok := PatchcontentofferstylingconfigurationMap["position"].(map[string]interface{}); ok {
		PositionString, _ := json.Marshal(Position)
		json.Unmarshal(PositionString, &o.Position)
	}
	
	if Offer, ok := PatchcontentofferstylingconfigurationMap["offer"].(map[string]interface{}); ok {
		OfferString, _ := json.Marshal(Offer)
		json.Unmarshal(OfferString, &o.Offer)
	}
	
	if CloseButton, ok := PatchcontentofferstylingconfigurationMap["closeButton"].(map[string]interface{}); ok {
		CloseButtonString, _ := json.Marshal(CloseButton)
		json.Unmarshal(CloseButtonString, &o.CloseButton)
	}
	
	if CtaButton, ok := PatchcontentofferstylingconfigurationMap["ctaButton"].(map[string]interface{}); ok {
		CtaButtonString, _ := json.Marshal(CtaButton)
		json.Unmarshal(CtaButtonString, &o.CtaButton)
	}
	
	if Title, ok := PatchcontentofferstylingconfigurationMap["title"].(map[string]interface{}); ok {
		TitleString, _ := json.Marshal(Title)
		json.Unmarshal(TitleString, &o.Title)
	}
	
	if Headline, ok := PatchcontentofferstylingconfigurationMap["headline"].(map[string]interface{}); ok {
		HeadlineString, _ := json.Marshal(Headline)
		json.Unmarshal(HeadlineString, &o.Headline)
	}
	
	if Body, ok := PatchcontentofferstylingconfigurationMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchcontentofferstylingconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
