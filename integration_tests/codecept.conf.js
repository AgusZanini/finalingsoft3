/** @type {CodeceptJS.MainConfig} */
exports.config = {
  tests: './*_test.js',
  output: './output',
  helpers: {
    Playwright: {
      browser: 'chromium',
      url: 'https://frontend-fuz7myfmqq-uc.a.run.app',
      show: false,
      timeout: 60000
    }
  },
  include: {
    I: './steps_file.js'
  },
  name: 'integration_tests'
}