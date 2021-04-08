import React from 'react';
import { Wrapper } from './style';
import DropArea from '../../components/DropArea';
import GlobalStyle from '../../global-styles';

const App = () => {
  return (
    <>
      <Wrapper>
        <DropArea />
      </Wrapper>
      <GlobalStyle />
    </>
  );
};

export default App;
