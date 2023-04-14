//Package control facilitates the background process and thread waiting for program flow.
//Essentially allows the thread to wait for other listeners to continue working until a termination signal is given to the program
package control

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/maple-tech/core/log"
)

var signalChan chan os.Signal

//BlockUntilSignal locks the thread until a system interupt signal is recieved.
//Waits for os.Interrupt, SIGINT, SIGTERM, or SIGHUP
func BlockUntilSignal() error {
	if signalChan != nil {
		return errors.New("control.BlockUntilSignal() was called, but the signal channel is already initialized, possible duplicate call?")
	}

	signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	//Block thread
	<-signalChan

	//Notify we are cleaning up
	log.Info("control received termination signal, cleaning up")

	funcLock.Lock()
	defer funcLock.Unlock()

	//Fire off we are starting cleanup
	if preExitCB != nil {
		preExitCB()
	}

	//Fire off all the module cleanup callbacks
	for mod, cb := range funcs {
		log.Debugf("calling shutdown function for module '%s'", mod)
		cb()
	}

	//Fire off we are ending cleanup
	if postExitCB != nil {
		postExitCB()
	}

	log.Info("shutdown completed")
	return nil
}
