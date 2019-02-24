import React, { Component } from "react"

import Button from "@material-ui/core/Button"
import Card from "./Card"
import Check from "@material-ui/icons/Check"
import Fab from "@material-ui/core/Fab"
import Typography from "@material-ui/core/Typography"
import axios from "axios"
import { connect } from "react-redux"
import { playCardsUrl } from "../restUrls"
import pushError from "../actions/pushError"
import { withStyles } from "@material-ui/core/styles"
import withWidth from "@material-ui/core/withWidth"

const styles = theme => ({
  hand: {
    maxWidth: 800,
    boxShadow: theme.shadows[8],
    background: "#0004",
    padding: theme.spacing.unit,
    paddingTop: theme.spacing.unit * 2,
    marginLeft: "auto",
    marginRight: "auto",
    textAlign: "center",
  },
  cardsInHand: {
    textAlign: "center",
    paddingBottom: 8, // Inverse of card's negative top margin
  },
  largeScreenButton: {
    marginTop: theme.spacing.unit,
  },
  smallScreenButton: {
    position: "fixed",
    right: theme.spacing.unit,
    bottom: theme.spacing.unit,
  },
})

let PlayCardsButton = ({ classes, width, playCards }) => {
  if (width === "sm" || width === "xs") {
    return (
      <Fab
        aria-label="Play selected cards"
        color="primary"
        onClick={playCards}
        className={classes.smallScreenButton}
      >
        <Check />
      </Fab>
    )
  } else {
    return (
      <Button
        variant="contained"
        color="primary"
        onClick={playCards}
        className={classes.largeScreenButton}
      >
        Play cards
      </Button>
    )
  }
}
PlayCardsButton = withWidth()(PlayCardsButton)

const CardsToPlay = ({ state }) => {
  const isCzar = state.myPlayer.id === state.currentCzarID
  if (isCzar) {
    return null
  }
  const cardsToPlay =
    state.blackCardInPlay.blanks - state.myPlayer.whiteCardsInPlay.length
  if (cardsToPlay === 0) {
    return null
  }
  return (
    <Typography variant="h6" gutterBottom>
      Play {cardsToPlay} cards
    </Typography>
  )
}

class Hand extends Component {
  state = { cardIndexes: [], errormsg: null }

  render() {
    const { gamestate, classes } = this.props
    return (
      <div className={classes.hand}>
        <CardsToPlay state={gamestate} />
        <div className={classes.cardsInHand}>
          {gamestate.myPlayer.hand.map((c, i) => (
            <Card
              {...c}
              handIndex={i}
              elevated
              inHand
              glowing={this.state.cardIndexes.includes(i)}
              onClick={() => this.handleCardClick(i)}
            />
          ))}
        </div>
        {this.canPlayCards() ? (
          <PlayCardsButton playCards={this.playCards} classes={classes} />
        ) : null}
      </div>
    )
  }

  canPlayCards = () => {
    const gamestate = this.props.gamestate
    const isCzar = gamestate.currentCzarID === gamestate.myPlayer.id
    const validPhase = gamestate.phase === "Sinners playing their cards"
    return !isCzar && validPhase
  }

  handleCardClick = i => {
    if (!this.canPlayCards()) {
      return
    }
    let newList = this.state.cardIndexes.slice()
    if (newList.includes(i)) {
      newList.splice(newList.indexOf(i), 1)
    } else {
      newList.push(i)
    }
    this.setState({ cardIndexes: newList })
  }

  playCards = () => {
    const gamestate = this.props.gamestate
    axios
      .post(playCardsUrl(gamestate.id), {
        cardIndexes: this.state.cardIndexes,
      })
      .then(r => {
        this.setState({ cardIndexes: [] })
      })
      .catch(r => this.props.pushError(r))
  }
}

export default connect(
  () => {},
  { pushError }
)(withWidth()(withStyles(styles)(Hand)))
