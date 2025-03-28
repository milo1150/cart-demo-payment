package enums

type PaymentStatus string

const (
	PENDING PaymentStatus = "PENDING"
	DONE    PaymentStatus = "DONE"
)

func (p PaymentStatus) ToString() string {
	switch p {
	case PENDING:
		return "PENDING"
	case DONE:
		return "DONE"
	default:
		return ""
	}
}
