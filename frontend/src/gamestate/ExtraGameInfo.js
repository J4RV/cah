import React from "react"
import Typography from "@material-ui/core/Typography"
import { withStyles } from "@material-ui/core/styles"
import withWidth from "@material-ui/core/withWidth"

const styles = (theme) => ({
  container: {},
})

const ExtraGameInfo = ({ state, classes }) => {
  let roundText = state.currRound
  if (state.maxRounds) {
    roundText += " of " + state.maxRounds
  }
  return (
    <Typography align="center">
      <div className={classes.container}>
        <p>Round: {roundText}</p>
      </div>
    </Typography>
  )
}

export default withWidth()(withStyles(styles)(ExtraGameInfo))
