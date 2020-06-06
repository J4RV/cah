import React, { Component } from "react"

import CircularProgress from "@material-ui/core/CircularProgress"
import axios from "axios"
import { connect } from "react-redux"
import processLoginResponse from "../actions/processLoginResponse"
import pushError from "../actions/pushError"
import { validCookieUrl } from "../restUrls"

class LoggedInControl extends Component {
  componentWillMount() {
    axios
      .get(validCookieUrl)
      .then((r) => this.props.processLoginResponse(r))
      .catch((r) => this.props.processLoginResponse(r))
  }
  render() {
    const { validCookie } = this.props
    if (validCookie == null) {
      return <CircularProgress />
    }
    if (validCookie) {
      return this.props.children
    }
    window.location.href = "/login"
  }
}

/*
using withRouter to prevent connect to block updates
see: https://github.com/ReactTraining/react-router/blob/master/packages/react-router/docs/guides/blocked-updates.md
*/
export default connect((state) => ({ validCookie: state.validCookie }), {
  processLoginResponse,
  pushError,
})(LoggedInControl)
