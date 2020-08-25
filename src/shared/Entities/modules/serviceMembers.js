import { swaggerRequest } from 'shared/Swagger/request';
import { getClient } from 'shared/Swagger/api';
import { get } from 'lodash';
import { selectLoggedInUser } from './user';

const createBackupContactLabel = 'ServiceMember.createBackupContact';
const loadBackupContactsLabel = 'ServiceMember.loadBackupContacts';
const updateBackupContactLabel = 'ServiceMember.updateBackupContact';

const createServiceMemberLabel = 'ServiceMember.createServiceMember';
export const loadServiceMemberLabel = 'ServiceMember.loadServiceMember';
export const updateServiceMemberLabel = 'ServiceMember.updateServiceMember';

/** Swagger Requests */
export function createBackupContact(serviceMemberId, backupContact) {
  const swaggerTag = 'backup_contacts.createServiceMemberBackupContact';
  return swaggerRequest(
    getClient,
    swaggerTag,
    {
      serviceMemberId,
      createBackupContactPayload: backupContact,
    },
    { label: createBackupContactLabel },
  );
}

export function loadBackupContacts(serviceMemberId, label = loadBackupContactsLabel) {
  const swaggerTag = 'backup_contacts.indexServiceMemberBackupContacts';
  return swaggerRequest(getClient, swaggerTag, { serviceMemberId }, { label });
}

export function updateBackupContact(backupContactId, backupContact, label = updateBackupContactLabel) {
  const swaggerTag = 'backup_contacts.updateServiceMemberBackupContact';
  return swaggerRequest(
    getClient,
    swaggerTag,
    { backupContactId, updateServiceMemberBackupContactPayload: backupContact },
    { label },
  );
}

export function createServiceMember(serviceMember) {
  const swaggerTag = 'service_members.createServiceMember';
  return swaggerRequest(
    getClient,
    swaggerTag,
    { createServiceMemberPayload: serviceMember },
    { label: createServiceMemberLabel },
  );
}

export function loadServiceMember(serviceMemberId, label = loadServiceMemberLabel) {
  const swaggerTag = 'service_members.showServiceMember';
  return swaggerRequest(getClient, swaggerTag, { serviceMemberId }, { label });
}

export function updateServiceMember(serviceMemberId, serviceMember, label = updateServiceMemberLabel) {
  const swaggerTag = 'service_members.patchServiceMember';
  return swaggerRequest(
    getClient,
    swaggerTag,
    { serviceMemberId, patchServiceMemberPayload: serviceMember },
    { label },
  );
}

/** Selectors */
export function selectServiceMember(state, serviceMemberId) {
  return get(state, `entities.serviceMembers.${serviceMemberId}`, {});
}

export function selectServiceMemberFromLoggedInUser(state) {
  const user = selectLoggedInUser(state);
  const serviceMemberId = user.service_member;
  return selectServiceMember(state, serviceMemberId);
}

export function selectServiceMemberForOrders(state, ordersId) {
  const orders = get(state, `entities.orders.${ordersId}`);
  if (!orders) {
    return {};
  }
  const serviceMember = get(state, `entities.serviceMembers.${orders.service_member_id}`);
  return serviceMember || {};
}

export function selectServiceMemberForMove(state, moveId) {
  const move = get(state, `entities.moves.${moveId}`);
  if (!move) {
    return {};
  }
  const serviceMemberId = move.service_member_id;
  const serviceMember = selectServiceMember(state, serviceMemberId);
  if (!serviceMember) {
    return {};
  }
  return serviceMember;
}

export function selectBackupContactForServiceMember(state, serviceMemberId) {
  const backupContact = Object.values(state.entities.backupContacts).find((backupContact) => {
    return backupContact.service_member_id === serviceMemberId;
  });
  return backupContact || {};
}
