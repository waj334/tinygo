package net

import "context"

func lookupProtocol(ctx context.Context, name string) (int, error) {
	panic("Not implemented")
}

func (r *Resolver) lookupHost(ctx context.Context, host string) (addrs []string, err error) {
	panic("Not implemented")
}

func (r *Resolver) lookupIP(ctx context.Context, network, host string) (addrs []IPAddr, err error) {
	panic("Not implemented")
}

func (r *Resolver) lookupPort(ctx context.Context, network, service string) (int, error) {
	panic("Not implemented")
}

func (r *Resolver) lookupCNAME(ctx context.Context, name string) (string, error) {
	panic("Not implemented")
}

func (r *Resolver) lookupSRV(ctx context.Context, service, proto, name string) (string, []*SRV, error) {
	panic("Not implemented")
}

func (r *Resolver) lookupMX(ctx context.Context, name string) ([]*MX, error) {
	panic("Not implemented")
}

func (r *Resolver) lookupNS(ctx context.Context, name string) ([]*NS, error) {
	panic("Not implemented")
}

func (r *Resolver) lookupTXT(ctx context.Context, name string) ([]string, error) {
	panic("Not implemented")
}

func (r *Resolver) lookupAddr(ctx context.Context, addr string) ([]string, error) {
	panic("Not implemented")
}