import React from 'react'
import CloseIcon from '@material-ui/icons/Close'
import IconButton from "@material-ui/core/IconButton"

const CloseButton = (props) => (
  <IconButton
    key="close"
    aria-label="Close"
    color="inherit"
    onClick={props.onClick}
    className={props.className}
  >
    <CloseIcon />
  </IconButton>
)

export default CloseButton