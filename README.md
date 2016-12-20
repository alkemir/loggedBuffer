# loggedBuffer
Golang buffers which keep track of their operations so you can debug easily.
Note that logging is not safe for concurrent access, as the original package.

Just plug the LoggedBuffer where you would otherwise use a bytes.Buffer (even
  its zero-value) and use LoggedBuffer.Logs() to retrieve the logs when you
  are done.
