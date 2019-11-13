import React, {Component} from 'react'
import './console.css'

class Console extends Component {
    constructor(props) {
        super(props)
    }


    render = () => {
        return (
            <div id="Console">
                <ul>
                    <li><strong>Temps:</strong> {this.props.console.time}</li>
                    <li><strong>Erreur:</strong> {this.props.console.stderr}</li>
                    <li><strong>Result:</strong> {this.props.console.stdout}</li>
                </ul>
                

            </div>
        )
    }
}

export default Console