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

func (u *Patchcontentofferstylingconfiguration) MarshalJSON() ([]byte, error) {
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
func (o *Patchcontentofferstylingconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
