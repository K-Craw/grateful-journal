import React from 'react';
import FetchDemo from './components/FetchDemo';
import Entries from './components/Entries';

function App() {
  return (
    <div className="App">
      <h1>Post what you want</h1>
      <Entries/>
      <h1>Posts</h1>
      <FetchDemo />
    </div>
  );
}
export default App;