# goreactive - reactive template variables with WebSockets

`goreactive` allows you to build web applications such as dashboards that update in real-time without writing any JavaScript. Use standard Go templates, embed them with `ReactiveVar` objects and the library takes care of the rest. Updates are pushed to clients over Websockets.

See the [example](./example/) folder for a runnable example.

## Major missing features

- Server-side
  - Authentication
  - Review 'nicer'/proper ways to handle closing the messageBroker channel etc. when the application is stopped
  - Some way to customise the printing format of variables, or allow further customisation of what is sent to the client
  - Tests + load tests
- Client-side
  - Inform client on error/disconnect
  - Automatic attempts to redirect to client

