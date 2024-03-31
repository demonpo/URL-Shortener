package gateways

type RedirectInput struct {
	Url string
}

type GinHttpGateway struct {
}

func (g *GinHttpGateway) redirect(input RedirectInput) error {
	return nil
}
