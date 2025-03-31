package schemas

type CreateCheckoutEventPayload struct {
	UserId     uint `json:"user_id"`
	CheckoutId uint `json:"checkout_id"`
}
