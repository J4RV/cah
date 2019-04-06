import React, { Component } from "react"

import Hand from "../gamestate/Hand"
import PlayersInfo from "../gamestate/PlayersInfo"
import Table from "../gamestate/Table"
import { connect } from "react-redux"
import { gameStateWSocketAbsUrl } from "../restUrls"
import pushError from "../actions/pushError"

const MAX_RETRIES_ON_SOCKET_CLOSED = 5

class Game extends Component {
  render() {
    if (this.state == null) return null
    return (
      <div className="cah-game">
        <PlayersInfo state={this.state} />
        <Table state={this.state} />
        <Hand gamestate={this.state} />
      </div>
    )
  }

  componentWillMount() {
    const stateID = this.props.stateID
    this.startWebsocket(stateID, MAX_RETRIES_ON_SOCKET_CLOSED)
  }

  startWebsocket = (stateID, retries) => {
    const sock = new WebSocket(gameStateWSocketAbsUrl(stateID))

    sock.onmessage = e => {
      this.setState(JSON.parse(e.data))
      retries = MAX_RETRIES_ON_SOCKET_CLOSED
    }

    sock.onclose = () => {
      retries--
      if (retries <= 0) {
        this.props.pushError("Could not reconnect to server.")
        return
      }
      console.error("Server Connection was lost, reconnecting...")
      setTimeout(() => this.startWebsocket(stateID, retries), 2000 / retries)
    }

    sock.onerror = err => {
      this.props.pushError("Server Connection error: " + err.message)
      sock.close()
    }
  }
}

export default connect(
  null,
  { pushError }
)(Game)
