package custom

import (
	"context"
	"time"
)

type Context struct {
	pai     context.Context
	valores map[string]any
}

func NewContext(pai context.Context, valores map[string]any) *Context {
	return &Context{
		pai:     pai,
		valores: valores,
	}
}

// Meu custom context não tem deadline definido (ignora o do ctx pai)
func (Context) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

// Ignora o canal done do ctx pai
func (Context) Done() <-chan struct{} {
	return nil
}

func (Context) Err() error {
	return nil
}

func (c Context) Value(chave string) any {
	if v, ok := c.valores[chave]; ok {
		return v
	}
	return c.pai.Value(chave)
}
