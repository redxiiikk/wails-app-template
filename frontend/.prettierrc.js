/**
 * @see https://prettier.io/docs/en/configuration.html
 * @type {import('prettier').Config}
 */
const config = {
  trailingComma: 'all',
  printWidth: 80,
  tabWidth: 2,
  semi: true,
  singleQuote: true,
  overrides: [
    // ignore wails generated files
    {
      files: 'src/wailsjs/**/*.js',
      options: {},
    },
  ],
};

export default config;
