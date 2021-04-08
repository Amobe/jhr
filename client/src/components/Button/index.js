import { Wrapper } from './style';

const Button = ({ children, size = 14, onClick = () => {} }) => {
  return (
    <Wrapper size={size} onClick={onClick}>
      {children}
    </Wrapper>
  );
};

export default Button;
