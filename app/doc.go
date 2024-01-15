// SPDX-License-Identifier: Unlicense OR MIT

/*
Package app provides a platform-independent interface to operating system
functionality for running graphical user interfaces.

See https://github.com/Seikaijyu/gio for instructions to set up and run Gio programs.

# Windows

Create a new Window by calling NewWindow. On mobile platforms or when Gio
is embedded in another project, NewWindow merely connects with a previously
created window.

A Window is run by calling NextEvent in a loop. The most important event is
FrameEvent that prompts an update of the window contents.

For example:

	import "github.com/Seikaijyu/gio/unit"

	w := app.NewWindow()
	for {
		e := w.NextEvent()
		if e, ok := e.(system.FrameEvent); ok {
			ops.Reset()
			// Add operations to ops.
			...
			// Completely replace the window contents and state.
			e.Frame(ops)
		}
	}

A program must keep receiving events from the event channel until
DestroyEvent is received.

# Main

The Main function must be called from a program's main function, to hand over
control of the main thread to operating systems that need it.

Because Main is also blocking on some platforms, the event loop of a Window must run in a goroutine.

For example, to display a blank but otherwise functional window:

	package main

	import "github.com/Seikaijyu/gio/app"

	func main() {
		go func() {
			w := app.NewWindow()
			for {
				w.NextEvent()
			}
		}()
		app.Main()
	}

# Event queue

A FrameEvent's Queue method returns an event.Queue implementation that distributes
incoming events to the event handlers declared in the last frame.
See the github.com/Seikaijyu/gio/io/event package for more information about event handlers.

# Permissions

The packages under github.com/Seikaijyu/gio/app/permission should be imported
by a Gio program or by one of its dependencies to indicate that specific
operating-system permissions are required.  Please see documentation for
package github.com/Seikaijyu/gio/app/permission for more information.
*/
package app
