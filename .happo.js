const { RemoteBrowserTarget } = require('happo.io');
const happoPluginStorybook = require('happo-plugin-storybook');

require('dotenv').config();

module.exports = {
  apiKey: process.env.HAPPO_API_KEY,
  apiSecret: process.env.HAPPO_API_SECRET,
  targets: {
    chrome: new RemoteBrowserTarget('chrome', {
      viewport: '1024x768',
    }),
    // HERE IS A TEST COMMENT TO TRIGGER A HAPPO RUN!
    // TODO - IE is failing because Storybook causes syntax error. Need to investigate
    /*
    'internet explorer': new RemoteBrowserTarget('internet explorer', {
      viewport: '1024x768',
    }),
    */
    // TODO - iOS is failing as of updating Storybook to v6. Need to investigate & fix.
    /*
    'ios-safari': new RemoteBrowserTarget('ios-safari', {
      viewport: '375x667',
    }),
    */
  },
  plugins: [
    happoPluginStorybook({
      outputDir: 'storybook-static',
    }),
  ],
};
