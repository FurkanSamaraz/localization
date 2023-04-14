package control

import (
	"errors"
	"sync"
)

//ShutdownFunc callback function to be called when the service is shutting down
type ShutdownFunc func()

var (
	funcLock   = &sync.Mutex{}
	funcs      = make(map[string]ShutdownFunc, 0)
	preExitCB  ShutdownFunc
	postExitCB ShutdownFunc
)

//SetPreExitCallback sets a callback that can be executed during termination before clean up callbacks are issued
func SetPreExitCallback(cb ShutdownFunc) {
	preExitCB = cb
}

//SetPostExitCallback sets a callback that can be executed during termination after clean up callbacks are issued
func SetPostExitCallback(cb ShutdownFunc) {
	postExitCB = cb
}

//AddShutdownCallback adds a callback function for the specified module to the list.
//These callbacks are executed when the service shutsdown
func AddShutdownCallback(mod string, cb ShutdownFunc) error {
	funcLock.Lock()
	defer funcLock.Unlock()

	_, exists := funcs[mod]
	if exists {
		return errors.New("attempted to add shutdown callback via control.AddShutdownCallback with an existing key " + mod)
	}

	funcs[mod] = cb

	return nil
}

//RemoveShutdownCallback removes a previously registed shutdown callback provided it's module key
func RemoveShutdownCallback(mod string) error {
	funcLock.Lock()
	defer funcLock.Unlock()

	_, exists := funcs[mod]
	if exists {
		delete(funcs, mod)
		return nil
	}

	return errors.New("attempted to remove a shutdown callback via control.RemoveShutdownCallback with a key '" + mod + "' that did not exist")
}
