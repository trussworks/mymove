import React from 'react';
import classNames from 'classnames';
import * as PropTypes from 'prop-types';

import styles from './ShipmentContainer.module.scss';

import { SHIPMENT_OPTIONS } from 'shared/constants';
import { ShipmentOptionsOneOf } from 'types/shipment';

const ShipmentContainer = ({ className, children, shipmentType }) => {
  const containerClasses = classNames(
    'container',
    styles.shipmentContainer,
    {
      'container--accent--hhg':
        shipmentType === SHIPMENT_OPTIONS.HHG ||
        shipmentType === SHIPMENT_OPTIONS.HHG_SHORTHAUL_DOMESTIC ||
        shipmentType === SHIPMENT_OPTIONS.HHG_LONGHAUL_DOMESTIC,
      'container--accent--nts': shipmentType === SHIPMENT_OPTIONS.NTS,
    },
    className,
  );

  return <div className={`${containerClasses}`}>{children}</div>;
};

ShipmentContainer.propTypes = {
  className: PropTypes.string,
  children: PropTypes.element.isRequired,
  /** Describes the type of shipment container. */
  shipmentType: ShipmentOptionsOneOf,
};

ShipmentContainer.defaultProps = {
  shipmentType: SHIPMENT_OPTIONS.HHG,
  className: '',
};

export default ShipmentContainer;
