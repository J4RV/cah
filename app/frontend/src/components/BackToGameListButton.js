import Button from '@material-ui/core/Button'
import { Link } from 'react-router-dom'
import React from "react"

const BackToGameListButton = ({ className }) => (
  <Link to="/game/list/open">
    <Button className={className}>Back to games list</Button>
  </Link>
)

export default BackToGameListButton