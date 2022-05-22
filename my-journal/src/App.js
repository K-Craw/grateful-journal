import React from 'react';
import FetchDemo from './components/FetchDemo';
import Entries from './components/Entries';
import Post from './components/Post';

function App() {
  return (
    <div className="App">
      <h1>Post what you want</h1>
      <Post/>
      <FetchDemo />
    </div>
  );
}
export default App;