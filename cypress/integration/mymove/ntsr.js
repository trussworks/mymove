describe('Customer NTSr Setup flow', function () {
  // profile@comple.te
  const profileCompleteUser = '3b9360a3-3304-4c60-90f4-83d687884077';
  // nts@ntsr.unsubmitted
  const ntsUser = '583cfbe1-cb34-4381-9e1f-54f68200da1b';

  before(() => {
    cy.prepareCustomerApp();
  });

  it('Sets up an NTSr shipment', function () {
    cy.apiSignInAsUser(profileCompleteUser);
    customerCreatesAnNTSRShipment();
    customerReviewsNTSRMoveDetails();
  });

  it('Edits an NTSr shipment from review page', function () {
    cy.apiSignInAsUser(ntsUser);
    customerVisitsReviewPage();
    customerEditsNTSRShipmentFromReviewPage();
  });

  it('Edits an NTSr shipment from home page', function () {
    cy.apiSignInAsUser(ntsUser);
    customerEditsNTSRShipmentFromHomePage();
  });
});

function customerEditsNTSRShipmentFromHomePage() {
  cy.get('[data-testid="shipment-list-item-container"]').contains('NTS-R').click();
  cy.get('input[data-testid="remarks"]').clear().type('Warning: glass').blur();

  cy.get('button').contains('Save').click();
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/moves\/[^/]+\/review/);
  });
}

function customerReviewsNTSRMoveDetails() {
  cy.get('[data-testid="review-move-header"]').contains('Review your details');

  // Requested delivery date
  cy.get('[data-testid="ntsr-summary"]').last().contains('02 Jan 2020');

  // Destination
  cy.get('[data-testid="ntsr-summary"]').last().contains('30813');

  // Receiving agent
  cy.get('[data-testid="ntsr-summary"]').last().contains('James Bond');
  cy.get('[data-testid="ntsr-summary"]').last().contains('777-777-7777');
  cy.get('[data-testid="ntsr-summary"]').last().contains('007@example.com');

  // Remarks
  cy.get('[data-testid="ntsr-summary"]').last().contains('some other customer remark');
}

function customerEditsNTSRShipmentFromReviewPage() {
  cy.get('button[data-testid="edit-ntsr-shipment-btn"]').contains('Edit').click();
  cy.get('input[name="delivery.requestedDate"]').clear().type('01/01/2022').blur();
  cy.get('[data-testid="mailingAddress1"]').clear().type('123 Maple street');
  cy.get('input[data-testid="firstName"]').clear().type('Ketchum').blur();
  cy.get('input[data-testid="remarks"]').clear().type('Warning: fragile').blur();
  cy.get('button').contains('Save').click();
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/moves\/[^/]+\/review/);
  });
  cy.get('[data-testid="ntsr-summary"]').contains('01 Jan 2022');
  cy.get('[data-testid="ntsr-summary"]').contains('123 Maple street');
  cy.get('[data-testid="ntsr-summary"]').contains('Ketchum Ash');
  cy.get('[data-testid="ntsr-summary"]').contains('Warning: fragile');
}

function customerVisitsReviewPage() {
  cy.get('button[data-testid="review-and-submit-btn"]').contains('Review and submit').click();
  cy.get('h2[data-testid="review-move-header"]').contains('Review your details');
}

function customerCreatesAnNTSRShipment() {
  cy.get('[data-testid="shipment-selection-btn"]').contains('Plan your shipments').click();
  cy.get('h1').contains('Figure out your shipments');
  cy.nextPage();
  cy.get('h1').contains('How do you want to move your belongings?');
  cy.get('input[type="radio"]').eq(3).check({ force: true });
  cy.nextPage();
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ntsr-start/);
  });

  // pickup date
  cy.get('input[name="delivery.requestedDate"]').type('01/02/2020').blur();

  // no delivery location

  // receiving agent
  cy.get(`[data-testid="firstName"]`).type('James');
  cy.get(`[data-testid="lastName"]`).type('Bond');
  cy.get(`[data-testid="phone"]`).type('7777777777');
  cy.get(`[data-testid="email"]`).type('007@example.com').blur();

  // remarks
  cy.get(`[data-testid="remarks"]`).first().type('some other customer remark');

  cy.nextPage();
}
