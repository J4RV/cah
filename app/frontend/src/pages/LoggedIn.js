import React, { Component } from "react"
import { loginUrl, registerUrl, validCookieUrl } from "../restUrls"

import Button from "@material-ui/core/Button"
import Card from "../gamestate/Card"
import CircularProgress from '@material-ui/core/CircularProgress'
import Footer from "../Footer"
import FormControl from "@material-ui/core/FormControl"
import TextField from "@material-ui/core/TextField"
import Typography from "@material-ui/core/Typography"
import axios from "axios"
import { connect } from "react-redux"
import processLoginResponse from "../actions/processLoginResponse"
import pushError from "../actions/pushError"
import { withStyles } from "@material-ui/core/styles"

const styles = theme => ({
  container: {
    textAlign: "center",
    marginTop: theme.spacing.unit * 2,
  },
  form: {
    maxWidth: 260,
    marginTop: theme.spacing.unit * 2,
    marginBottom: theme.spacing.unit * 2,
    padding: theme.spacing.unit * 2,
    display: "inline-block",
  },
})

class LoginForm extends Component {
  state = { username: "", password: "", disabled: false }

  handleChangeUser = event => {
    let newState = Object.assign({}, this.state)
    newState.username = event.target.value.trim()
    this.setState(newState)
  }

  handleChangePass = event => {
    let newState = Object.assign({}, this.state)
    newState.password = event.target.value.trim()
    this.setState(newState)
  }

  handleSubmit = (event, url) => {
    event.preventDefault() // necessary for firefox, or submitting the form will make the page reload!
    this.setState({ ...this.state, disabled: true })
    let payload = {
      username: this.state.username,
      password: this.state.password,
    }
    axios
      .post(url, payload)
      .then(r => this.props.onSubmitResponse(r))
      .catch(r => {
        this.setState({
          ...this.state,
          disabled: false,
        })
        this.props.onError(r.response.data)
      })
    return false
  }

  render() {
    const classes = this.props.classes
    return (
      <div className={classes.container}>
        <Typography variant="h2" gutterBottom>
          Cards Against Humanity
        </Typography>
        <Typography variant="h4" gutterBottom>
          A party game for horrible people.
        </Typography>
        <form
          className={classes.form}
          onSubmit={e => this.handleSubmit(e, loginUrl)}
        >
          <Card
            isBlack
            text="I'm _ and my password is _."
            expansion="Security questions"
          />
          <TextField
            required
            fullWidth
            margin="normal"
            label="Username"
            autoComplete="username"
            onChange={this.handleChangeUser}
          />
          <TextField
            required
            fullWidth
            margin="normal"
            label="Password"
            type="password"
            autoComplete="password"
            onChange={this.handleChangePass}
          />
          <FormControl margin="normal" fullWidth>
            <Button
              type="submit"
              variant="contained"
              color="primary"
              onClick={e => this.handleSubmit(e, loginUrl)}
              disabled={this.state.disabled}
            >
              Log in
            </Button>
          </FormControl>
          <FormControl margin="normal" fullWidth>
            <Button
              type="button"
              variant="contained"
              onClick={e => this.handleSubmit(e, registerUrl)}
              disabled={this.state.disabled}
            >
              Register
            </Button>
          </FormControl>
          <Footer />
        </form>
      </div>
    )
  }
}

class LoggedInControl extends Component {
  componentWillMount() {
    axios
      .get(validCookieUrl)
      .then(r => this.props.processLoginResponse(r))
      .catch(r => this.props.processLoginResponse(r))
  }
  render() {
    const { validCookie, processLoginResponse, pushError, classes } = this.props
    if (validCookie == null) {
      return <CircularProgress />
    }
    if (validCookie) {
      return this.props.children
    }
    return (
      <LoginForm onSubmitResponse={processLoginResponse} onError={pushError} classes={classes} />
    )
  }
}

/*
using withRouter to prevent connect to block updates
see: https://github.com/ReactTraining/react-router/blob/master/packages/react-router/docs/guides/blocked-updates.md
*/
export default connect(
  state => ({ validCookie: state.validCookie }),
  { processLoginResponse, pushError }
)(withStyles(styles)(LoggedInControl))
