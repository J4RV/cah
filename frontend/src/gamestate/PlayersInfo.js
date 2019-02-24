import React from 'react'
import Typography from '@material-ui/core/Typography'
import { withStyles } from "@material-ui/core/styles"
import withWidth from '@material-ui/core/withWidth'

const styles = theme => ({
  container: {
    display: "flex",
    flexWrap: "wrap",
    [theme.breakpoints.up('md')]: {
      position: "fixed",
      right: 8,
      top: 72,
    },
  },
  playerInfo: {
    background: "#EEE8",
    color: "#111",
    margin: 4,
    padding: 4,
    borderRadius: 3,
    boxShadow: theme.shadows[8],
    flexGrow: 1,
  },
})

const PlayerInfo = ({ player, itsYou, isCzar, classes }) => (
  <div className={classes.playerInfo}>
    <div>{player.name} {itsYou ? <b>(You)</b> : null}</div>
    <div>{player.points.length} points</div>
    <div>{isCzar ? <b>Current Czar</b> : `${player.whiteCardsInPlay} card(s) in play`}</div>
  </div>
)

const PlayersInfo = ({ state, classes }) => {
  return <Typography>
    <div className={classes.container}>
      {state.players.map((p) =>
        <PlayerInfo
          player={p}
          itsYou={p.id === state.myPlayer.id}
          isCzar={p.id === state.currentCzarID}
          classes={classes}
        />
      )}
    </div>
  </Typography>
}

export default withWidth()(withStyles(styles)(PlayersInfo))
