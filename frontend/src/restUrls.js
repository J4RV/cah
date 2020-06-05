// User
export const loginUrl = "/api/user/login"
export const logoutUrl = "/api/user/logout"
export const registerUrl = "/api/user/register"
export const validCookieUrl = "/api/user/valid-cookie"

// Game state
export const gameStateUrl = (stateID) => `/api/gamestate/${stateID}/state`
export const playCardsUrl = (stateID) => `/api/gamestate/${stateID}/play-cards`
export const chooseWinnerUrl = (stateID) =>
  `/api/gamestate/${stateID}/choose-winner`

export const gameStateWSocketAbsUrl = (stateID) =>
  (document.location.protocol === "http:" ? "ws:" : "wss:") +
  `//${window.location.host}/api/gamestate/${stateID}/state-websocket`

// Game
export const openGamesUrl = "/api/game/list-open"
export const myGamesInProgress = "/api/game/list-in-progress"
export const createGameUrl = "/api/game/create"
export const joinGameUrl = "/api/game/join"
export const roomStateUrl = (gameID) => `/api/game/${gameID}/room-state`
export const startGameUrl = `/api/game/start`
export const availableExpansionsUrl = `/api/game/available-expansions`
