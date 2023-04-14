# Core - Control

This package is small and limited. It basically just provides the function for blocking application flow (thread blocking) until a viable system interrupt is received, in which it releases. This is fundamental in the multi-threaded style this library operates. Essentially any interrupt and the system signals of `SIGINT`, `SIGTERM`, and `SIGHUP` will release the thread. This function is called near the end of the `core.Run()` function to hold the application while all the packages are running (mainly just `web`). After being released, the cleanup and shutdown functions will begin to be called.

In addition to the blocking function. It also tracks shutdown callbacks for an event-driven style shutdown. Packages can add a call-back that will be called when shutdown is in progress using the function `control.AddShutdownCallback()`.  At this time they are called in order, so **do not rely on calling order** of the callbacks. Internally, core maintains a pre- and post- shutdown callback for the packages it provides.

Since core handles everything for you, no cleanup of core packages is required.

