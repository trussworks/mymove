import React from 'react';
import { string, object } from 'prop-types';

import styles from './Home.module.scss';

const Helper = ({ containerStyles, helpList, description, title }) => (
  <div
    className={`${styles['helper-container']} padding-top-2 padding-right-1 padding-bottom-2 padding-left-1 margin-bottom-4`}
    style={containerStyles}
  >
    <h3 className="margin-bottom-1">{title}</h3>
    {description && <p>{description}</p>}
    {helpList && (
      <ul>
        {helpList.map((helpText, index) => (
          <li key={helpText} className={index > 0 ? 'margin-top-1' : ''}>
            <span>{helpText}</span>
          </li>
        ))}
      </ul>
    )}
  </div>
);

Helper.propTypes = {
  // eslint-disable-next-line react/forbid-prop-types
  containerStyles: object,
  title: string.isRequired,
  helpList: (props, propName, componentName) => {
    if (!props.helpList && props.description) {
      return '';
    }
    console.log('what', props.helpList, props.description);
    if (!Array.isArray(props.helpList)) {
      return new Error('Prop helpList must be of type array');
    }

    if (!props.helpList.length && !props.description) {
      return new Error(`One of 'helpList' or 'description' was not specified in ${componentName}`);
    }

    if (props.helpList.length && typeof props.helpList[0] !== 'string') {
      return new Error(`TypeError: Prop helpList requires an array of strings`);
    }

    return '';
  },
  description: (props, propNam, componentName) => {
    if (!props.description && props.helpList && props.helpList.length) {
      return '';
    }

    if (!props.description && !props.helpList) {
      return new Error(`One of 'helpList' or 'description' was not specified in ${componentName}`);
    }

    if (typeof props.description !== 'string') {
      return new Error('TypeError: Prop description requires a string');
    }

    if (!props.description) {
      return new Error('Prop description cannot be empty');
    }

    return '';
  },
};

Helper.defaultProps = {
  containerStyles: {},
  helpList: [],
  description: '',
};

export default Helper;
