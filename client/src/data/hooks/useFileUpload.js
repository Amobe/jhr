import React from 'react';
import LoadingIcon from '../../images/loading.gif';
import Modal from '../../components/Modal';
import JKModal from 'jk-modal';

const useFileUpload = ({ setData }) => {
  const sendToConvert = ({
    filePath,
    doSuccess = () => {},
    doError = () => {},
  }) => {
    console.log('send to api', filePath);

    if (!filePath) return doError();
    doSuccess();
  };

  /**
   * Modal Popup
   */
  const callModalPopup = ({ type, title = '', desc = '' }) => {
    JKModal(
      <Modal
        error
        type={type}
        customTitle={() => title}
        customDesc={() => desc}
      />,
    );
  };

  /**
   * analyze file's type
   */
  const readFileHandler = file => {
    // 不是圖片的檔案類型
    if (!file.type.includes('image')) {
      fileProcess(file);
    } else {
      imageProcess(file);
    }
  };

  const fileProcess = file => {
    setData(<img width="30" height="30" src={LoadingIcon} alt="" />);

    // 將檔案路徑傳給後端，後端去解析
    sendToConvert({
      filePath: file.path,
      doSuccess: params => {
        callModalPopup({ type: 'sendSuccess', ...params });
      },
      doError: params => {
        callModalPopup({ type: 'sendError', ...params });
      },
    });
  };

  const imageProcess = file => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = loadEvt => {
      setData(<img src={loadEvt.target.result} alt="" />);
    };
  };

  return { readFileHandler };
};

export default useFileUpload;
