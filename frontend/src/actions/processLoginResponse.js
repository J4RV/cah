export const PROCESS_LOGIN_RESPONSE = 'PROCESS_LOGIN_RESPONSE'

export const processLoginReduce = (state, action) => {
  let { response } = action.payload
  if (response.status !== 200) {
    return { ...state, validCookie: false }
  }
  const userInfo = response.data
  return {
    ...state,
    validCookie: true,
    userID: userInfo.id,
    username: userInfo.username,
  }
}

export default (validCookieResponse) => ({
  type: PROCESS_LOGIN_RESPONSE,
  payload: { response: validCookieResponse }
})