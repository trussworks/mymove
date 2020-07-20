import React from 'react';
import PropTypes from 'prop-types';
import { Fieldset, Label, TextInput } from '@trussworks/react-uswds';
import { v4 as uuidv4 } from 'uuid';

export const AddressFields = ({ legend, className, values, handleChange }) => {
  const addressFieldsUUID = uuidv4();

  return (
    <Fieldset legend={legend} className={className}>
      <Label htmlFor={`mailingAddress1_${addressFieldsUUID}`}>Street address 1</Label>
      <TextInput
        id={`mailingAddress1_${addressFieldsUUID}`}
        data-testid="mailingAddress1"
        name="mailingAddress1"
        type="text"
        onChange={handleChange}
        value={values.mailingAddress1}
      />
      <Label hint=" (optional)" htmlFor={`mailingAddress2_${addressFieldsUUID}`}>
        Street address 2
      </Label>
      <TextInput
        id={`mailingAddress2_${addressFieldsUUID}`}
        data-testid="mailingAddress2"
        name="mailingAddress2"
        type="text"
        onChange={handleChange}
        value={values.mailingAddress2}
      />
      <Label htmlFor={`city_${addressFieldsUUID}`}>City</Label>
      <TextInput
        id={`city_${addressFieldsUUID}`}
        data-testid="city"
        name="city"
        type="text"
        onChange={handleChange}
        value={values.city}
      />
      <Label htmlFor={`state_${addressFieldsUUID}`}>State</Label>
      <TextInput
        id={`state_${addressFieldsUUID}`}
        data-testid="state"
        name="state"
        type="text"
        onChange={handleChange}
        value={values.state}
      />
      <Label htmlFor={`zip_${addressFieldsUUID}`}>ZIP</Label>
      <TextInput
        id={`zip_${addressFieldsUUID}`}
        data-testid="zip"
        inputSize="medium"
        name="zip"
        pattern="[\d]{5}(-[\d]{4})?"
        type="text"
        onChange={handleChange}
        value={values.zip}
      />
    </Fieldset>
  );
};

AddressFields.propTypes = {
  legend: PropTypes.string,
  className: PropTypes.string,
  values: PropTypes.shape({
    mailingAddress1: PropTypes.string,
    mailingAddress2: PropTypes.string,
    city: PropTypes.string,
    state: PropTypes.string,
    zip: PropTypes.string,
  }),
  handleChange: PropTypes.func.isRequired,
};

AddressFields.defaultProps = {
  legend: '',
  className: '',
  values: {},
};

export default AddressFields;
