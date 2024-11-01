package goreactive

// This constant provides the script block for HTML pages to connect to desired websocket.
const (
	WebsocketJavascriptBlock = `
      <script>
    const socket = new WebSocket("ws://"+location.host+"/reactiveVarsWebsocket");
        socket.onopen = function (event) {
            console.log('Connected to reactiveVarsWebsocket');
        };
        socket.onmessage = function (event) {
            jsonEvent = JSON.parse(event.data)
            const target = document.getElementById(jsonEvent.key);
            console.log(jsonEvent)
            target.innerHTML = jsonEvent.value;
        };
        socket.onclose = function (event) {
            console.log('Disconnected to reactiveVarsWebsocket');
        };
    </script>

  `
)
