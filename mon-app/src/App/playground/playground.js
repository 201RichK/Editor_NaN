import React, {Component} from 'react';
import './playground.css';
import Editor from './Editor/editor'
import Exercice from './Exercice/exercice'
import Timeline from './Timeline/timeline'

export default class Playground extends Component {
  render = () => {
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
    )
  }
}
