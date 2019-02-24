import React, { Component } from "react"

import Chip from "@material-ui/core/Chip"
import FormControl from "@material-ui/core/FormControl"
import InputLabel from "@material-ui/core/InputLabel"
import MenuItem from "@material-ui/core/MenuItem"
import Select from "@material-ui/core/Select"
import { availableExpansionsUrl } from "../restUrls"
import axios from "axios"
import { withStyles } from "@material-ui/core/styles"

const styles = theme => ({
  select: {
    width: "100%",
    minHeight: 36,
  },
  chips: {
    display: "flex",
    flexWrap: "wrap",
  },
  chip: {
    margin: theme.spacing.unit / 4,
  },
})

const ITEM_HEIGHT = 48
const ITEM_PADDING_TOP = 8
const MenuProps = {
  PaperProps: {
    style: {
      maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
      width: 250,
    },
  },
}

class ExpansionsSelect extends Component {
  state = {
    expansions: ["Loading..."],
    selected: [],
  }

  render() {
    const { classes } = this.props
    const { selected, expansions } = this.state
    return (
      <FormControl required className={classes.select}>
        <select multiple hidden id="selectedExpansions" name="selectedExpansions">
          {selected.map(name => (
            <option selected key={name} value={name}>
              {name}
            </option>
          ))}
        </select>
        <InputLabel>Expansions</InputLabel>
        <Select
          multiple
          displayEmpty
          value={selected}
          onChange={this.handleChangeSelect}
          renderValue={selected => (
            <div className={classes.chips}>
              {selected.map(value => (
                <Chip key={value} label={value} className={classes.chip} />
              ))}
            </div>
          )}
          MenuProps={MenuProps}
        >
          {expansions.map(name => (
            <MenuItem key={name} value={name}>
              {name}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    )
  }

  handleChangeSelect = event => {
    const newValue = event.target.value
    this.setState({
      ...this.state,
      selected: newValue,
    })
    this.props.onSelectedChange(newValue)
  }

  componentWillMount() {
    axios.get(availableExpansionsUrl).then(r =>
      this.setState({
        ...this.state,
        expansions: r.data,
      })
    )
  }
}

export default withStyles(styles)(ExpansionsSelect)
