import React, { Component, Fragment } from 'react';

class ProcessingUpload extends Component {
  render() {
    return (
      <Fragment>
        <div className="usa-grid">
          <h2>Your file is being scanned for viruses </h2>
          <p>It will be available within a few minutes.</p>
        </div>
      </Fragment>
    );
  }
}

export default ProcessingUpload;
