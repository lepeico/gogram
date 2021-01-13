package gogram

import "reflect"

type middleware interface {
	Call(ctx *Context) *Context
}

type Middleware func(*Context) *Context

func (m Middleware) Call(ctx *Context) *Context {
	return m(ctx)
}

func fromFn(fn func(*Context)) Middleware {
	return func(ctx *Context) *Context {
		fn(ctx)
		return ctx
	}
}

func (gg *Gogram) Use(mws ...Middleware) {
	for _, fn := range mws {
		gg.middlewares = append(gg.middlewares, fn)
	}
}

func (gg *Gogram) On(event string, fns ...func(*Context)) {
	for _, fn := range fns {
		filtered := func(ctx *Context) {
			c := reflect.ValueOf(ctx)
			if reflect.Indirect(c).FieldByName(event).IsNil() {
				fn(ctx)
			}
		}
		gg.middlewares = append(gg.middlewares, fromFn(filtered))
	}
}

func (gg *Gogram) Command(command string, fns ...func(*Context)) {
	for _, fn := range fns {
		filtered := func(ctx *Context) {
			e := (*ctx.Message.Entities)[0]
			if e.Type == "bot_command" && ctx.Message.Text[1:e.Length] == command {
				fn(ctx)
			}
		}
		gg.middlewares = append(gg.middlewares, fromFn(filtered))
	}
}

func (gg *Gogram) Start(fns ...func(*Context)) {
	gg.Command("start", fns...)
}

func (gg *Gogram) Help(fns ...func(*Context)) {
	gg.Command("help", fns...)
}


func (gg *Gogram) intercept(ctx *Context) *Context {
	for _, m := range gg.middlewares {
		ctx = m.Call(ctx)
	}
	return ctx
}
