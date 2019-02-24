import GitHubIcon from './icons/GitHub'
import React from 'react'
import ShoppingCart from '@material-ui/icons/ShoppingCart'
import Typography from '@material-ui/core/Typography'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  linkContainer: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
  icon: {
    margin: theme.spacing.unit,
  },
});

const Footer = ({ classes }) => {
  return <div className={classes.linkContainer}>
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
  </div>
}

export default withStyles(styles)(Footer)