import React, { useState } from 'react';
import { Grid, GridContainer, Alert } from '@trussworks/react-uswds';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import ScrollToTop from 'components/ScrollToTop';
import ContactInfoForm from 'components/Customer/ContactInfoForm';
import { ServiceMemberShape } from 'types/customerShapes';
import { patchServiceMember, getResponseError } from 'services/internalApi';
import { updateServiceMember as updateServiceMemberAction } from 'store/entities/actions';
import { selectServiceMemberFromLoggedInUser } from 'store/entities/selectors';
import requireCustomerState from 'containers/requireCustomerState/requireCustomerState';
import { profileStates } from 'constants/customerStates';
import { customerRoutes } from 'constants/routes';

export const ContactInfo = ({ serviceMember, updateServiceMember, push }) => {
  const initialValues = {
    telephone: serviceMember?.telephone || '',
    secondary_telephone: serviceMember?.secondary_telephone || '',
    personal_email: serviceMember?.personal_email || '',
    phone_is_preferred: serviceMember?.phone_is_preferred,
    email_is_preferred: serviceMember?.email_is_preferred,
  };
  const [serverError, setServerError] = useState(null);

  const handleNext = () => {
    return push(customerRoutes.CURRENT_DUTY_STATION_PATH);
  };

  const handleBack = () => {
    return push(customerRoutes.NAME_PATH);
  };

  const handleSubmit = (values) => {
    const { nextPage } = values;

    if (values) {
      const payload = {
        id: serviceMember.id,
        telephone: values?.telephone,
        secondary_telephone: values?.secondary_telephone,
        personal_email: values?.personal_email,
        phone_is_preferred: values?.phone_is_preferred,
        email_is_preferred: values?.email_is_preferred,
      };
      return patchServiceMember(payload)
        .then(updateServiceMember)
        .then(() => {
          if (nextPage === 'back') {
            handleBack();
          } else {
            handleNext();
          }
        })
        .catch((e) => {
          // TODO - error handling - below is rudimentary error handling to approximate existing UX
          // Error shape: https://github.com/swagger-api/swagger-js/blob/master/docs/usage/http-client.md#errors
          const { response } = e;
          const errorMessage = getResponseError(response, 'failed to update service member due to server error');
          setServerError(errorMessage);
        });
    }

    return Promise.resolve();
  };

  return (
    <GridContainer>
      <ScrollToTop otherDep={serverError} />

      {serverError && (
        <Grid row>
          <Grid col desktop={{ col: 8, offset: 2 }}>
            <Alert type="error" heading="An error occurred">
              {serverError}
            </Alert>
          </Grid>
        </Grid>
      )}

      <Grid row>
        <Grid col desktop={{ col: 8, offset: 2 }}>
          <ContactInfoForm initialValues={initialValues} onSubmit={handleSubmit} />
        </Grid>
      </Grid>
    </GridContainer>
  );
};

ContactInfo.propTypes = {
  serviceMember: ServiceMemberShape.isRequired,
  updateServiceMember: PropTypes.func.isRequired,
  push: PropTypes.func.isRequired,
};

const mapDispatchToProps = {
  updateServiceMember: updateServiceMemberAction,
};

const mapStateToProps = (state) => ({
  serviceMember: selectServiceMemberFromLoggedInUser(state),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(requireCustomerState(ContactInfo, profileStates.NAME_COMPLETE));
