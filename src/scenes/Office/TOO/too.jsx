import React from 'react';
import PrivateRoute from 'shared/User/PrivateRoute';

function TOO() {
  return <h1>TOO Placeholder Page</h1>;
}

function TOORoute() {
  return <PrivateRoute path="/ghc/too" component={TOO} />;
}

export default TOORoute;
