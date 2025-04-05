package enums

type PaymentStatus string

const (
	PENDING   PaymentStatus = "PENDING"
	COMPLETED PaymentStatus = "COMPLETED"
)

func (p PaymentStatus) ToString() string {
	switch p {
	case PENDING:
		return "PENDING"
	case COMPLETED:
		return "COMPLETED"
	default:
		return ""
	}
}
