import React from 'react';
import { storiesOf } from '@storybook/react';
import { withKnobs, text, object } from '@storybook/addon-knobs';

import OrdersTable from './OrdersTable';

storiesOf('TOO/TIO Components|OrdersTable', module)
  .addDecorator(withKnobs)
  .add('Orders Table', () => (
    <OrdersTable
      ordersInfo={{
        currentDutyStation: object('ordersInfo.currentDutyStation', { name: 'JBSA Lackland' }),
        newDutyStation: object('ordersInfo.newDutyStation', { name: 'JB Lewis-McChord' }),
        issuedDate: text('ordersInfo.issuedDate', '8 Mar 2020'),
        reportByDate: text('ordersInfo.reportByDate', '1 Apr 2020'),
        departmentIndicator: text('ordersInfo.departmentIndicator', '17 (Navy and Marine Corps)'),
        ordersNumber: text('ordersInfo.ordersNumber', '999999999'),
        ordersType: text('ordersInfo.ordersType', 'Permanent Change of Duty Station'),
        ordersTypeDetail: text('ordersInfo.ordersTypeDetail', 'Shipment of HHG permitted'),
        tacMDC: text('ordersInfo.tacMDC', '9999'),
        sacSDN: text('ordersInfo.sacSDN', '999 999999 999'),
      }}
    />
  ));
