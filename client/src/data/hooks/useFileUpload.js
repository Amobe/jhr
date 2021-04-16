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

    // let output = async() => {
    //   let response = await fetch('http://localhost:8000/excel', {
    //     method: 'post',
    //     body: data,
    //   });
    //   let data = await response.blob()
    //   console.log(data)
    //   let objectUrl = URL.createObjectURL(data);
    //   console.log(objectUrl)
    //   window.open(objectUrl);
    //   doSuccess()
    // }
    // output();


    fetch('http://localhost:8000/excel', {
      method: 'post',
      body: data,
    }).then(function(response) {
      const resp = response.clone();
      console.log(resp);
      resp.blob().then(data => {
        var reader = new FileReader();
        reader.readAsDataURL(data);
        reader.onloadend = function() {
          var base64data = reader.result;
          console.log(base64data);
          var a = document.createElement("a");
          a.href = base64data;
          a.download = "hello.xlsx";
          a.click();
          // window.open(base64data)
          doSuccess();
      }
      });
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
    // 不是圖片的檔案類型
    if (!file.type.includes('image')) {
      fileProcess(file);
    } else {
      imageProcess(file);
    }
  };

  const fileProcess = file => {
    setData(<img width="30" height="30" src={LoadingIcon} alt="" />);
    sendToConvert({
      file: file,
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
