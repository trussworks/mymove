import React from 'react';
import { Formik } from 'formik';
import * as Yup from 'yup';
import PropTypes from 'prop-types';

import { CustomerContactInfoFields } from 'components/form/CustomerContactInfoFields';
import SectionWrapper from 'components/Customer/SectionWrapper';
import { Form } from 'components/form/Form';
import WizardNavigation from 'components/Customer/WizardNavigation/WizardNavigation';
import formStyles from 'styles/form.module.scss';

const ContactInfoForm = ({ initialValues, onSubmit, onBack }) => {
  const validationSchema = Yup.object().shape({
    telephone: Yup.string().min(10, 'Number must have 10 digits and a valid area code').required('Required'),
    secondary_phone: Yup.string().min(10, 'Number must have 10 digits and a valid area code'),
    personal_email: Yup.string()
      .matches(/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$/, 'Must be a valid email address')
      .required('Required'),
    phone_is_preferred: Yup.bool(),
    email_is_preferred: Yup.bool(),
  });

  return (
    <Formik initialValues={initialValues} onSubmit={onSubmit} validationSchema={validationSchema} validateOnMount>
      {({ isValid, handleSubmit, isSubmitting }) => {
        return (
          <Form>
            <h1>Your contact info</h1>
            <SectionWrapper>
              <div className="tablet:margin-top-neg-3">
                <CustomerContactInfoFields />
              </div>
            </SectionWrapper>
            <div className={formStyles.formActions}>
              <WizardNavigation
                onBackClick={onBack}
                disableNext={!isValid || isSubmitting}
                onNextClick={handleSubmit}
              />
            </div>
          </Form>
        );
      }}
    </Formik>
  );
};

ContactInfoForm.propTypes = {
  initialValues: PropTypes.shape({
    telephone: PropTypes.string,
    secondary_phone: PropTypes.string,
    personal_email: PropTypes.string,
    phone_is_preferred: PropTypes.bool,
    email_is_preferred: PropTypes.bool,
  }).isRequired,
  onSubmit: PropTypes.func.isRequired,
  onBack: PropTypes.func.isRequired,
};

export default ContactInfoForm;
