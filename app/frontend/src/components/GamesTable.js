import { Link, Redirect } from 'react-router-dom'
import React, { Component } from 'react'

import { Button } from '@material-ui/core';
import CircularProgress from '@material-ui/core/CircularProgress'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import TableCell from '@material-ui/core/TableCell'
import TableHead from '@material-ui/core/TableHead'
import TableRow from '@material-ui/core/TableRow'
import Typography from "@material-ui/core/Typography"
import axios from 'axios'
import { connect } from 'react-redux'
import { joinGameUrl } from '../restUrls'
import pushError from '../actions/pushError'

/*import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
})*/

const PrimaryButton = ({ username, game, joinGame }) => {
  if (game.players.includes(username)) {
    return <Link to={`/game/room/${game.id}`}>
      <Button
        color="primary"
        variant="contained">
        Enter
      </Button>
    </Link>
  } else {
    return <Button
      color="primary"
      variant="contained"
      onClick={() => joinGame(game.id)}>
      Join
    </Button>
  }
}

class GamesTable extends Component {
  state = { games: undefined, joinedGame: false }

  render() {
    if (this.state.joinedGame) {
      return <Redirect to={`/game/room/${this.state.joinedGame}`} />
    }
    const { username } = this.props
    const { games } = this.state
    if (games == null) {
      return <CircularProgress />
    }
    if (games.length === 0) {
      return <Typography variant="h6" align="center">
        No games found
      </Typography>
    }
    return (
      <Table>
        <TableHead>
          <TableRow>
            <TableCell align="left">Name</TableCell>
            <TableCell align="center">Owner</TableCell>
            {/*<TableCell align="right">Has password</TableCell>*/}
            {/*<TableCell align="left">Current players</TableCell>*/}
            <TableCell align="center">Actions</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {games.map(game => (
            <TableRow key={game.id}>
              <TableCell align="left">{game.name}</TableCell>
              <TableCell align="center">{game.owner}</TableCell>
              {/*<TableCell align="right">{game.hasPassword ? "Yes" : "No"}</TableCell>*/}
              {/*<TableCell align="left">{game.players.join(", ")}</TableCell>*/}
              <TableCell align="center">
                <PrimaryButton game={game} joinGame={this.joinGame} username={username} />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    )
  }

  componentWillMount() {
    this.refreshGames()
  }

  refreshGames = () => {
    axios.get(this.props.fetchGamesUrl)
      .then(r => {
        this.setState({ ...this.state, games: r.data })
      })
      .catch(e => this.props.pushError(e))
  }

  joinGame = (gameID) => {
    axios.post(joinGameUrl, { id: gameID })
      .then(this.setState({ ...this.state, joinedGame: gameID }))
      .catch(e => this.props.pushError(e))
  }
}

export default connect(
  state => ({ username: state.username }),
  { pushError }
)(/*withStyles(styles)*/(GamesTable))