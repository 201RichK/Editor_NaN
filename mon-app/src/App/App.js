import React from 'react';
import './App.css';
import Editor from './Editor/editor'
import Exercice from './Exercice/exercice'
import Timeline from './Timeline/timeline'

function App() {


  return (
    <div className="App flex column one">
      <div className="flex row three">
        <div className="one">
          <Exercice></Exercice>
        </div>
        <Editor></Editor>
      </div>
      <div className="one">
        <Timeline></Timeline>
      </div>
    </div>
  );
}

export default App;
