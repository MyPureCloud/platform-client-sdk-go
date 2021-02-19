package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Contentofferstylingconfiguration) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
