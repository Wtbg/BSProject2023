package mqtt

import (
	"bytes"
	"encoding/json"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/model"
)

func Init() {
	server := mqtt.New(&mqtt.Options{})
	_ = server.AddHook(new(auth.AllowHook), nil)
	tcp := listeners.NewTCP("t1", ":1883", nil)
	err := server.AddListener(tcp)
	if err != nil {
		logrus.Fatal(err)
	}
	err = server.AddHook(new(ExampleHook), &ExampleHookOptions{
		Server: server,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	go func() {
		err := server.Serve()
		if err != nil {
			logrus.Fatal(err)
		}
	}()
}

// Options contains configuration settings for the hook.
type ExampleHookOptions struct {
	Server *mqtt.Server
}

type ExampleHook struct {
	mqtt.HookBase
	config *ExampleHookOptions
}

func (h *ExampleHook) ID() string {
	return "events-example"
}

func (h *ExampleHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnect,
		mqtt.OnDisconnect,
		mqtt.OnPublished,
	}, []byte{b})
}

func (h *ExampleHook) Init(config any) error {
	if _, ok := config.(*ExampleHookOptions); !ok && config != nil {
		return mqtt.ErrInvalidConfigType
	}
	h.config = config.(*ExampleHookOptions)
	if h.config.Server == nil {
		return mqtt.ErrInvalidConfigType
	}
	return nil
}

func (h *ExampleHook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	logrus.Info("client connected", "client", cl.ID)
	//try insert the client to database, fail if the client already exist
	client := model.Device{
		DeviceID: cl.ID,
	}
	err := model.DB.FirstOrCreate(&client).Error
	if err != nil {
		logrus.Info("client already exist", "client", cl.ID)
	} else {
		logrus.Info("client inserted", "client", cl.ID)
	}
	//update the client status to connected
	err = model.DB.Model(&client).Update("is_connected", true).Error
	if err != nil {
		logrus.Error("error update client status", "error", err)
	} else {
		logrus.Info("client connecting status updated to connected", "client", cl.ID)
	}
	return nil
}

func (h *ExampleHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	if err != nil {
		logrus.Info("client disconnected", "client", cl.ID, "expire", expire, "error", err)
	} else {
		logrus.Info("client disconnected", "client", cl.ID, "expire", expire)
	}
	//update the client status to disconnected
	client := model.Device{
		DeviceID: cl.ID,
	}
	err = model.DB.Model(&client).Update("is_connected", false).Error
	if err != nil {
		logrus.Error("error update client status", "error", err)
	} else {
		logrus.Info("client connecting status updated to disconnected", "client", cl.ID)
	}
}

func (h *ExampleHook) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	//analysis the payload to a IOTMessage
	msg := model.DeviceMessage{}
	err := json.Unmarshal(pk.Payload, &msg)
	if err != nil {
		logrus.Error("error unmarshal payload", "error", err)
		return
	} else {
		logrus.Info("client published", "client", cl.ID, "topic", pk.TopicName, "payload", msg)
	}
	//insert the message to database
	if model.DB.Create(&msg).Error != nil {
		logrus.Error("error insert message", "error", err)
	} else {
		logrus.Info("message inserted", "client", cl.ID, "topic", pk.TopicName, "payload", msg)
	}
	if model.DB.Model(&model.Device{}).Where("device_id = ?", cl.ID).Update("alert", msg.Alert).Error != nil {
		logrus.Error("error update client warning status", "error", err)
	} else {
		logrus.Info("client warning status updated", "client", cl.ID, "topic", pk.TopicName, "payload", msg)
	}
}
