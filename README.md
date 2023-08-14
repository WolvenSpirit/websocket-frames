### RFC6455 implementation

This is a websocket protocol implementation without third-party libraries.
The library is built mainly for educational purposes and to basically asimilate more in-depth knowledge of the protocol.

While it will give full control over the frame except for some primitives and helper functions for parsing and building, it is probably safe to assume the same level of control can be achieved by using the standard library websocket package or gorilla websocket and digging a bit.

The only extras you might find here are exported constants that are defined as described in the websocket RFC and that are probably unexported in more packages with higher level of abstraction.

References used:

https://www.rfc-editor.org/rfc/rfc7692.html
https://datatracker.ietf.org/doc/html/rfc6455
https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_server