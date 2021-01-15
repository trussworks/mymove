export default {
  PENDING: 'PENDING',
  REVIEWED: 'REVIEWED',
  SENT_TO_GEX: 'SENT_TO_GEX',
  RECEIVED_BY_GEX: 'RECEIVED_BY_GEX',
  PAID: 'PAID',
};

export const PAYMENT_REQUEST_STATUS_LABELS = {
  PENDING: 'Payment requested',
  REVIEWED: 'Reviewed',
  SENT_TO_GEX: 'Reviewed',
  RECEIVED_BY_GEX: 'Reviewed',
  REVIEWED_AND_ALL_SERVICE_ITEMS_REJECTED: 'Rejected',
  PAID: 'Paid',
};
