/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import { render, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';

import SubmitMoveForm from './SubmitMoveForm';

describe('SubmitMoveForm component', () => {
  const testProps = {
    onSubmit: jest.fn(),
    onPrint: jest.fn(),
  };

  it('renders the signature and date inputs', () => {
    const { getByLabelText } = render(<SubmitMoveForm {...testProps} />);
    expect(getByLabelText('Signature')).toBeInTheDocument();
    expect(getByLabelText('Signature')).toBeRequired();
    expect(getByLabelText('Date')).toBeInTheDocument();
    expect(getByLabelText('Date')).toBeDisabled();
  });

  it('shows an error message if trying to submit an invalid form', async () => {
    const { getByTestId, getByText } = render(<SubmitMoveForm {...testProps} />);
    const submitBtn = getByTestId('wizardCompleteButton');

    userEvent.click(submitBtn);

    await waitFor(() => {
      expect(getByText('Required')).toBeInTheDocument();
    });
  });

  it('submits the form when its valid', async () => {
    const { getByLabelText, getByTestId } = render(<SubmitMoveForm {...testProps} />);

    const signatureInput = getByLabelText('Signature');
    const submitBtn = getByTestId('wizardCompleteButton');

    userEvent.type(signatureInput, 'My Name');
    userEvent.click(submitBtn);

    await waitFor(() => {
      expect(testProps.onSubmit).toHaveBeenCalled();
    });
  });

  it('implements the onPrint handler', () => {
    const { getByText } = render(<SubmitMoveForm {...testProps} />);

    const printBtn = getByText('Print');
    userEvent.click(printBtn);

    expect(testProps.onPrint).toHaveBeenCalled();
  });
});
