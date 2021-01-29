import React from 'react';
import DayPickerInput from 'react-day-picker/DayPickerInput';
import { parseDate, formatDate, defaultDateFormat } from 'shared/dates';

import 'react-day-picker/lib/style.css';

const getDayPickerProps = (disabledDays) => {
  return {
    modifiers: {
      disabled: disabledDays,
    },
  };
};

export default function SingleDatePicker(props) {
  const {
    value = null,
    format = defaultDateFormat,
    onChange,
    onBlur,
    disabled,
    name,
    disabledDays,
    placeholder,
    inputClassName,
  } = props;
  const formatted = parseDate(value);

  return (
    <DayPickerInput
      onDayChange={onChange}
      onDayPickerHide={onBlur}
      placeholder={placeholder}
      parseDate={parseDate}
      formatDate={formatDate}
      format={format}
      value={formatted}
      dayPickerProps={getDayPickerProps(disabledDays)}
      inputProps={{
        disabled,
        name,
        autoComplete: 'off',
        className: inputClassName || 'usa-input',
      }}
    />
  );
}
