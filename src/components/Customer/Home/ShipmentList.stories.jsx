/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';

import ShipmentList from './ShipmentList';

import { SHIPMENT_OPTIONS } from 'shared/constants';

export const Basic = () => (
  <div className="grid-container">
    <h3>Single Shipment</h3>
    <ShipmentList shipments={[{ id: '#0001', type: SHIPMENT_OPTIONS.PPM }]} onShipmentClick={() => {}} />
    <br />
    <h3>Multiple shipments</h3>
    <ShipmentList
      shipments={[
        { id: '#0001', type: SHIPMENT_OPTIONS.HHG },
        { id: '#0002', type: SHIPMENT_OPTIONS.NTS },
        { id: '#0003', type: SHIPMENT_OPTIONS.PPM },
      ]}
      onShipmentClick={() => {}}
    />
  </div>
);

export default {
  title: 'Customer Components | ShipmentList',
};
