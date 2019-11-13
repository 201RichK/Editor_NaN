import React, {Component} from 'react'
import EditorAce from 'react-ace'
import Console from './Console/Console'
import './modes'
import './editor.css'
import Axios from 'axios'

export default class Editor extends Component {
    constructor() {
        super()
        this.state = {
            lang: "golang",
            theme: "monokai",
            program: "package main\n\n\n\nfunc main() { \n\n}",
            console: Object
        }
    }

    print = () => {
        Axios.post("http://localhost:8080/run", { source_code: this.refs.ace.editor.getSession().getValue()})
            .then((response) => {
                this.setState({
                    console: response.data
                })
            })
    }


    render = () => {
        return (
            <div id="Editor" className="flex just three column">
                <EditorAce ref="ace" mode={this.state.lang} theme={this.state.theme} width="100%" fontSize="16px" defaultValue={this.state.program} height="100%"/>
                <button onClick={this.print}>Run</button>
                <Console console={this.state.console}></Console>
            </div>
        )
    }
} 