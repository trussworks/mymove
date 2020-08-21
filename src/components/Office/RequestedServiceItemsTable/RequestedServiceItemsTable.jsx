import React from 'react';
import PropTypes from 'prop-types';

import styles from './RequestedServiceItemsTable.module.scss';

import ServiceItemTableHasImg from 'components/ServiceItemTableHasImg';

const RequestedServiceItemsTable = ({ serviceItems }) => {
  return (
    <div className={styles.RequestedServiceItemsTable}>
      <h4>
        Requested service items&nbsp;
        <span>
          ({serviceItems.length} {serviceItems.length === 1 ? 'item' : 'items'})
        </span>
      </h4>
      <ServiceItemTableHasImg serviceItems={serviceItems} />
    </div>
  );
};

RequestedServiceItemsTable.propTypes = {
  serviceItems: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string,
      submittedAt: PropTypes.string,
      serviceItem: PropTypes.string,
      code: PropTypes.string,
      details: PropTypes.object,
    }),
  ).isRequired,
};

export default RequestedServiceItemsTable;