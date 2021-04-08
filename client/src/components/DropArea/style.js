import styled from 'styled-components';

const Wrapper = styled('div')``;

const DropContainer = styled('div')`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 480px;
  height: 300px;
  border: 2px dashed ${props => (props.active ? '#808080' : '#ccc')};
  border-radius: 20px;
  overflow: hidden;
`;

const ShowText = styled('span')`
  color: #ccc;
  font-size: 20px;
`;

const ControlPanel = styled('div')`
  display: flex;
  padding: 30px 20px;
  width: 100%;
`;

const ErrorText = styled('div')`
  color: #d44c4c;
`;

const PanelRight = styled('div')`
  display: flex;
  width: 100%;
  justify-content: flex-end;
`;

export {
  Wrapper,
  DropContainer,
  ShowText,
  ControlPanel,
  PanelRight,
  ErrorText,
};
