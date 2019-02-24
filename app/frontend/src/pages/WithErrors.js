import CloseButton from '../components/CloseButton'
import ErrorIcon from '@material-ui/icons/Error'
import React from 'react'
import Snackbar from '@material-ui/core/Snackbar'
import SnackbarContent from '@material-ui/core/SnackbarContent'
import { connect } from 'react-redux'
import removeError from '../actions/removeError';
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  error: {
    color: theme.palette.getContrastText(theme.palette.error.dark),
    background: theme.palette.error.dark,
    display: 'flex',
    alignItems: 'center',
  },
  icon: {
    marginRight: theme.spacing.unit,
  },
  message: {
    display: 'flex',
    alignItems: 'center',
  },
})

const WithErrors = ({ errors, children, removeError, classes }) => (
  <React.Fragment>
    {children}
    <Snackbar
      anchorOrigin={{
        vertical: 'bottom',
        horizontal: 'left',
      }}
      open={errors != null && errors.length > 0}
    >
      <SnackbarContent
        className={classes.error}
        message={
          errors.map((error, i) =>
            <div className={classes.message}>
              <ErrorIcon className={classes.icon} />
              {error}
              <CloseButton onClick={() => removeError(i)} />
            </div>
          )
        }
      />
    </Snackbar >
  </React.Fragment>
)

export default connect(
  state => ({ errors: state.errors }),
  { removeError }
)(withStyles(styles)(WithErrors))