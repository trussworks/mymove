import React from 'react';
import PropTypes from 'prop-types';

import { ErrorMessage as UswdsErrorMessage } from '@trussworks/react-uswds';

export const ErrorMessage = ({ display, children, ...props }) => {
  // eslint-disable-next-line react/jsx-props-no-spreading
  return display ? <UswdsErrorMessage {...props}>{children}</UswdsErrorMessage> : null;
};

ErrorMessage.propTypes = {
  display: PropTypes.bool.isRequired,
  children: PropTypes.string.isRequired,
};

export default ErrorMessage;