import React from 'react';
import PropTypes from 'prop-types';
import dialogMap from './dialogMap';
import {
  Container,
  Content,
  ImageDiv,
  Title,
  Description,
  Hint,
  ButtonGroup,
  Button,
  Mask,
} from './style';

const Modal = ({
  type,
  customImage,
  customTitle = () => {},
  customDesc = () => {},
  customConfirmText,
  customCancelText,
  customCelecrateImg,
  closeModal,
  onConfirmClick = () => {},
  onCancelClick = () => {},
}) => {
  const pool = dialogMap;
  const data = type && pool[type] ? pool[type] : {};

  const { image, title, desc, hint, confirmText, cancelText, celecrateImg, htmlTitle } = data;

  const handleConfirmOnClick = () => {
    onConfirmClick();
    closeModal();
  };

  const handleCancelOnClick = () => {
    onCancelClick();
    closeModal();
  };

  const secondBtnShow = customCancelText || cancelText;
  const celebrateImgShow = customCelecrateImg || celecrateImg

  return (
    <Container>
      <Content>
        {customImage && <ImageDiv imageSrc={customImage || image} />}
        <Title>
          {customTitle() || title} {htmlTitle && htmlTitle()}
        </Title>
        <Description>{customDesc() || desc}</Description>
        {hint && <Hint>{hint}</Hint>}
        {
          celebrateImgShow && (
            <img width="240" height="240" src={celecrateImg} alt="" />
          )
        }
        <ButtonGroup>
          <Button primary onClick={handleConfirmOnClick}>
            {customConfirmText || confirmText}
          </Button>
          {secondBtnShow && (
            <Button onClick={handleCancelOnClick}>
              {customCancelText || cancelText}
            </Button>
          )}
        </ButtonGroup>
      </Content>
      <Mask onClick={closeModal} />
    </Container>
  );
};

Modal.propTypes = {
  type: PropTypes.string,
  error: PropTypes.bool,
  customImage: PropTypes.string,
  customTitle: PropTypes.func,
  customDesc: PropTypes.func,
  customConfirmText: PropTypes.string,
  customCancelText: PropTypes.string,
  closeModal: PropTypes.func,
  onConfirmClick: PropTypes.func,
  onCancelClick: PropTypes.func,
};

export default Modal;
