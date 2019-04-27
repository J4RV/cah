import React from 'react'
import Card from './Card'
import WhiteCardsPlayed from './WhiteCardsPlayed'

const Table = ({ state }) => {
  const card = state.blackCardInPlay
  if (card == null) return null
  return (
    <div className='cah-table'>
      <Card {...card} isBlack style={{ margin: '.5rem' }} />
      <WhiteCardsPlayed state={state} />
    </div>)
}

export default Table
