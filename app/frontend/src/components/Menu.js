import withWidth, { isWidthUp } from '@material-ui/core/withWidth'

import IconButton from '@material-ui/core/IconButton';
import { Link } from 'react-router-dom'
import Menu from '@material-ui/core/Menu';
import MenuIcon from '@material-ui/icons/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import React from 'react'
import { logoutUrl } from '../restUrls'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  menuButton: {
    marginRight: -12,
    marginLeft: theme.spacing.unit * 2,
  },
});

const MenuElements = ({ onElementClick }) => (
  <React.Fragment>
    <Link to="/game/list/my-games-in-progress">
      <MenuItem onClick={onElementClick}>My games in progress</MenuItem>
    </Link>
    <Link to="/game/list/open">
      <MenuItem onClick={onElementClick}>Open games</MenuItem>
    </Link>
    <a href={logoutUrl}>
      <MenuItem onClick={onElementClick}>Logout</MenuItem>
    </a>
  </React.Fragment>
)

class AppBarMenu extends React.Component {
  state = { open: false }

  render() {
    const { classes, width } = this.props
    const { open } = this.state

    if (isWidthUp("md", width, true)) {
      return <MenuElements />
    }

    return (
      <div>
        <IconButton
          className={classes.menuButton}
          color="inherit"
          aria-label="Menu"
          onClick={this.handleToggle}
          buttonRef={node => this.anchorEl = node}
        >
          <MenuIcon />
        </IconButton>
        <Menu
          open={open}
          anchorEl={this.anchorEl}
          onClose={this.handleClose}
        >
          <MenuElements onElementClick={this.handleClose} />
        </Menu>
      </div>
    )
  }

  handleToggle = () => {
    this.setState(state => ({ open: !state.open }))
  }

  handleClose = event => {
    if (this.anchorEl.contains(event.target)) {
      // anchorEl is the Menu button, which will toggle the open property on click
      return;
    }
    this.setState({ open: false })
  }
}

export default withWidth()(withStyles(styles)(AppBarMenu))
