import { PROCESS_LOGIN_RESPONSE, processLoginReduce } from './actions/processLoginResponse'
import { PUSH_ERROR, pushErrorReduce } from './actions/pushError'
import { REMOVE_ERROR, removeErrorReduce } from './actions/removeError'

const initState = {
  validCookie: undefined,
  userID: undefined,
  username: undefined,
  errors: [],
}

export default (state = initState, action) => {
  switch (action.type) {
    case PROCESS_LOGIN_RESPONSE:
      return processLoginReduce(state, action)
    case PUSH_ERROR:
      return pushErrorReduce(state, action)
    case REMOVE_ERROR:
      return removeErrorReduce(state, action)
    default:
      return state
  }
}