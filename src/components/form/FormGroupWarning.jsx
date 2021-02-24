import React from 'react';
import PropTypes from 'prop-types';
import { Label, TextInput } from '@trussworks/react-uswds';
import classNames from 'classnames/bind';

import styles from './FormGroupWarning.module.scss';

export const FormGroupWarning = ({ inputLabel, warningMessage }) => {
  return (
    <div className={`${styles.formGroupWarning}`}>
      <Label className={`${styles.label}`} htmlFor="input-type-text">
        {inputLabel}
      </Label>
      <TextInput id="input-type-text" name="input-type-text" type="text" validationStatus="warning" />
      <em>
        <p className={`${styles.hintText}`}>{warningMessage}</p>
      </em>
    </div>
  );
};

FormGroupWarning.propTypes = {
  inputLabel: PropTypes.string.isRequired,
  warningMessage: PropTypes.string.isRequired,
};

export default FormGroupWarning;
