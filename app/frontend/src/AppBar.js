import AppBar from '@material-ui/core/AppBar'
import GitHubIcon from './icons/GitHub'
import Menu from './components/Menu';
import React from 'react'
import ShoppingCart from '@material-ui/icons/ShoppingCart'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import { connect } from 'react-redux'
import { withStyles } from '@material-ui/core/styles'
import withWidth from '@material-ui/core/withWidth'

const styles = theme => ({
  appbar: {
    color: theme.palette.blackcard.text,
    backgroundColor: theme.palette.blackcard.background,
  },
  title: {
    flexGrow: 1,
  },
  user: {
    margin: theme.spacing.unit,
  },
  icon: {
    margin: theme.spacing.unit,
    color: theme.palette.grey[50],
  },
})

function TopAppBar({ username, title, shortTitle, width, classes }) {
  return (
    <div>
      <AppBar position='static' className={classes.appbar} >
        <Toolbar>
          <Typography variant='h6' color='inherit' className={classes.user} >
            {username}
          </Typography>
          <Typography variant='h6' color='inherit' className={classes.title} >
            {width === "xs" ? shortTitle : title}
          </Typography>
          <Typography>
            <a target="blank" href="https://github.com/J4RV">
              <GitHubIcon className={classes.icon} />
            </a>
          </Typography>
          <Typography>
            <a target="blank" href="https://store.cardsagainsthumanity.com">
              <ShoppingCart className={classes.icon} />
            </a>
          </Typography>
          <Menu />
        </Toolbar>
      </AppBar>
    </div>
  )
}

export default connect(
  state => ({ username: state.username }),
)(withWidth()(withStyles(styles)(TopAppBar)))
