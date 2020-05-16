import React, { Component } from "react"

import ExtraGameInfo from "../gamestate/ExtraGameInfo"
import Hand from "../gamestate/Hand"
import PlayersInfo from "../gamestate/PlayersInfo"
import Table from "../gamestate/Table"
import Typography from "@material-ui/core/Typography"
import { connect } from "react-redux"
import { gameStateWSocketAbsUrl } from "../restUrls"
import pushError from "../actions/pushError"

const MAX_RETRIES_ON_SOCKET_CLOSED = 5

class Game extends Component {
  render() {
    if (this.state == null) return null

    if (this.state.phase === "Finished") {
      let winner = this.state.players[0]
      for (let i in this.state.players) {
        let player = this.state.players[i]
        console.log(player)
        if (player.points.length > winner.points.length) {
          winner = player
        }
      }
      return (
        <div className="cah-game">
          <PlayersInfo state={this.state} />
          <Typography align="center">
            <h1>Game finished!</h1>
            <h2>Winner: {winner.name}</h2>
            <h3>Black cards earned:</h3>
            {winner.points.map((p) => (
              <p>{p.text}</p>
            ))}
          </Typography>
        </div>
      )
    }

    return (
      <div className="cah-game">
        <Typography align="center">
          <h2>{this.state.phase}</h2>
        </Typography>
        <PlayersInfo state={this.state} />
        <Table state={this.state} />
        <Hand gamestate={this.state} />
        <ExtraGameInfo state={this.state} />
      </div>
    )
  }

  componentWillMount() {
    const stateID = this.props.stateID
    this.startWebsocket(stateID, MAX_RETRIES_ON_SOCKET_CLOSED)
  }

  startWebsocket = (stateID, retries) => {
    const sock = new WebSocket(gameStateWSocketAbsUrl(stateID))

    sock.onmessage = (e) => {
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

    sock.onerror = (err) => {
      this.props.pushError("Server Connection error: " + err.message)
      sock.close()
    }
  }
}

export default connect(
  null,
  { pushError }
)(Game)
