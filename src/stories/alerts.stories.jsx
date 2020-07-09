import React from 'react';
import { Alert } from '@trussworks/react-uswds';

export default {
  title: 'Components|Alerts',
};

export const success = () => (
  <div>
    <Alert heading="Success status" type="success">
      This is a succinct, helpful success message.
    </Alert>
    <Alert slim type="success">
      This is a succinct, helpful success message.
    </Alert>
  </div>
);

export const warning = () => (
  <div>
    <Alert heading="Warning status" type="warning">
      This is a succinct, helpful warning message.
    </Alert>
    <Alert slim type="warning">
      This is a succinct, helpful warning message.
    </Alert>
  </div>
);

export const error = () => (
  <div>
    <Alert heading="Error status" type="error">
      This is a succinct, helpful error message.
    </Alert>
    <Alert slim type="error">
      This is a succinct, helpful error message.
    </Alert>
  </div>
);

export const info = () => (
  <div>
    <Alert heading="Informative status" type="info">
      This is a succinct, helpful info message.
    </Alert>
    <Alert slim type="info">
      This is a succinct, helpful info message.
    </Alert>
  </div>
);

export const systemError = () => (
  <div>
    <Alert className="usa-alert--system-error">
      This is a succinct, helpful error message. Also inlcuded is an example of some&nbsp;
      <a href="#">link text</a>
      .
      <br />
      This is a second line to test the line height.
    </Alert>
  </div>
);
