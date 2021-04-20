import React from 'react';
import LoadingIcon from '../../images/loading.gif';
import Modal from '../../components/Modal';
import JKModal from 'jk-modal';

const useFileUpload = ({ setData }) => {
  const sendToConvert = ({
    file,
    doSuccess = () => {},
    doError = () => {},
  }) => {
    console.log('send to api', file);

    const data = new FormData();
    data.append('file', file, file.name)

    fetch('http://localhost:8000/excel', {
      method: 'post',
      body: data,
    }).then(function(response) {
      doSuccess();
    }).catch(function(e) {
      console.log(e);
      doError()
    });

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
    console.log('read file handler')
    fileProcess(file);
  };

  const fileProcess = file => {
    setData(<img width="30" height="30" src={LoadingIcon} alt="" />);
    sendToConvert({
      file: file,
      doSuccess: params => {
        callModalPopup({ type: 'sendSuccess', ...params });
        setData(false);
      },
      doError: params => {
        callModalPopup({ type: 'sendError', ...params });
      },
    });
  };

  return { readFileHandler };
};

export default useFileUpload;
