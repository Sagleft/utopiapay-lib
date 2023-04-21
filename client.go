package utopiapaylib

type Client interface {
}

type defaultClient struct{}

func NewClient() Client {
	return &defaultClient{}
}
