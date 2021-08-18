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

func (u *Contentofferstylingconfiguration) MarshalJSON() ([]byte, error) {
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
		Position: u.Position,
		
		Offer: u.Offer,
		
		CloseButton: u.CloseButton,
		
		CtaButton: u.CtaButton,
		
		Title: u.Title,
		
		Headline: u.Headline,
		
		Body: u.Body,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contentofferstylingconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
