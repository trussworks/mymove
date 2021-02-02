import React from 'react';
import { string } from 'prop-types';

import styles from './index.module.scss';

const BypassBlock = ({ anchorLink }) => (
  <div className={styles.bypassBlock}>
    <a href={anchorLink}>Skip to Content</a>
  </div>
);

BypassBlock.propTypes = {
  anchorLink: string,
};

BypassBlock.defaultProps = {
  anchorLink: '#main',
};

export default BypassBlock;
