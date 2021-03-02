import { adminBaseURL } from '../../support/constants';

describe('Office Users List Page', function () {
  before(() => {
    cy.prepareAdminApp();
  });

  it('successfully navigates to Office Users page', function () {
    cy.signInAsNewAdminUser();
    // we should be at the office users list page by default,
    // but let's go somewhere else and then come back to make sure the side nav link works
    cy.get('a[href*="system/moves"]').click();
    cy.url().should('eq', adminBaseURL + '/system/moves');

    // now we'll come back to the office users page:
    cy.get('a[href*="system/office_users"]').click();
    cy.url().should('eq', adminBaseURL + '/system/office_users');
    cy.get('header').contains('Office users');

    const columnLabels = ['Id', 'Email', 'First name', 'Last name', 'Transportation Office', 'User Id', 'Active'];
    columnLabels.forEach((label) => {
      cy.get('table').contains(label);
    });
  });
});

describe('Office Users Show Page', function () {
  before(() => {
    cy.prepareAdminApp();
  });

  it('pulls up details page for an office user', function () {
    cy.signInAsNewAdminUser();
    // we tested the side nav in the previous test,
    // so let's work with the assumption that we were already redirected to this page:
    cy.url().should('eq', adminBaseURL + '/system/office_users');
    cy.get('tr[resource="office_users"]').first().click();

    // check that the office user's name is shown in the page title
    cy.get('.ra-field-firstName span.MuiTypography-root')
      .invoke('text')
      .then((firstName) => {
        cy.get('.ra-field-lastName span.MuiTypography-root')
          .invoke('text')
          .then((lastName) => {
            cy.get('#react-admin-title').should('contain', firstName + ' ' + lastName);
          });
      });

    const labels = [
      'Id',
      'User Id',
      'Email',
      'First name',
      'Middle initials',
      'Last name',
      'Telephone',
      'Active',
      'Roles',
      'Transportation Office',
      'Created at',
      'Updated at',
    ];
    labels.forEach((label) => {
      cy.get('.MuiCardContent-root label').contains(label);
    });
  });
});

describe('Office Users Edit Page', function () {
  before(() => {
    cy.prepareAdminApp();
  });

  it('pulls up edit page for an office user', function () {
    cy.signInAsNewAdminUser();
    cy.url().should('eq', adminBaseURL + '/system/office_users');
    cy.get('tr[resource="office_users"]').first().click();

    // grab the office user's ID to check that the correct value is in the url
    cy.get('.ra-field-id span.MuiTypography-root')
      .invoke('text')
      .as('officeUserID')
      .then((officeUserID) => {
        // continue to the edit page
        cy.get('a').contains('Edit').click();
        cy.url().should('eq', adminBaseURL + '/system/office_users/' + officeUserID);
      });

    const disabledFields = ['id', 'email', 'userId', 'createdAt', 'updatedAt'];
    disabledFields.forEach((label) => {
      cy.get('[id="' + label + '"]').should('be.disabled');
    });

    // todo: save update
  });
});
