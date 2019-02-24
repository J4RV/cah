import React, { Component } from 'react'

import Button from '@material-ui/core/Button'
import DialogActions from '@material-ui/core/DialogActions'
import { Redirect } from 'react-router-dom'
import TextField from '@material-ui/core/TextField'
import Typography from '@material-ui/core/Typography'
import axios from 'axios'
import { connect } from "react-redux"
import { createGameUrl } from '../restUrls'
import pushError from "../actions/pushError"
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  form: {
    padding: theme.spacing.unit * 2,
    maxWidth: 360,
    marginLeft: "auto",
    marginRight: "auto",
  },
})

class CreateGameDialog extends Component {
  state = { name: "", password: "", waitingResponse: false, createdSuccessfully: false }

  render() {
    const { classes, history } = this.props
    if (this.state.createdSuccessfully) {
      return <Redirect to="/game/list/open" />
    }
    return (
      <form className={classes.form} onSubmit={this.handleSubmit} >
        <Typography variant="h6" className={classes.formLabel} >
          Create a new game
        </Typography>
        <TextField required fullWidth margin="normal"
          label="Room name"
          autoComplete="roomName" s
          onChange={this.handleChangeName}
        />
        {/*<TextField fullWidth margin="normal"
            label="Room password"
            autoComplete="roomPassword"
            onChange={this.handleChangePass}
          />*/}
        <DialogActions>
          <Button margin="normal"
            onClick={history.goBack}
            disabled={this.state.waitingResponse}
          >Cancel</Button>
          <Button margin="normal" autoFocus
            type="submit"
            variant="contained"
            color="primary"
            disabled={this.state.waitingResponse}
          >Create Game</Button>
        </DialogActions>
      </form>
    )
  }

  handleSubmit = (event) => {
    event.preventDefault() // necessary for firefox, or submitting the form will make the page reload!
    this.setState({ ...this.state, waitingResponse: true })
    let payload = {
      name: this.state.name,
      password: this.state.password,
    }
    axios.post(createGameUrl, payload)
      .then(this.createdSuccessfully)
      .catch(r => {
        this.props.pushError(r)
        this.setState({
          ...this.state,
          waitingResponse: false
        })
      })
    return false
  }

  createdSuccessfully = () => {
    this.setState({ ...this.state, createdSuccessfully: true })
  }

  handleChangeName = (event) => {
    let newState = Object.assign({}, this.state)
    newState.name = event.target.value.trim()
    this.setState(newState)
  }

  handleChangePass = (event) => {
    let newState = Object.assign({}, this.state)
    newState.password = event.target.value.trim()
    this.setState(newState)
  }
}

export default connect(
  null,
  { pushError }
)(withStyles(styles)(CreateGameDialog))