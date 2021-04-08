import styled, { css } from 'styled-components';

const getBtn = props => {
  if (props.primary) {
    return css`
      margin-top: 20px;
      color: #ffffff;
      font-size: 14px;
      font-weight: 600;
      padding: 8px 20px;
      background: #349077;
    `;
  }
  return css`
    height: auto;
    margin-top: 15px;
    color: #7f7f7f;
    font-size: 15px;
    font-weight: 500;
    line-height: 21px;
    background: #ffffff;
  `;
};

export const Container = styled.div`
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100vh;
`;

export const Content = styled.div`
  position: relative;
  background: #ffffff;
  border-radius: 16px;
  padding: 20px 36px;
  min-width: 292px;
  text-align: center;
  max-width: 480px;
  z-index: 2;
`;

export const ImageDiv = styled.div`
  width: 100%;
  max-width: 165px;
  margin: auto;
  padding-top: calc(154 / 198 * 100%);
  background-image: url(${({ imageSrc }) => imageSrc || ''});
  background-size: contain;
  background-position: center;
  transform: scale(1.2);
`;

export const Title = styled.h1`
  margin: 0;
  color: #363535;
  font-size: 17px;
  font-weight: 500;
  line-height: 23px;
`;

export const Description = styled.p`
  margin: 5px auto 0;
  max-width: 220px;
  color: #6a6666;
  font-size: 14px;
  font-weight: normal;
  line-height: 150%;

  & > span {
    color: #1f8af4;
    text-decoration: underline;
    cursor: pointer;
  }
`;

export const Hint = styled.p`
  margin: 20px 0 0;
  max-width: 220px;
  max-height: 75px;
  overflow: auto;
  color: #42a1fe;
  font-size: 14px;
  font-weight: normal;
  line-height: 150%;
  text-align: left;
`;

export const ButtonGroup = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

export const Button = styled.button`
  /* width: 160px; */
  text-align: center;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  padding: 10px 20px;
  min-width: 100px;

  ${getBtn}
`;

export const Mask = styled.div`
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1;
`;
