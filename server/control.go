package server

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type Control struct {
	l        *logrus.Logger
	start    func() error
	stop     func()
	stopChan chan interface{}
}

func (c *Control) Start() {
	if err := c.start(); err != nil {
		c.l.WithError(err).Error("Error when starting")
		go func() {
			c.stopChan <- 1
		}()
	}
}

func (c *Control) Stop() {
	c.stop()
}

// ShutdownBlock will listen for and block on term and interrupt signals, calling Control.Stop() once signalled
func (c *Control) ShutdownBlock() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGTERM)
	signal.Notify(sigChan, syscall.SIGINT)

	select {
	case _ = <-c.stopChan:
		c.l.Println("Received signal, shutting down")
	case rawSig := <-sigChan:
		sig := rawSig.String()
		c.l.WithField("signal", sig).Info("Caught signal, shutting down")
	}

	c.Stop()
}
