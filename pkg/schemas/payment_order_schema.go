package schemas

type PublishCreatedPaymentOrderPayload struct {
	CheckoutId uint `json:"checkout_id"`
	PaymentId  uint `json:"payment_id"`
}

type CreateCheckoutEventPayload struct {
	UserId     uint `json:"user_id"`
	CheckoutId uint `json:"checkout_id"`
}
