package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Patchcontentofferstylingconfiguration) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
