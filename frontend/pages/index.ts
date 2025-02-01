import React from 'react';
import { NextPage } from 'next';

const HomePage: NextPage = () => {
  return React.createElement(
    'div',
    null,
    React.createElement('h1', null, 'ホームページへようこそ'),
    React.createElement('p', null, 'これが基本的なNext.jsのホームページの例です。')
  );
};

export default HomePage;
