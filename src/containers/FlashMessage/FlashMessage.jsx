import React, { useEffect, useRef } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { useLocation } from 'react-router-dom';

import { clearFlashMessage as clearFlashMessageAction } from 'store/flash/actions';
import Alert from 'shared/Alert';

export const FlashMessage = ({ children, flash, clearFlashMessage }) => {
  const { pathname } = useLocation();

  const isInitialMount = useRef(true);

  useEffect(() => {
    if (isInitialMount.current) {
      isInitialMount.current = false;
    } else {
      // only clear flash if an update occurs after mount
      clearFlashMessage();
    }
  }, [pathname, clearFlashMessage]);

  if (!flash) return null;

  const { message, title, type, key } = flash;

  // display children if passed in as a custom flash (using the key)
  if (key && children) {
    return children;
  }

  // otherwise fall back to alert
  // message & type are required
  if (!message || !type) return null;
  return (
    <Alert type={type} heading={title}>
      {message}
    </Alert>
  );
};

FlashMessage.propTypes = {
  children: PropTypes.node,
  flash: PropTypes.shape({
    type: PropTypes.string,
    title: PropTypes.string,
    message: PropTypes.string,
    key: PropTypes.string,
  }),
  clearFlashMessage: PropTypes.func.isRequired,
};

FlashMessage.defaultProps = {
  children: null,
  flash: null,
};

const mapDispatchToProps = {
  clearFlashMessage: clearFlashMessageAction,
};

const mapStateToProps = (state) => ({
  flash: state.flash.flashMessage,
});

export default connect(mapStateToProps, mapDispatchToProps)(FlashMessage);
