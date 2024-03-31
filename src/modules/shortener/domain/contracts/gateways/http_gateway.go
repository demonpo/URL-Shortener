package gateways

type RedirectInput struct {
	Url string
}

type HttpGateway interface {
	redirect(input RedirectInput) error
}
