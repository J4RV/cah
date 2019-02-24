import { Button, Typography } from '@material-ui/core';
import React, { Component } from 'react'

import GamesTable from '../components/GamesTable'
import { Link } from 'react-router-dom'
import Paper from '@material-ui/core/Paper'
import { connect } from 'react-redux'
import { openGamesUrl } from '../restUrls'
import pushError from '../actions/pushError'
import { withStyles } from '@material-ui/core/styles'
import withWidth from '@material-ui/core/withWidth';

const styles = theme => ({
  root: {
    maxWidth: 960 - theme.spacing.unit * 4,
    marginTop: theme.spacing.unit * 3,
    marginLeft: "auto",
    marginRight: "auto",
    padding: theme.spacing.unit * 2,
    [theme.breakpoints.down('sm')]: {
      padding: theme.spacing.unit,
    },
  },
  tableContainer: {
    overflowX: 'auto',
  },
  title: {
    textAlign: "center",
    marginBottom: theme.spacing.unit * 2,
  },
  createBtn: {
    float: "right",
    marginTop: theme.spacing.unit * 2,
  },
})

class OpenGames extends Component {
  render() {
    const { classes } = this.props
    return <div className={classes.root}>
      <Typography variant="h5" className={classes.title}>
        Open games
      </Typography>
      <Paper className={classes.tableContainer}>
        <GamesTable fetchGamesUrl={openGamesUrl} />
      </Paper>
      <Link to="/game/list/create">
        <Button
          type="button"
          className={classes.createBtn}>
          Create new game
        </Button>
      </Link>
    </div>
  }
}

export default connect(
  null,
  { pushError }
)(withWidth()(withStyles(styles)(OpenGames)))