// User
export const loginUrl = "user/login"
export const logoutUrl = "user/logout"
export const registerUrl = "user/register"
export const validCookieUrl = "user/valid-cookie"

// Game state
export const gameStateUrl = (stateID) => `gamestate/${stateID}/state`
export const playCardsUrl = (stateID) => `gamestate/${stateID}/play-cards`
export const chooseWinnerUrl = (stateID) => `gamestate/${stateID}/choose-winner`

// Game
export const openGamesUrl = "game/list-open"
export const myGamesInProgress = "game/list-in-progress"
export const createGameUrl = "game/create"
export const joinGameUrl = "game/join"
export const roomStateUrl = (gameID) => `game/${gameID}/room-state`
export const startGameUrl = `game/start`
export const availableExpansionsUrl = `game/available-expansions`