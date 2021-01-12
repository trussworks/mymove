import { get, isEmpty } from 'lodash';

import * as ReduxHelpers from 'shared/ReduxHelpers';
import { GET_LOGGED_IN_USER } from 'shared/Data/users';
import { fetchActive, fetchActivePPM } from 'shared/utils';
import { selectCurrentMove } from 'store/entities/selectors';

// Types
export const CREATE_OR_UPDATE_PPM = ReduxHelpers.generateAsyncActionTypes('CREATE_OR_UPDATE_PPM');
export const GET_PPM = ReduxHelpers.generateAsyncActionTypes('GET_PPM');
export const GET_SIT_ESTIMATE = ReduxHelpers.generateAsyncActionTypes('GET_SIT_ESTIMATE');

// Selectors
export function getMaxAdvance(state) {
  const maxIncentive = get(state, 'ppm.incentive_estimate_max');
  // we are using 20000000 since it is the largest number MacRae found that could be stored in table
  // and we don't want to block the user from requesting an advance if the rate engine fails
  return maxIncentive ? 0.6 * maxIncentive : 20000000;
}

export function getPPM(state) {
  const move = selectCurrentMove(state) || {};
  const moveId = move.id;
  const ppmFromEntities = Object.values(state.entities.personallyProcuredMoves).find((ppm) => ppm.move_id === moveId);
  const tempPPM = state.ppm.currentPpm;
  // temp fix while redux refactor is in progress when statuses aren't updated on ppms from both places
  const ppmStates = ['DRAFT', 'SUBMITTED', 'APPROVED', 'PAYMENT_REQUESTED', 'CANCELED'];

  if (!isEmpty(ppmFromEntities) && !isEmpty(tempPPM)) {
    const entitiesPPMStatus = ppmFromEntities.status;
    const tempPPMStatus = tempPPM.status;
    const indexOfEntitiesPPMStatus = ppmStates.indexOf(entitiesPPMStatus);
    const indexOfTempPPMStatus = ppmStates.indexOf(tempPPMStatus);

    if (entitiesPPMStatus === 'CANCELED') {
      return ppmFromEntities;
    } else if (tempPPMStatus === 'CANCELED') {
      return tempPPM;
    }

    if (indexOfEntitiesPPMStatus > indexOfTempPPMStatus) {
      return ppmFromEntities;
    }
    if (indexOfEntitiesPPMStatus < indexOfTempPPMStatus) {
      return tempPPM;
    }
    return ppmFromEntities;
  } else if (tempPPM) {
    return tempPPM;
  }
  return {};
}

// Reducer
const initialState = {
  incentive: null,
  sitReimbursement: null,
  pendingPpmWeight: null,
  currentPpm: null,
  hasSubmitError: false,
  hasSubmitSuccess: false,
  hasLoadSuccess: false,
  hasLoadError: false,
  hasEstimateSuccess: false,
  hasEstimateError: false,
  hasEstimateInProgress: false,
};
export function ppmReducer(state = initialState, action) {
  switch (action.type) {
    case GET_LOGGED_IN_USER.success:
      // Initialize state when we get the logged in user
      const activeOrders = fetchActive(get(action.payload, 'service_member.orders'));
      const activeMove = fetchActive(get(activeOrders, 'moves'));
      const activePpm = fetchActivePPM(get(activeMove, 'personally_procured_moves'));
      return Object.assign({}, state, {
        currentPpm: activePpm,
        pendingPpmWeight: get(activePpm, 'weight_estimate', null),
        incentive_estimate_min: get(activePpm, 'incentive_estimate_min', null),
        incentive_estimate_max: get(activePpm, 'incentive_estimate_max', null),
        sitReimbursement: get(activePpm, 'estimated_storage_reimbursement', null),
        hasLoadSuccess: true,
        hasLoadError: false,
      });
    case CREATE_OR_UPDATE_PPM.start:
      return Object.assign({}, state, {
        hasSubmitSuccess: false,
        hasSubmitInProgress: true,
      });
    case CREATE_OR_UPDATE_PPM.success:
      return Object.assign({}, state, {
        currentPpm: action.payload,
        sitReimbursement: get(action.payload, 'estimated_storage_reimbursement', null),
        pendingPpmWeight: null,
        hasSubmitSuccess: true,
        hasSubmitError: false,
        hasSubmitInProgress: false,
        ...action.payload,
      });
    case CREATE_OR_UPDATE_PPM.failure:
      return Object.assign({}, state, {
        hasSubmitSuccess: false,
        hasSubmitError: true,
        hasSubmitInProgress: false,
        error: action.error,
      });
    case GET_PPM.start:
      return Object.assign({}, state, {
        hasLoadSuccess: false,
      });
    case GET_PPM.success:
      return Object.assign({}, state, {
        currentPpm: get(action.payload, '0', null),
        pendingPpmWeight: get(action.payload, '0.weight_estimate', null),
        incentive_estimate_min: get(action.payload, '0.incentive_estimate_min', null),
        incentive_estimate_max: get(action.payload, '0.incentive_estimate_max', null),
        sitReimbursement: get(action.payload, '0.estimated_storage_reimbursement', null),
        hasLoadSuccess: true,
        hasLoadError: false,
      });
    case GET_PPM.failure:
      return Object.assign({}, state, {
        currentPpm: null,
        hasLoadSuccess: false,
        hasLoadError: true,
        error: action.error,
      });
    default:
      return state;
  }
}
