import React, { Component } from "react"

import Hand from "../gamestate/Hand"
import PlayersInfo from "../gamestate/PlayersInfo"
import Table from "../gamestate/Table"
import { connect } from "react-redux"
import { gameStateWSocketAbsUrl } from "../restUrls"
import pushError from "../actions/pushError"

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
    const sock = new WebSocket(gameStateWSocketAbsUrl(stateID))
    sock.onmessage = e => {
      console.debug("updating game state", e)
      this.setState(JSON.parse(e.data))
    }
  }
}

export default connect(
  null,
  { pushError }
)(Game)
