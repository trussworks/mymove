/* eslint-disable react/jsx-props-no-spreading */
import React, { Component, lazy, Suspense } from 'react';
import PropTypes from 'prop-types';
import { Route, Switch, withRouter, matchPath, Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import classnames from 'classnames';

import 'uswds';
import '../../../node_modules/uswds/dist/css/uswds.css';
import 'scenes/Office/office.scss';

// API / Redux actions
import { loadUser as loadUserAction } from 'store/auth/actions';
import { selectCurrentUser } from 'shared/Data/users';
import {
  loadInternalSchema as loadInternalSchemaAction,
  loadPublicSchema as loadPublicSchemaAction,
} from 'shared/Swagger/ducks';
// Shared layout components
import ConnectedLogoutOnInactivity from 'shared/User/LogoutOnInactivity';
import PrivateRoute from 'containers/PrivateRoute';
import SomethingWentWrong from 'shared/SomethingWentWrong';
import { QueueHeader } from 'shared/Header/Office';
import FOUOHeader from 'components/FOUOHeader';
import { ConnectedSelectApplication } from 'pages/SelectApplication/SelectApplication';
import { roleTypes } from 'constants/userRoles';
import LoadingPlaceholder from 'shared/LoadingPlaceholder';
import { withContext } from 'shared/AppContext';
import { LocationShape, UserRolesShape } from 'types/index';

// Lazy load these dependencies (they correspond to unique routes & only need to be loaded when that URL is accessed)
const SignIn = lazy(() => import('shared/User/SignIn'));
// PPM pages (TODO move into src/pages)
const MoveInfo = lazy(() => import('scenes/Office/MoveInfo'));
const Queues = lazy(() => import('scenes/Office/Queues'));
const OrdersInfo = lazy(() => import('scenes/Office/OrdersInfo'));
const DocumentViewer = lazy(() => import('scenes/Office/DocumentViewer'));
// TXO
const TXOMoveInfo = lazy(() => import('pages/Office/TXOMoveInfo/TXOMoveInfo'));
// TOO pages (TODO move into src/pages)
const TOO = lazy(() => import('scenes/Office/TOO/too'));
const MoveQueue = lazy(() => import('pages/Office/MoveQueue/MoveQueue'));
const CustomerDetails = lazy(() => import('scenes/Office/TOO/customerDetails'));
// TIO pages
const PaymentRequestIndex = lazy(() => import('pages/Office/PaymentRequestIndex/PaymentRequestIndex'));
const PaymentRequestQueue = lazy(() => import('pages/Office/PaymentRequestQueue/PaymentRequestQueue'));

export class OfficeApp extends Component {
  constructor(props) {
    super(props);

    this.state = {
      hasError: false,
      error: undefined,
      info: undefined,
    };
  }

  componentDidMount() {
    document.title = 'Transcom PPP: Office';

    const { loadUser, loadInternalSchema, loadPublicSchema } = this.props;

    loadInternalSchema();
    loadPublicSchema();
    loadUser();
  }

  componentDidCatch(error, info) {
    this.setState({
      hasError: true,
      error,
      info,
    });
  }

  render() {
    const { hasError, error, info } = this.state;
    const {
      activeRole,
      userIsLoggedIn,
      userRoles,
      location: { pathname },
    } = this.props;

    const selectedRole = userIsLoggedIn && activeRole;

    // TODO - test login page?

    // TODO - I don't love this solution but it will work for now. Ideally we can abstract the page layout into a separate file where each route can use it or not
    // Don't show Header on OrdersInfo or DocumentViewer pages (PPM only)
    const hideHeaderPPM =
      selectedRole === roleTypes.PPM &&
      (matchPath(pathname, {
        path: '/moves/:moveId/documents/:moveDocumentId?',
        exact: true,
      }) ||
        matchPath(pathname, {
          path: '/moves/:moveId/orders',
          exact: true,
        }));

    const displayChangeRole =
      userIsLoggedIn &&
      userRoles?.length > 1 &&
      !matchPath(pathname, {
        path: '/select-application',
        exact: true,
      });

    const ppmRoutes = [
      <PrivateRoute
        key="ppmMoveOrdersRoute"
        path="/moves/:moveId/orders"
        component={OrdersInfo}
        requiredRoles={[roleTypes.PPM]}
      />,
      <PrivateRoute
        key="ppmMoveDocumentRoute"
        path="/moves/:moveId/documents/:moveDocumentId?"
        component={DocumentViewer}
        requiredRoles={[roleTypes.PPM]}
      />,
    ];

    const txoRoutes = [
      <PrivateRoute
        key="txoMoveInfoRoute"
        path="/moves/:moveOrderId"
        component={TXOMoveInfo}
        requiredRoles={[roleTypes.TOO, roleTypes.TIO]}
      />,
    ];

    const isFullscreenPage = matchPath(pathname, {
      path: '/moves/:moveOrderId/payment-requests/:id',
    });

    const siteClasses = classnames('site', {
      [`site--fullscreen`]: isFullscreenPage,
    });

    return (
      <div className={siteClasses}>
        <FOUOHeader />
        {displayChangeRole && <Link to="/select-application">Change user role</Link>}
        {!hideHeaderPPM && <QueueHeader />}
        <main role="main" className="site__content site-office__content">
          <ConnectedLogoutOnInactivity />

          {hasError && <SomethingWentWrong error={error} info={info} />}

          <Suspense fallback={<LoadingPlaceholder />}>
            {!hasError && (
              <Switch>
                {/* no auth */}
                <Route path="/sign-in" component={SignIn} />

                {/* PPM */}
                <PrivateRoute
                  path="/queues/:queueType/moves/:moveId"
                  component={MoveInfo}
                  requiredRoles={[roleTypes.PPM]}
                />
                <PrivateRoute path="/queues/:queueType" component={Queues} requiredRoles={[roleTypes.PPM]} />

                {/* TXO */}
                <PrivateRoute path="/moves/queue" exact component={MoveQueue} requiredRoles={[roleTypes.TOO]} />
                <PrivateRoute path="/invoicing/queue" component={PaymentRequestQueue} requiredRoles={[roleTypes.TIO]} />

                {/* PPM & TXO conflicting routes - select based on user role */}
                {selectedRole === roleTypes.PPM ? ppmRoutes : txoRoutes}

                {/* TEMP FOR DEV ONLY */}
                <PrivateRoute
                  path="/too/:moveOrderId/customer/:customerId"
                  component={CustomerDetails}
                  requiredRoles={[roleTypes.TOO]}
                />

                <PrivateRoute exact path="/select-application" component={ConnectedSelectApplication} />
                {/* ROOT */}
                <PrivateRoute
                  exact
                  path="/"
                  render={(routeProps) => {
                    switch (selectedRole) {
                      case roleTypes.PPM:
                        return <Queues queueType="new" {...routeProps} />;
                      case roleTypes.TIO:
                        return <PaymentRequestIndex {...routeProps} />;
                      case roleTypes.TOO:
                        return <TOO {...routeProps} />;
                      default:
                        // User has unknown role or shouldn't have access
                        return <div />;
                    }
                  }}
                />
              </Switch>
            )}
          </Suspense>
        </main>
      </div>
    );
  }
}

OfficeApp.propTypes = {
  loadInternalSchema: PropTypes.func.isRequired,
  loadPublicSchema: PropTypes.func.isRequired,
  loadUser: PropTypes.func.isRequired,
  location: LocationShape,
  userIsLoggedIn: PropTypes.bool,
  userRoles: UserRolesShape,
  activeRole: PropTypes.string,
};

OfficeApp.defaultProps = {
  location: { pathname: '' },
  userIsLoggedIn: false,
  userRoles: [],
  activeRole: null,
};

const mapStateToProps = (state) => {
  const user = selectCurrentUser(state);
  return {
    swaggerError: state.swaggerInternal.hasErrored,
    userIsLoggedIn: user.isLoggedIn,
    userRoles: user.roles,
    activeRole: state.auth.activeRole,
  };
};

const mapDispatchToProps = (dispatch) =>
  bindActionCreators(
    {
      loadInternalSchema: loadInternalSchemaAction,
      loadPublicSchema: loadPublicSchemaAction,
      loadUser: loadUserAction,
    },
    dispatch,
  );

export default withContext(withRouter(connect(mapStateToProps, mapDispatchToProps)(OfficeApp)));
