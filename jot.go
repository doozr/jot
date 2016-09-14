/*Package jot contains a simple logger for developers making notes during development.

It is similar in concept and use to the `debug` log level that many loggers
provide. A way of  making notes and annotations to the code that appear at
runtime.

Getting A Jotter Instance

The `Jotter` instance can be created manually, or a global instance can be used
via the top level functions. `Jotter` wraps an instance of `Printer`, which can
be any object that implements `Print`, `Printf` and `Println` in the same way as
the `fmt` package. Coincidentally the `log.Logger` type also conforms to this
interface.

	logger := log.New(os.Stderr, "", 0)
	jotter := jot.New(logger)

By default a standard `Jotter` instance is created using a `log.Logger` instance
configured to write to Stderr and use standard log formatting. Override this by
called `SetPrinter` and pass in a custom `Printer` instance.

Enabling Jotter

The `Jotter` instance can be enabled by called `Enable()` on it, or the on the
package to enable the standard `Jotter`. It can be turned off again by calling
`Disable()`. This is useful for being able to turn on `Jotter` at runtime via
some secret API call.

	jot.Enable()
	jot.Print("This is printed")
	jot.Disable()
	jot.Print("This is not")

A useful way to enable Jotter could be to use an environment variable. This is
not enabled by default to prevent a generic way of enabling detailed output for
any program that uses Jotter, but it is easy to add.

	if os.Getenv("JOTTER_ENABLE") == "true" {
		jot.Enable()
	}

Then, to use it:

	export JOTTER_ENABLE=true
	my_program

Example

	jot.Print("Calling connectToThing", someParam)
	jot.Printf("User: %s ACL: %s", user, acl)
	result, err := connectToThing(someParam)

	if err != nil {
		log.Printf("Error connecting to thing with %s: %v", someParam, err)
	}

	if (result == "specific value of interest to developer") {
		jot.Printf("TRACER: result is %s", result)
	}

In this example the log line will be printed in case of error as is proper, but
the jot lines are only printed if the standard `Jotter` instance is enabled.
This ensures that noisy log lines that only help developers can be avoided
unless absolutely required.

*/
package jot

import (
	"log"
	"os"
)

var jotter = Jotter{
	enabled: false,
	printer: log.New(os.Stderr, "", log.LstdFlags),
}

// SetPrinter changes the Printer instance used by the standard Jotter.
func SetPrinter(printer Printer) {
	jotter.printer = printer
}

// Enable output from the standard Jotter.
func Enable() {
	jotter.Enable()
}

// Disable output from the standard Jotter.
func Disable() {
	jotter.Disable()
}

// Print via the standard Jotter.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	jotter.Print(v...)
}

// Printf via the standard Jotter.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	jotter.Printf(format, v...)
}

// Println via the standard Jotter.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	jotter.Println(v...)
}