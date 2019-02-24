import './App.css'

import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import React, { Component } from 'react'
import { Redirect, Route, BrowserRouter as Router, withRouter } from 'react-router-dom'

import AppBar from './AppBar'
import CssBaseline from '@material-ui/core/CssBaseline'
import GameCreate from './pages/GameCreate'
import GameRoom from './pages/GameRoom'
import LoggedIn from './pages/LoggedIn'
import MyGamesInProgress from './pages/MyGamesInProgress'
import OpenGames from './pages/OpenGames'
import { Provider } from 'react-redux'
import WithErrors from './pages/WithErrors'
import { createStore } from 'redux'
import cyan from '@material-ui/core/colors/cyan'
import grey from '@material-ui/core/colors/grey'
import red from '@material-ui/core/colors/red'
import reducer from './reducer'

const theme = createMuiTheme({
  palette: {
    primary: red,
    secondary: grey,
    whitecard: { text: "#161616", background: "#FAFAFA" },
    blackcard: { text: "#FAFAFA", background: "#161616" },
    expansion: "#888888",
    type: 'dark',
  },
  lights: {
    glow: `0 0 4px 2px ${cyan[100]}, 0 0 24px 2px ${cyan[500]}`,
  },
})

const RealApp = withRouter(() => (
  <WithErrors>
    <LoggedIn>
      <AppBar />
      <Route exact path="/" render={() => <Redirect to="/game/list/open" />} />
      <Route path="/game/list/create" component={GameCreate} />
      <Route path="/game/list/my-games-in-progress" component={MyGamesInProgress} />
      <Route path="/game/list/open" component={OpenGames} />
      <Route path="/game/room/:gameID" component={GameRoom} />
    </LoggedIn>
  </WithErrors>
))

class App extends Component {
  render() {
    return <Router>
      <MuiThemeProvider theme={theme}>
        <Provider store={createStore(reducer)}>
          <CssBaseline />
          <RealApp />
        </Provider>
      </MuiThemeProvider>
    </Router>
  }
}

export default App
