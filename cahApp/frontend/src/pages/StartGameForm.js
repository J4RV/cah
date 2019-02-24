import React, { Component } from "react"

import BackToGameListButton from '../components/BackToGameListButton'
import Button from '@material-ui/core/Button'
import Checkbox from "@material-ui/core/Checkbox"
import ExpansionsSelect from "../components/ExpansionsSelect"
import FormControl from "@material-ui/core/FormControl"
import FormControlLabel from "@material-ui/core/FormControlLabel"
import { TextField } from "@material-ui/core"
import Typography from "@material-ui/core/Typography"
import axios from "axios";
import { connect } from "react-redux"
import pushError from "../actions/pushError"
import { startGameUrl } from "../restUrls"
import { withStyles } from "@material-ui/core/styles"

const styles = theme => ({
  container: {
    padding: theme.spacing.unit * 2,
    maxWidth: 480,
    marginLeft: "auto",
    marginRight: "auto",
  },
  formLabel: {
    textAlign: "left",
  },
  form: {
    marginTop: theme.spacing.unit * 2,
    display: "inline-block",
    width: "100%",
    textAlign: "right",
  },
  button: {
    margin: theme.spacing.unit,
  },
})

const StartButton = className => (
  <Button
    variant="contained"
    color="primary"
    type="submit"
    className={className}
  >
    Start game
  </Button>
)

class StartGameForm extends Component {
  state = {
    gameID: this.props.gameID,
    expansions: [],
    handSize: 10,
    randomFirstCzar: true,
  }

  render() {
    const { classes, enoughPlayers } = this.props
    const { handSize, randomFirstCzar } = this.state
    return (
      <div className={classes.container}>
        <form className={classes.form} onSubmit={this.handleSubmit}>
          <Typography variant="h6" className={classes.formLabel} >
            Game Options
          </Typography>
          <FormControl fullWidth margin="normal">
            <ExpansionsSelect onSelectedChange={this.handleExpansionSelected} />
          </FormControl>
          <FormControl required fullWidth margin="normal">
            <TextField
              label="Hand size"
              id="handSize"
              name="handSize"
              type="number"
              onChange={this.handleHandSizeChange}
              value={handSize}
            />
          </FormControl>
          <FormControl fullWidth margin="normal">
            <FormControlLabel
              control={<Checkbox id="randomFirstCzar" name="randomFirstCzar" color="primary" value={randomFirstCzar} />}
              label="First Czar chosen randomly"
            />
          </FormControl>
          <BackToGameListButton className={classes.button} />
          {enoughPlayers
            ? <StartButton className={classes.button} />
            : null}
        </form>
      </div>
    )
  }

  handleSubmit = event => {
    event.preventDefault()
    if (this.props.enoughPlayers) {
      console.log("Starting game with options", this.state)
      axios.post(startGameUrl, this.state)
        .catch(e => this.props.pushError(e))
    } else {
      console.error("Tried to start a game without enough players")
    }
    return false
  }

  handleHandSizeChange = event => {
    let newValue = parseInt(event.target.value)
    newValue = Math.min(Math.max(newValue, 0), 30);
    this.setState({ ...this.state, handSize: newValue })
  }

  handleExpansionSelected = selected => {
    this.setState({ ...this.state, expansions: selected })
  }
}

export default connect(
  () => { },
  { pushError }
)(withStyles(styles)(StartGameForm))
