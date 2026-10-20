package tcp

import (
	"errors"
	"gameSrv/pkg/log"
	"time"

	"github.com/panjf2000/gnet/v2"
)

type Client struct{}

type gnetHandler struct {
	gnet.BuiltinEventEngine
	gameHandler EventHandler
	codec       ICodec
}

var gClient *gnet.Client

func ClientStart(handler EventHandler, opts ...gnet.Option) error {
	gnetHandler := &gnetHandler{
		gameHandler: handler,
		codec:       &DefaultCodec{},
	}

	client, err := gnet.NewClient(gnetHandler, opts...)
	if err != nil {
		return err
	}

	gClient = client
	return client.Start()
}

func ClientStop() {
	if gClient != nil {
		gClient.Stop()
	}
}

func ClientInited() bool {
	return gClient != nil
}

func Dial(network, address string) (Channel, error) {
	if gClient == nil {
		return nil, errors.New("tcp client not started, call ClientStart first")
	}

	conn, err := gClient.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &ChannelGnet{conn: conn}, nil
}

func (h *gnetHandler) OnBoot(eng gnet.Engine) (action gnet.Action) {
	return gnet.None
}

func (h *gnetHandler) OnShutdown(srv gnet.Engine) {}

func (h *gnetHandler) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	ctx := &ChannelGnet{conn: c, ctx: nil}
	c.SetContext(ctx)
	opened, a := h.gameHandler.OnOpened(ctx)
	return opened, gnet.Action(a)
}

func (h *gnetHandler) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	ctx, ok := c.Context().(Channel)
	if !ok {
		return gnet.Close
	}
	h.gameHandler.OnClosed(ctx, err)
	return gnet.Close
}

func (h *gnetHandler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	ctx, ok := c.Context().(Channel)
	if !ok {
		return gnet.Close
	}

	for c.InboundBuffered() > 0 {
		packet, err := h.codec.Decode(c)
		if err != nil {
			if err == ErrIncompletePacket {
				return gnet.None
			}
			log.Errorf("Decode error: %v, addr=%s", err, c.RemoteAddr())
			return gnet.Close
		}
		a := h.gameHandler.React(packet, ctx)
		if gnet.Action(a) == gnet.Close {
			return gnet.Close
		}
	}
	return gnet.None
}

func (h *gnetHandler) OnTick() (delay time.Duration, action gnet.Action) {
	tick, a := h.gameHandler.Tick()
	return tick, gnet.Action(a)
}
