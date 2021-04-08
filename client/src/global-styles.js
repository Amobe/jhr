import { createGlobalStyle } from 'styled-components';

const GlobalStyle = createGlobalStyle`
  :root {
    --gray-1000: #171718;
    --gray-900: #292734;
    --gray-800: #42434a;
    --gray-700: #5c5d65;
    --gray-600: #7d7d87;
    --gray-500: #9fa0ab;
    --gray-400: #b7b8c4;
    --gray-300: #d9dae4;
    --gray-200: #ededf1;
    --gray-100: #f4f4f6;
    --white: #ffffff;
    --black: #000000;
  }

  html,
  body {
    height: 100%;
    width: 100%;
    color: var(--gray-800);
    margin: 0;
    user-select:none;
    /* touch-action: manipulation; */
    -webkit-touch-callout: none;
    -webkit-overflow-scrolling: touch;
    -webkit-tap-highlight-color: transparent;
  }

  body {
    font-family: -apple-system, BlinkMacSystemFont, "PingFang SC",sans-serif;
    overflow-x: hidden;
    background: var(--gray-100);
  }

  *, *:before, *:after {
    box-sizing: border-box;
  }
`;

export default GlobalStyle;
