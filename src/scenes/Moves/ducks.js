import { get, head, pick } from 'lodash';
import { GET_LOGGED_IN_USER_SUCCESS } from 'store/auth/actions';
import { fetchActive } from 'shared/utils';

import * as ReduxHelpers from 'shared/ReduxHelpers';
// Types
const SET_CONUS_STATUS = 'SET_CONUS_STATUS';
const SET_PENDING_MOVE_TYPE = 'SET_PENDING_MOVE_TYPE';
const SET_SELECTED_MOVE_TYPE = 'SET_SELECTED_MOVE_TYPE';

export const getMoveType = 'GET_MOVE';
export const GET_MOVE = ReduxHelpers.generateAsyncActionTypes(getMoveType);

export const createOrUpdateMoveType = 'CREATE_OR_UPDATE_MOVE';
export const CREATE_OR_UPDATE_MOVE = ReduxHelpers.generateAsyncActionTypes(createOrUpdateMoveType);

// Action creation
export function setPendingMoveType(value) {
  return { type: SET_PENDING_MOVE_TYPE, payload: value };
}

// TODO - deprecate this action & field
export function setSelectedMoveType(moveType) {
  return { type: SET_SELECTED_MOVE_TYPE, moveType };
}

// Reducer
const initialState = {
  currentMove: null,
  latestMove: null,
  pendingMoveType: null,
  hasSubmitError: false,
  hasSubmitSuccess: false,
  error: null,
};
function reshapeMove(move) {
  if (!move) return null;
  return pick(move, ['id', 'locator', 'orders_id', 'selected_move_type', 'status']);
}
export function moveReducer(state = initialState, action) {
  switch (action.type) {
    case GET_LOGGED_IN_USER_SUCCESS:
      const lastOrdersMoves = get(action.payload, 'service_member.orders.0.moves', []);
      const activeOrders = fetchActive(get(action.payload, 'service_member.orders'));

      const activeMove = fetchActive(get(activeOrders, 'moves'));

      return Object.assign({}, state, {
        latestMove: reshapeMove(head(lastOrdersMoves)),
        currentMove: reshapeMove(activeMove),
        hasLoadError: false,
        hasLoadSuccess: true,
      });
    case SET_PENDING_MOVE_TYPE:
      return Object.assign({}, state, {
        pendingMoveType: action.payload,
      });
    case CREATE_OR_UPDATE_MOVE.success:
      return Object.assign({}, state, {
        currentMove: reshapeMove(action.payload),
        latestMove: null,
        pendingMoveType: null,
        hasSubmitSuccess: true,
        hasSubmitError: false,
        error: null,
      });
    case CREATE_OR_UPDATE_MOVE.failure:
      return Object.assign({}, state, {
        currentMove: {},
        latestMove: null,
        hasSubmitSuccess: false,
        hasSubmitError: true,
        error: action.error,
      });
    case GET_MOVE.success:
      return Object.assign({}, state, {
        currentMove: reshapeMove(action.payload),
        latestMove: null,
        hasLoadSuccess: true,
        hasLoadError: false,
        error: null,
      });
    case GET_MOVE.failure:
      return Object.assign({}, state, {
        currentMove: {},
        latestMove: null,
        hasLoadSuccess: false,
        hasLoadError: true,
        error: action.error,
      });
    case SET_CONUS_STATUS:
      return Object.assign({}, state, {
        currentMove: {
          ...state.currentMove,
          conus_status: action.moveType,
        },
      });
    case SET_SELECTED_MOVE_TYPE:
      return Object.assign({}, state, {
        currentMove: {
          ...state.currentMove,
          selected_move_type: action.moveType,
        },
      });
    default:
      return state;
  }
}
