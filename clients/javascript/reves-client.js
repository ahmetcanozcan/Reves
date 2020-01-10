var GetRevesClient = (() => {
    /**
     * @name GetRevesClient
     * @param url {string} : base url of the reves server
     */
    var constructor = (url = "localhost:8080") => {
        var _url = `${url}/ws`
        console.log("Connecting to ", _url)
        var socket = new WebSocket(_url)
        /**
         * A map that stores handlers
         * @type {Map[string]function}
         */
        var events = {}

        var client = {}

        /**
         * @name client.on
         * @param name @type {string} : Name of the event
         * @param handler  @type {function} : Function hat handle data coming from the server
         */
        client.on = (name, handler) => {
            events[name] = handler
        }

        /**
         * @name client.emit
         * @param name @type {string} : Name of the event
         * @param payload @type {map[string]string} : Payload, as the function said
         */
        client.emit = (name, payload) => {
            var p = parsePayload(payload)
            socket.send(`${name};${p}`)
        }


        socket.onmessage = function (e) {
            console.log(e.data)
        }

        socket.onopen = function (e) {
            client.emit("Init")
            events["Init"]()
        }


        return client
    }


    /**
     * parses paylaod to string.
     * 
     * @param payload @type {map[string]string}
     * @returns {string} : 
     */
    var parsePayload = (payload) => {
        var result = ""
        if (!payload) {
            for (let key in payload) {
                result += `${key}:${payload[key]};`
            }
        }
        return result + "\n"
    }


    return constructor

})()