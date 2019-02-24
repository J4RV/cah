export const PUSH_ERROR = 'PUSH_ERROR'

const maxErrors = 4

export const pushErrorReduce = (state, action) => {
  let { msg } = action.payload
  console.error(msg)
  let newValue = state.errors.concat(msg)
  if (newValue.length > maxErrors) {
    newValue.splice(0, 1)
  }
  return { ...state, errors: newValue }
}

export default (error) => {
  const msg = error.response != null ? error.response.data : error.toString()
  return {
    type: PUSH_ERROR,
    payload: { msg: msg }
  }
}