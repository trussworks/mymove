import React from 'react';
import PropTypes from 'prop-types';
import { Route, Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { get } from 'lodash';

import LoadingPlaceholder from 'shared/LoadingPlaceholder';
import { fetchAccessCode as fetchAccessCodeAction } from 'shared/Entities/modules/accessCodes';
import { selectServiceMemberFromLoggedInUser } from 'store/entities/selectors';
import { selectGetCurrentUserIsLoading, selectIsLoggedIn } from 'store/auth/selectors';

class CustomerPrivateRoute extends React.Component {
  componentDidMount() {
    const { fetchAccessCode } = this.props;
    fetchAccessCode();
  }

  render() {
    const { loginIsLoading, userIsLoggedIn, requiresAccessCode, accessCode, path, ...routeProps } = this.props;
    if (loginIsLoading) return <LoadingPlaceholder />;

    if (!userIsLoggedIn) return <Redirect to="/sign-in" />;

    if (userIsLoggedIn && requiresAccessCode && !accessCode) return <Redirect to="/access-code" />;

    // eslint-disable-next-line react/jsx-props-no-spreading
    return <Route {...routeProps} />;
  }
}

CustomerPrivateRoute.propTypes = {
  fetchAccessCode: PropTypes.func.isRequired,
  loginIsLoading: PropTypes.bool,
  userIsLoggedIn: PropTypes.bool,
  requiresAccessCode: PropTypes.bool,
  accessCode: PropTypes.string,
  path: PropTypes.string,
};

CustomerPrivateRoute.defaultProps = {
  loginIsLoading: true,
  userIsLoggedIn: false,
  requiresAccessCode: false,
  accessCode: undefined,
  path: undefined,
};

const mapStateToProps = (state) => {
  const serviceMember = selectServiceMemberFromLoggedInUser(state);
  const accessCodes = get(state, 'entities.accessCodes');

  return {
    loginIsLoading: selectGetCurrentUserIsLoading(state),
    userIsLoggedIn: selectIsLoggedIn(state),
    requiresAccessCode: serviceMember?.requires_access_code,
    accessCode: accessCodes && Object.values(accessCodes).length > 0 ? Object.values(accessCodes)[0].code : null,
  };
};

const mapDispatchToProps = { fetchAccessCode: fetchAccessCodeAction };

export default connect(mapStateToProps, mapDispatchToProps)(CustomerPrivateRoute);
