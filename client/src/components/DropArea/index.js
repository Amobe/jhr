import React, { useState } from 'react';
import allowFiles from '../../data/methods/allowFIles';
import {
  Wrapper,
  DropContainer,
  ShowText,
  ControlPanel,
  ErrorText,
  PanelRight,
} from './style';
import Button from '../Button';
import useFileUpload from '../../data/hooks/useFileUpload';

const DropArea = () => {
  const [data, setData] = useState(false);
  const [error, setError] = useState(false);
  const [active, setActive] = useState(false);

  const { readFileHandler } = useFileUpload({ setData });

  const onDrop = e => {
    console.log('on drop');
    e.preventDefault();

    const {
      dataTransfer: { files },
    } = e;

    const { length } = files;
    if (length === 0) return false;

    removeFile();

    const file = files[0];

    try {
      allowFiles(file);
      setError(false);
      readFileHandler(file);
    } catch (e) {
      setError(e.message);
    }
  };

  const onDragOver = e => {
    e.preventDefault();
  };

  const removeFile = () => {
    setData(false);
    setActive(false);
  };

  return (
    <Wrapper>
      <DropContainer
        active={active}
        onDrop={e => onDrop(e)}
        onDragOver={e => onDragOver(e)}
        onDragEnter={() => setActive(true)}
        onDragLeave={() => setActive(false)}
      >
        {!data && <ShowText>Drag the file</ShowText>}
        {data}
      </DropContainer>

      <ControlPanel>
        {error && <ErrorText>{error}</ErrorText>}
        {data && (
          <PanelRight>
            <Button onClick={() => removeFile()}>Stop celebrate</Button>
          </PanelRight>
        )}
      </ControlPanel>
    </Wrapper>
  );
};

export default DropArea;
