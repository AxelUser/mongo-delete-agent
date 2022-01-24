package constants

const (
	Purchase  = "purchase"
	PageVisit = "pageVisit"
	LinkClick = "linkClick"
)

var (
	AllEventTypes = [...]string{
		Purchase, PageVisit, LinkClick,
	}
)
