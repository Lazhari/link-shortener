package shortener

// RedirectRepository represents a redirect repository interface
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
