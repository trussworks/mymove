/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import { mount } from 'enzyme';

import HHGShipmentCard from '.';

import { formatCustomerDate } from 'utils/formatters';

const defaultProps = {
  editPath: '',
  onEditClick: () => {},
  shipmentNumber: 1,
  shipmentId: '#ABC123K',
  requestedPickupDate: new Date('01/01/2020').toISOString(),
  pickupLocation: {
    street_address_1: '17 8th St',
    city: 'New York',
    state: 'NY',
    postal_code: '11111',
  },
  releasingAgent: {
    firstName: 'Jo',
    lastName: 'Xi',
    phone: '(555) 555-5555',
    email: 'jo.xi@email.com',
  },
  requestedDeliveryDate: new Date('03/01/2020').toISOString(),
  destinationZIP: '73523',
  receivingAgent: {
    firstName: 'Dorothy',
    lastName: 'Lagomarsino',
    phone: '(999) 999-9999',
    email: 'dorothy.lagomarsino@email.com',
  },
  remarks:
    'This is 500 characters of customer remarks right here. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
};

function mountHHGShipmentCard(props = defaultProps) {
  return mount(<HHGShipmentCard {...props} />);
}
describe('HHGShipmentCard component', () => {
  it('renders component with all fields', () => {
    const wrapper = mountHHGShipmentCard();
    const tableHeaders = [
      'Requested pickup date',
      'Pickup location',
      'Releasing agent',
      'Requested delivery date',
      'Destination',
      'Receiving agent',
      'Remarks',
    ];
    const { street_address_1: streetAddress1, city, state, postal_code: postalCode } = defaultProps.pickupLocation;
    const {
      firstName: releasingFirstName,
      lastName: releasingLastName,
      phone: releasingTelephone,
      email: releasingEmail,
    } = defaultProps.releasingAgent;
    const {
      firstName: receivingFirstName,
      lastName: receivingLastName,
      phone: receivingTelephone,
      email: receivingEmail,
    } = defaultProps.receivingAgent;
    const tableData = [
      formatCustomerDate(defaultProps.requestedPickupDate),
      `${streetAddress1} ${city}, ${state} ${postalCode}`,
      `${releasingFirstName} ${releasingLastName} ${releasingTelephone} ${releasingEmail}`,
      formatCustomerDate(defaultProps.requestedDeliveryDate),
      defaultProps.destinationZIP,
      `${receivingFirstName} ${receivingLastName} ${receivingTelephone} ${receivingEmail}`,
    ];

    tableHeaders.forEach((label, index) => expect(wrapper.find('dt').at(index).text()).toBe(label));
    tableData.forEach((label, index) => expect(wrapper.find('dd').at(index).text()).toBe(label));
    expect(wrapper.find('.remarksCell').text()).toBe(defaultProps.remarks);
  });

  it('should render without releasing/receiving agents and remarks', () => {
    const wrapper = mountHHGShipmentCard({ ...defaultProps, releasingAgent: null, receivingAgent: null, remarks: '' });
    const tableHeaders = ['Requested pickup date', 'Pickup location', 'Requested delivery date', 'Destination'];
    const { street_address_1: streetAddress1, city, state, postal_code: postalCode } = defaultProps.pickupLocation;
    const tableData = [
      formatCustomerDate(defaultProps.requestedPickupDate),
      `${streetAddress1} ${city}, ${state} ${postalCode}`,
      formatCustomerDate(defaultProps.requestedDeliveryDate),
      defaultProps.destinationZIP,
    ];
    tableHeaders.forEach((label, index) => expect(wrapper.find('dt').at(index).text()).toBe(label));
    tableData.forEach((label, index) => expect(wrapper.find('dd').at(index).text()).toBe(label));
    expect(wrapper.find('.remarksCell').length).toBe(0);
  });
});