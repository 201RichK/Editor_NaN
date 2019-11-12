import React, {Component} from 'react'
import EditorAce from 'react-ace'
import Console from './Console/Console'
import './modes'
import './editor.css'

export default class Editor extends Component {
    constructor() {
        super()
        this.state = {
            lang: "golang",
            theme: "monokai",
            program: "package main\n\n\n\nfunc main() { \n\n}"
        }
    }


    render = () => {
        return (
            <div id="Editor" className="flex just three column">
                <EditorAce mode={this.state.lang} theme={this.state.theme} width="100%" fontSize="16px" defaultValue={this.state.program} height="100%"/>
                <Console></Console>
            </div>
        )
    }
} 