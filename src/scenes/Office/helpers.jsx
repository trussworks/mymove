import moment from 'moment';

// Format a date and ignore any time values, e.g. 03-Jan-2018
export function formatDate(date) {
  if (date) {
    return moment(date).format('DD-MMM-YY');
  }
}

// Format a date and include its time, e.g. 03-Jan-2018 21:23
export function formatDateTime(date) {
  if (date) {
    return moment(date).format('DD-MMM-YY HH:mm');
  }
}
