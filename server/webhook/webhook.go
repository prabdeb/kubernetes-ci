package webhook

// Enables webhook based triggers, instead of polling
// disable polling when enabled for a repository

type (
	// WebHook type
	WebHook struct {
		Repository *Repository
		Builder    *Builder
		Pipeline   string
	}
	// Repository type
	Repository struct {
		Owner    string
		Name     string
		Refs     string
		Compare  string
		FullName string
		CloneURL string
	}
	// Builder type
	Builder struct {
		Name   string
		Email  string
		Avatar string
	}
	// SCM type
	SCM struct {
		UserName string
		Password string
		Type     string
	}
)

// New func
func New() *WebHook {
	return &WebHook{}
}
