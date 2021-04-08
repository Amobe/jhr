import { FILE_TYPES } from '../constants';

const allowFiles = file => {
  const { size, type } = file;

  if (!FILE_TYPES.includes(type)) {
    throw new Error('File format must be either png, jpg or xlsx');
  }
  if (size / 1024 / 1024 > 2) {
    throw new Error('File size exceeded the limit of 2MB');
  }
};

export default allowFiles;
