package gogram

import "github.com/lepeico/gogram/pkg/telegram"

type eventHandler func(*Context)

// Gogram is the set of fields for bot to work
type Gogram struct {
	Telegram telegram.Telegram
	Info     telegram.User
	Update   chan telegram.Update

	stop    chan struct{}
	context Context
	events  map[string]eventHandler
}

// New is constructor for Gogram
func New(token string, opts ...interface{}) (g Gogram, err error) {
	g.Telegram = *telegram.New(token)
	g.events = make(map[string]eventHandler)
	g.Update = make(chan telegram.Update, 100)
	if me, err := g.Telegram.GetMe(); err == nil {
		g.Info = me
	}
	g.context = Context{Telegram: &g.Telegram, Bot: &g.Info}
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
		gg.context.Update = &upd
		gg.events["text"](&gg.context)
	}
}

// Stop shuts down long-poller
func (gg *Gogram) Stop() {
	close(gg.stop)
}

func (gg *Gogram) On(event string, cb eventHandler) *Gogram {
	gg.events[event] = cb
	return gg
}
