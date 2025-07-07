// language has the single usecase of being a namespace for the Language type and its enum values
package language

// Language represents a language available in the HAFAS API
//
// 1.2.14. Languages:
// The journey planer supports multiple languages. [...] The chosen language only influences the returned Notes in the ReST responses.
type Language = string

const (
	AR Language = "ar" // Arabic
	CA Language = "ca" // Catalan, Valencian
	DA Language = "da" // Danish
	DE Language = "de" // German
	EL Language = "el" // Greek
	EN Language = "en" // English
	ES Language = "es" // Spanish
	FI Language = "fi" // Finnish
	FR Language = "fr" // French
	HI Language = "hi" // Hindi
	HR Language = "hr" // Croatian
	HU Language = "hu" // Hungarian
	IT Language = "it" // Italian
	NL Language = "nl" // Dutch
	NO Language = "no" // Norwegian
	PL Language = "pl" // Polish
	RU Language = "ru" // Russian
	SK Language = "sk" // Slovak
	SL Language = "sl" // Slovenian
	SV Language = "sv" // Swedish
	TL Language = "tl" // Tagalog
	TR Language = "tr" // Turkish
	UR Language = "ur" // Urdu
	ZH Language = "zh" // Chinese
)
