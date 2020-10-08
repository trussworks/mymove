import { takeLatest, put, call } from 'redux-saga/effects';

import {
  INIT_ONBOARDING,
  FETCH_CUSTOMER_DATA,
  initOnboardingFailed,
  initOnboardingComplete,
} from 'store/onboarding/actions';
import {
  getLoggedInUser,
  getMTOShipmentsForMove,
  createServiceMember as createServiceMemberApi,
} from 'services/internalApi';
import { addEntities } from 'shared/Entities/actions';
import { CREATE_SERVICE_MEMBER } from 'scenes/ServiceMembers/ducks';

export function* fetchCustomerData() {
  // First load the user & store in entities
  const user = yield call(getLoggedInUser);
  yield put(addEntities(user));

  // TODO - fork/spawn additional API calls
  // Load MTO shipments if there is a move
  const { moves } = user;
  if (moves && Object.keys(moves).length > 0) {
    const [moveId] = Object.keys(moves);
    // User has a move, load MTO shipments & store in entities
    const mtoShipments = yield call(getMTOShipmentsForMove, moveId);
    yield put(addEntities(mtoShipments));
  }

  return user;
}

export function* watchFetchCustomerData() {
  yield takeLatest(FETCH_CUSTOMER_DATA, fetchCustomerData);
}

export function* createServiceMember() {
  try {
    yield put({ type: CREATE_SERVICE_MEMBER.start });
    const serviceMember = yield call(createServiceMemberApi);
    yield put({ type: CREATE_SERVICE_MEMBER.success, payload: serviceMember });
    yield call(fetchCustomerData);
  } catch (e) {
    yield put({ type: CREATE_SERVICE_MEMBER.failure, error: e });
  }
}

export function* initializeOnboarding() {
  try {
    const user = yield call(fetchCustomerData);
    if (!user.serviceMembers) {
      yield call(createServiceMember);
    }
    yield put(initOnboardingComplete());
    yield call(watchFetchCustomerData);
  } catch (error) {
    yield put(initOnboardingFailed(error));
  }
}

export function* watchInitializeOnboarding() {
  yield takeLatest(INIT_ONBOARDING, initializeOnboarding);
}
