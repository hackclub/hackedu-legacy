exports.config = {
  baseUrl: 'http://localhost:9000/',
  specs: ['./test/integration/*.js'],

  capabilities: {
    'browserName': 'chrome'
  },

  jasmineNodeOpts: {
    showColors: true,
    defaultTimeoutInterval: 30000
  }
}
