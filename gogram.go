package gogram

import t "github.com/lepeico/gogram/telegram"

// Gogram is the set of fields for bot to work
type Gogram struct {
	Telegram t.Telegram
	Info     t.User
	Update   chan t.Update

	stop        chan struct{}
	context     Context
	middlewares []middleware
}

// New is constructor for Gogram
func New(token string, opts ...interface{}) (g Gogram, err error) {
	g.Telegram = *t.New(token)
	if me, err := g.Telegram.GetMe(); err == nil {
		g.Info = me
	}
	g.Update = make(chan t.Update, 100)

	g.context = Context{Telegram: &g.Telegram, Bot: &g.Info, State: make(map[string]interface{})}
	return
}

// Context is getter for private field Context
func (gg *Gogram) Context() Context {
	return gg.context
}

func (gg *Gogram) poll() {
	var LastUpdateID int
	go func() {
		for {
			select {
			case <-gg.stop:
				return
			default:
			}

			updates, err := gg.Telegram.GetUpdates(LastUpdateID+1, 100, 1, []string{})
			if err != nil {
				continue
			}

			for _, update := range updates {
				LastUpdateID = update.UpdateID
				gg.Update <- update
			}
		}
	}()
}

// Launch starts the process of long-polling or Webhook
func (gg *Gogram) Launch() {
	go gg.poll()

	for upd := range gg.Update {
		ctx := gg.context
		ctx.Update = &upd
		gg.intercept(&ctx)
	}
}

// Stop shuts down long-poller
func (gg *Gogram) Stop() {
	close(gg.stop)
}
