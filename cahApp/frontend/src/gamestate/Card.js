import React from 'react'
import { withStyles } from '@material-ui/core/styles'
import withWidth from '@material-ui/core/withWidth'

const styles = theme => ({
  card: {
    position: "relative",
    display: "inline-block",
    padding: theme.spacing.unit,
    width: "8rem",
    height: "10rem",
    borderRadius: 10,
    textAlign: "center",
    verticalAlign: "top",
    //transition: "transform ease-in-out .5s",
    transformOrigin: "50% 80%",
    [theme.breakpoints.down('sm')]: {
      padding: theme.spacing.unit * 0.5,
      width: "6.4rem",
      height: "8rem",
      borderRadius: 8,
    }
  },
  inHand: {
    margin: "0 0 -8px 0",
  },
  text: {
    fontFamily: '"Open Sans", "Roboto", "Helvetica", "Arial", sans-serif',
    fontWeight: "600",
    fontSize: ".8rem",
    whiteSpace: "pre-wrap",
    [theme.breakpoints.down('sm')]: {
      fontSize: ".64rem",
    }
  },
  black: {
    color: theme.palette.blackcard.text,
    background: theme.palette.blackcard.background,
  },
  white: {
    color: theme.palette.whitecard.text,
    background: theme.palette.whitecard.background,
  },
  expansion: {
    position: "absolute",
    bottom: theme.spacing.unit,
    right: theme.spacing.unit * 2,
    marginLeft: theme.spacing.unit,
    color: theme.palette.expansion,
    fontSize: ".8em",
    textAlign: "right",
    [theme.breakpoints.down('sm')]: {
      right: theme.spacing.unit,
    }
  },
  inTable: {
    boxShadow: theme.shadows[1],
  },
  floating: {
    boxShadow: theme.shadows[10],
  },
  glowing: {
    boxShadow: theme.lights.glow,
  },
});

class Card extends React.Component {
  render() {
    const { text, isBlack, elevated, glowing, inHand, expansion, className, classes, style, ...rest } = this.props
    let shadowClass
    if (glowing) {
      shadowClass = classes.glowing
    } else {
      shadowClass = elevated ? classes.floating : classes.inTable
    }
    const colorClass = isBlack ? classes.black : classes.white
    return <div
      style={{ transform: `rotate(${this.state.rotation}deg)`, ...style }}
      className={`${classes.card} ${classes.text} ${colorClass} ${shadowClass}    
        ${inHand ? classes.inHand : ""}
        ${className ? className : ""}`}
      {...rest}
    >
      <div >{text}</div>
      <div className={classes.expansion}>{expansion}</div>
    </div>
  }
  randomRotate = () => {
    this.setState({ rotation: Math.random() * 5 - 2.5 })
    /*
    Too expensive for mobile, check page width and only do this for large screens
    window.setTimeout(this.randomRotate, Math.random() * 2000 + 1000)    
    */
  }
  componentWillMount() {
    this.randomRotate()
  }
}

export default withWidth()(withStyles(styles)(Card))
