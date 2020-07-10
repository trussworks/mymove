import React from 'react';
import PropTypes from 'prop-types';
import { Fieldset, Label, TextInput } from '@trussworks/react-uswds';

export const AddressFields = ({ legend }) => {
  return (
    <Fieldset legend={legend}>
      <Label htmlFor="mailing-address-1">Street address 1</Label>
      <TextInput id="mailing-address-1" name="mailing-address-1" type="text" />
      <Label hint=" (optional)" htmlFor="mailing-address-2">
        Street address 2
      </Label>
      <TextInput id="mailing-address-2" name="mailing-address-2" type="text" />
      <Label htmlFor="city">City</Label>
      <TextInput id="city" name="city" type="text" />
      <Label htmlFor="state">State</Label>
      <TextInput id="state" name="state" type="text" />
      <Label htmlFor="zip">ZIP</Label>
      <TextInput id="zip" inputSize="medium" name="zip" pattern="[\d]{5}(-[\d]{4})?" type="text" />
    </Fieldset>
  );
};

AddressFields.propTypes = {
  legend: PropTypes.string,
};

AddressFields.defaultProps = {
  legend: '',
};

export default AddressFields;
