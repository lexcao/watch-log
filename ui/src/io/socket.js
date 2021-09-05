import { readable} from 'svelte/store'

class Socket {
    constructor(port) {
        this.port = port
        this.connect()
    }

    sendMessage(body) {
        if (!this.connected) {
            console.log('Can\'t send message for no connection', body)
            return
        }

        console.log('Sending...', body)
        this.socket.send(body)
    }

    connect() {
        this.socket = new WebSocket(`ws://localhost:${this.port}/ws`)

        this.connected = readable(false, set => {
            this.socket.onopen = () => {
                console.log('OPEN connection...')
                set(true)
            }
        })

        this.messages = readable([], set => {
            let data = []

            this.socket.onmessage = (event) => {
                console.log('Receiving event', event)

                data = [...data, JSON.parse(event.data)]

                set(data)
            }
        })
    }
}

export default Socket
