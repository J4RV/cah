import React, { Component } from 'react'

import BackToGameListButton from '../components/BackToGameListButton'
import CircularProgress from '@material-ui/core/CircularProgress'
import Game from './Game'
import StartGameForm from './StartGameForm';
import Typography from '@material-ui/core/Typography'
import axios from 'axios'
import { connect } from 'react-redux'
import pushError from "../actions/pushError"
import { roomStateUrl } from '../restUrls'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  container: {
    textAlign: "center",
    marginTop: theme.spacing.unit * 2,
  },
  button: {
    margin: theme.spacing.unit,
  },
  loading: {
    margin: theme.spacing.unit,
  },
});

class GameRoom extends Component {
  render() {
    const { classes, username } = this.props
    if (this.state == null) {
      return <CircularProgress className={classes.loading} />
    }
    const { room } = this.state
    if (room.phase !== "Not started") {
      return <Game stateID={room.stateID} />
    }
    const enoughPlayers = room.players.length > 2
    const imOwner = room.owner === username
    return (
      <div className={classes.container}>
        <Typography variant="h4" gutterBottom>
          {room.name}
        </Typography>
        {enoughPlayers
          ? <Typography variant="h6" gutterBottom>Waiting for the game creator to start the game</Typography>
          : <Typography variant="h6" gutterBottom>Waiting for more players to join</Typography>}
        <Typography>
          Creator: {room.owner}.
        </Typography>
        <Typography gutterBottom>
          Players: {room.players.join(", ")}.
        </Typography>
        {imOwner
          ? <StartGameForm gameID={room.id} enoughPlayers={enoughPlayers} />
          : <BackToGameListButton className={classes.button} />}
      </div>
    );
  }

  componentWillMount() {
    this.updateState()
  }

  updateState = () => {
    const gameID = this.props.match.params.gameID
    axios.get(roomStateUrl(gameID))
      .then(r => {
        this.setState({ room: r.data })
        if (r.data.phase === "Not started") {
          // this would be much better with websockets
          window.setTimeout(this.updateState, 5000)
        }
      })
      .catch(e => this.props.pushError(e))
  }
}

export default connect(
  state => ({ username: state.username }),
  { pushError }
)(withStyles(styles)(GameRoom))